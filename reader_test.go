package wire

import (
	"github.com/moov-io/base"
	"os"
	"testing"
)

// TestRead reads wire Files with different BusinessFunctionCodes
func TestRead(t *testing.T) {
	t.Run("BankTransfer", testRead("./test/testdata/fedWireMessage-BankTransfer.txt"))
	t.Run("CustomerTransfer", testRead("./test/testdata/fedWireMessage-CustomerTransfer.txt"))
	t.Run("CustomerTransferPlus", testRead("./test/testdata/fedWireMessage-CustomerTransferPlus.txt"))
	t.Run("CheckSameDaySettlement", testRead("./test/testdata/fedWireMessage-CheckSameDaySettlement.txt"))
	t.Run("DepositSendersAccount", testRead("./test/testdata/fedWireMessage-DepositSendersAccount.txt"))
	t.Run("FEDFundsReturned", testRead("./test/testdata/fedWireMessage-FEDFundsReturned.txt"))
	t.Run("FEDFundsSold", testRead("./test/testdata/fedWireMessage-FEDFundsSold.txt"))
	t.Run("DrawdownRequest", testRead("./test/testdata/fedWireMessage-DrawdownRequest.txt"))
	t.Run("BankDrawdownRequest", testRead("./test/testdata/fedWireMessage-BankDrawdownRequest.txt"))
	t.Run("CustomerCorporateDrawdownRequest", testRead("./test/testdata/fedWireMessage-CustomerCorporateDrawdownRequest.txt"))
	t.Run("ServiceMessage", testRead("./test/testdata/fedWireMessage-ServiceMessage.txt"))
	t.Run("CustomerTransferPlusCOVS", testRead("./test/testdata/fedWireMessage-CustomerTransferPlusCOVS.txt"))
	t.Run("CustomerTransferPlusUnstructuredAddenda", testRead("./test/testdata/fedWireMessage-CustomerTransferPlusUnstructuredAddenda.txt"))
	t.Run("CustomerTransferPlusStructuredRemittance", testRead("./test/testdata/fedWireMessage-CustomerTransferPlusStructuredRemittance.txt"))
}

func testRead(filePathName string) func(t *testing.T) {
	return func(t *testing.T) {
		f, err := os.Open(filePathName)
		if err != nil {
			t.Errorf("%T: %s", err, err)
		}
		defer f.Close()
		r := NewReader(f)

		fwmFile, err := r.Read()
		if err != nil {
			t.Errorf("%T: %s", err, err)
		}
		// ensure we have a validated file structure
		if err = fwmFile.Validate(); err != nil {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestReadInvalidTag(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-InvalidTag.txt")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)

	_, err = r.Read()
	if err != nil {
		if !base.Has(err, NewErrInvalidTag(r.line[:6])) {
			t.Errorf("%T: %s", err, err)
		}
	}

}
