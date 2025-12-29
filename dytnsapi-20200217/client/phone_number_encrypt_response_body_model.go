// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberEncryptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PhoneNumberEncryptResponseBody
	GetCode() *string
	SetData(v []*PhoneNumberEncryptResponseBodyData) *PhoneNumberEncryptResponseBody
	GetData() []*PhoneNumberEncryptResponseBodyData
	SetMessage(v string) *PhoneNumberEncryptResponseBody
	GetMessage() *string
	SetRequestId(v string) *PhoneNumberEncryptResponseBody
	GetRequestId() *string
}

type PhoneNumberEncryptResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details about the returned entries.
	Data []*PhoneNumberEncryptResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberEncryptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptResponseBody) GetCode() *string {
	return s.Code
}

func (s *PhoneNumberEncryptResponseBody) GetData() []*PhoneNumberEncryptResponseBodyData {
	return s.Data
}

func (s *PhoneNumberEncryptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PhoneNumberEncryptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PhoneNumberEncryptResponseBody) SetCode(v string) *PhoneNumberEncryptResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberEncryptResponseBody) SetData(v []*PhoneNumberEncryptResponseBodyData) *PhoneNumberEncryptResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberEncryptResponseBody) SetMessage(v string) *PhoneNumberEncryptResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberEncryptResponseBody) SetRequestId(v string) *PhoneNumberEncryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *PhoneNumberEncryptResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PhoneNumberEncryptResponseBodyData struct {
	// The encrypted phone number.
	//
	// example:
	//
	// 1400513****
	EncryptedNumber *string `json:"EncryptedNumber,omitempty" xml:"EncryptedNumber,omitempty"`
	// The time when the phone number expires.
	//
	// example:
	//
	// 2022-05-27 16:05:23
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The original phone number.
	//
	// example:
	//
	// 1390000****
	OriginalNumber *string `json:"OriginalNumber,omitempty" xml:"OriginalNumber,omitempty"`
	OutId          *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
}

func (s PhoneNumberEncryptResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberEncryptResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptResponseBodyData) GetEncryptedNumber() *string {
	return s.EncryptedNumber
}

func (s *PhoneNumberEncryptResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *PhoneNumberEncryptResponseBodyData) GetOriginalNumber() *string {
	return s.OriginalNumber
}

func (s *PhoneNumberEncryptResponseBodyData) GetOutId() *string {
	return s.OutId
}

func (s *PhoneNumberEncryptResponseBodyData) SetEncryptedNumber(v string) *PhoneNumberEncryptResponseBodyData {
	s.EncryptedNumber = &v
	return s
}

func (s *PhoneNumberEncryptResponseBodyData) SetExpireTime(v string) *PhoneNumberEncryptResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *PhoneNumberEncryptResponseBodyData) SetOriginalNumber(v string) *PhoneNumberEncryptResponseBodyData {
	s.OriginalNumber = &v
	return s
}

func (s *PhoneNumberEncryptResponseBodyData) SetOutId(v string) *PhoneNumberEncryptResponseBodyData {
	s.OutId = &v
	return s
}

func (s *PhoneNumberEncryptResponseBodyData) Validate() error {
	return dara.Validate(s)
}
