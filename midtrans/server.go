package midtrans

import (
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

const SandboxServerKey1 = "SB-Mid-server-0NrVq555rAslGxfgAKi_NqMn"

const SandboxClientKey2 = "SB-Mid-client-mFPo7qREcM9p6x0y"


const IrisCreatorKeySandbox = "IRIS-330198f0-e49d-493f-baae-585cfded355d"
const IrisApproverKeySandbox = "IRIS-1595c12b-6814-4e5a-bbbb-9bc18193f47b"


func SnapParamWithMap() *snap.RequestParamWithMap {
	req := &snap.RequestParamWithMap{
		"transaction_details": map[string]interface{}{
			"order_id":     "MID-GO-TEST-" + Random(),
			"gross_amount": 10000,
		},
	}
	return req

}

func SnapParam() *snap.Request {
	req := & snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-GO-TEST-" + Random(),
			GrossAmt: 100000,
		},
	}
	return req
}

func CoreParam() *coreapi.ChargeReqWithMap {
	req := &coreapi.ChargeReqWithMap{
		"payment_type": "gopay",
		"transaction_details": map[string]interface{}{
			"order_id":     "MID-GO-TEST-" + Random(),
			"gross_amount": 10000,
		},
	}
	return req
}

func Random() string {
	time.Sleep(500 * time.Millisecond)
	return strconv.FormatInt(time.Now().Unix(), 10)
}