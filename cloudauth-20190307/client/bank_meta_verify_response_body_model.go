// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBankMetaVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BankMetaVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *BankMetaVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *BankMetaVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *BankMetaVerifyResponseBodyResultObject) *BankMetaVerifyResponseBody
	GetResultObject() *BankMetaVerifyResponseBodyResultObject
}

type BankMetaVerifyResponseBody struct {
	// Return code: 200 for success, others for failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 473469C7-A***B-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
	ResultObject *BankMetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s BankMetaVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *BankMetaVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BankMetaVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BankMetaVerifyResponseBody) GetResultObject() *BankMetaVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *BankMetaVerifyResponseBody) SetCode(v string) *BankMetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *BankMetaVerifyResponseBody) SetMessage(v string) *BankMetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *BankMetaVerifyResponseBody) SetRequestId(v string) *BankMetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BankMetaVerifyResponseBody) SetResultObject(v *BankMetaVerifyResponseBodyResultObject) *BankMetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *BankMetaVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type BankMetaVerifyResponseBodyResultObject struct {
	// Verification result.
	//
	// - 1: Consistent (billable)
	//
	// - 2: Inconsistent (billable)
	//
	// - 3: No record found (non-billable)
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Verification details:
	//
	// - **101**: Verification passed.
	//
	// - **201**: Authentication information does not match, cardholder information is incorrect.
	//
	// - **202**: Authentication information does not match, bank card has not enabled authentication payment.
	//
	// - **203**: Authentication information does not match, bank card has expired.
	//
	// - **204**: Authentication information does not match, bank card is a restricted card.
	//
	// - **205**: Authentication information does not match, this card has been confiscated.
	//
	// - **206**: Authentication information does not match, bank card is invalid.
	//
	// - **207**: Authentication information does not match, this card has no corresponding issuing bank.
	//
	// - **208**: Authentication information does not match, the card is uninitialized or a dormant card.
	//
	// - **209**: Authentication information does not match, this card is a cheating card or swallowed card.
	//
	// - **210**: Authentication information does not match, this card has been reported lost.
	//
	// - **211**: Authentication information does not match, the number of password errors exceeds the limit.
	//
	// - **212**: Authentication information does not match, the issuing bank does not support this transaction.
	//
	// - **213**: Authentication information does not match, the card status is abnormal or the card is invalid.
	//
	// - **214**: Authentication information does not match, no mobile phone number reserved.
	//
	// - **215**: Authentication information does not match, the entered password, expiration date, or CVN2 is incorrect.
	//
	// - **216**: Authentication information does not match, other card anomalies.
	//
	// - **301**: Unable to verify, the bank card does not support this service.
	//
	// - **302**: Unable to verify, verification failed or the bank refused to verify, please contact the issuing bank.
	//
	// - **303**: Unable to verify, the bank card does not currently support mobile phone number verification.
	//
	// - **304**: Unable to verify, the bank card number is incorrect.
	//
	// - **305**: Unable to verify, other reasons.
	//
	// - **306**: Unable to verify, the number of verifications exceeds the limit.
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s BankMetaVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *BankMetaVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *BankMetaVerifyResponseBodyResultObject) SetBizCode(v string) *BankMetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *BankMetaVerifyResponseBodyResultObject) SetSubCode(v string) *BankMetaVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *BankMetaVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
