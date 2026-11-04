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
	// The response code. 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-A***B-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification result.
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
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BankMetaVerifyResponseBodyResultObject struct {
	// The verification result.
	//
	// - 1: verification passed (billed).
	//
	// - 2: verification failed (billed).
	//
	// - 3: no records found (not billed).
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The verification details.
	//
	// - **101**: verification passed.
	//
	// - **201**: verification failed. The cardholder information is incorrect.
	//
	// - **202**: verification failed. Authentication payment is not enabled for the bank card.
	//
	// - **203**: verification failed. The bank card has expired.
	//
	// - **204**: verification failed. The bank card is restricted.
	//
	// - **205**: verification failed. The card has been confiscated.
	//
	// - **206**: verification failed. The bank card is invalid.
	//
	// - **207**: verification failed. No matching issuing bank exists.
	//
	// - **208**: verification failed. The card is not initialized or is a dormant card.
	//
	// - **209**: verification failed. The card is fraudulent or has been retained.
	//
	// - **210**: verification failed. The card has been reported lost.
	//
	// - **211**: verification failed. The maximum number of password attempts has been exceeded.
	//
	// - **212**: verification failed. The issuing bank does not support this transaction.
	//
	// - **213**: verification failed. The card status is abnormal or the card is invalid.
	//
	// - **214**: verification failed. No mobile phone number is registered.
	//
	// - **215**: verification failed. The password, expiration date, or CVN2 is incorrect.
	//
	// - **216**: verification failed. Other card exceptions.
	//
	// - **301**: unable to verify. The bank card does not support this service.
	//
	// - **302**: unable to verify. Verification failed or the bank refused verification. Contact the issuing bank.
	//
	// - **303**: unable to verify. The bank card does not support mobile phone number verification.
	//
	// - **304**: unable to verify. The bank card number is incorrect.
	//
	// - **305**: unable to verify. Other reasons.
	//
	// - **306**: unable to verify. The maximum number of verification attempts has been exceeded.
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
