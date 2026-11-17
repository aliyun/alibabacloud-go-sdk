// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBankMetaVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BankMetaVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *BankMetaVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *BankMetaVerifyIntlResponseBody
	GetRequestId() *string
	SetResultObject(v *BankMetaVerifyIntlResponseBodyResultObject) *BankMetaVerifyIntlResponseBody
	GetResultObject() *BankMetaVerifyIntlResponseBodyResultObject
}

type BankMetaVerifyIntlResponseBody struct {
	// The return code. A value of 200 indicates success, and other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result information.
	ResultObject *BankMetaVerifyIntlResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s BankMetaVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *BankMetaVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BankMetaVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BankMetaVerifyIntlResponseBody) GetResultObject() *BankMetaVerifyIntlResponseBodyResultObject {
	return s.ResultObject
}

func (s *BankMetaVerifyIntlResponseBody) SetCode(v string) *BankMetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) SetMessage(v string) *BankMetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) SetRequestId(v string) *BankMetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) SetResultObject(v *BankMetaVerifyIntlResponseBodyResultObject) *BankMetaVerifyIntlResponseBody {
	s.ResultObject = v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BankMetaVerifyIntlResponseBodyResultObject struct {
	// The verification result code. Valid values:
	//
	// - 1: The verification is consistent. You are billed for this call.
	//
	// - 2: The verification is inconsistent. You are billed for this call.
	//
	// - 3: No record is found. You are not billed for this call.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The verification details. Valid values:
	//
	// - **101**: The verification is successful.
	//
	// - **201**: The authentication information is inconsistent. The cardholder information is incorrect.
	//
	// - **202**: The authentication information is inconsistent. The bank card is not enabled for authenticated payment.
	//
	// - **203**: The authentication information is inconsistent. The bank card has expired.
	//
	// - **204**: The authentication information is inconsistent. The bank card is restricted.
	//
	// - **205**: The authentication information is inconsistent. The bank card has been confiscated.
	//
	// - **206**: The authentication information is inconsistent. The bank card is invalid.
	//
	// - **207**: The authentication information is inconsistent. The bank card has no corresponding issuing bank.
	//
	// - **208**: The authentication information is inconsistent. The bank card is not initialized or is a dormant card.
	//
	// - **209**: The authentication information is inconsistent. The bank card is a fraudulent or captured card.
	//
	// - **210**: The authentication information is inconsistent. The bank card has been reported as lost.
	//
	// - **211**: The authentication information is inconsistent. The number of incorrect password attempts has exceeded the limit.
	//
	// - **212**: The authentication information is inconsistent. The issuing bank does not support this transaction.
	//
	// - **213**: The authentication information is inconsistent. The card status is abnormal or the card is invalid.
	//
	// - **214**: The authentication information is inconsistent. No mobile phone number is reserved.
	//
	// - **215**: The authentication information is inconsistent. The entered password, expiration date, or CVN2 is incorrect.
	//
	// - **216**: The authentication information is inconsistent. Other card exceptions occurred.
	//
	// - **301**: The verification cannot be performed. The bank card does not support this service.
	//
	// - **302**: The verification cannot be performed. The verification failed or the bank rejected the verification. Contact the issuing bank.
	//
	// - **303**: The verification cannot be performed. The bank card does not support mobile phone number verification.
	//
	// - **304**: The verification cannot be performed. The bank card number is incorrect.
	//
	// - **305**: The verification cannot be performed. Other reasons.
	//
	// - **306**: The verification cannot be performed. The number of verification attempts has exceeded the limit.
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s BankMetaVerifyIntlResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyIntlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) SetBizCode(v string) *BankMetaVerifyIntlResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) SetSubCode(v string) *BankMetaVerifyIntlResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
