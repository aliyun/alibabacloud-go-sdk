// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidPhoneNumberFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvalidPhoneNumberFilterResponseBody
	GetCode() *string
	SetData(v []*InvalidPhoneNumberFilterResponseBodyData) *InvalidPhoneNumberFilterResponseBody
	GetData() []*InvalidPhoneNumberFilterResponseBodyData
	SetMessage(v string) *InvalidPhoneNumberFilterResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvalidPhoneNumberFilterResponseBody
	GetRequestId() *string
}

type InvalidPhoneNumberFilterResponseBody struct {
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **MobileNumberIllegal**: The phone number is invalid.
	//
	// 	- **EncyrptTypeIllegal**: The encryption type is invalid.
	//
	// 	- **MobileNumberTypeNotMatch**: The phone number does not match the encryption type.
	//
	// 	- **CarrierIllegal**: The carrier type is invalid.
	//
	// 	- **AuthCodeNotExist**: The authorization code does not exist.
	//
	// 	- **PortabilityNumberNotSupported**: Mobile number portability is not supported.
	//
	// 	- **Unknown**: An unknown exception occurred.
	//
	// 	- **AuthCodeAndApiNotMatch**: A system exception occurred.
	//
	// 	- **AuthCodeAndApiNotMatch**: The authorization code does not match the API operation.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details about the returned entries.
	Data []*InvalidPhoneNumberFilterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s InvalidPhoneNumberFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvalidPhoneNumberFilterResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvalidPhoneNumberFilterResponseBody) GetData() []*InvalidPhoneNumberFilterResponseBodyData {
	return s.Data
}

func (s *InvalidPhoneNumberFilterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvalidPhoneNumberFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvalidPhoneNumberFilterResponseBody) SetCode(v string) *InvalidPhoneNumberFilterResponseBody {
	s.Code = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) SetData(v []*InvalidPhoneNumberFilterResponseBodyData) *InvalidPhoneNumberFilterResponseBody {
	s.Data = v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) SetMessage(v string) *InvalidPhoneNumberFilterResponseBody {
	s.Message = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) SetRequestId(v string) *InvalidPhoneNumberFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) Validate() error {
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

type InvalidPhoneNumberFilterResponseBodyData struct {
	// The returned filter results.
	//
	// 	- **YES**: the valid phone number. The mappings are returned.
	//
	// 	- **NO**: the invalid phone number. No mappings are returned.
	//
	// example:
	//
	// YES
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
}

func (s InvalidPhoneNumberFilterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InvalidPhoneNumberFilterResponseBodyData) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *InvalidPhoneNumberFilterResponseBodyData) GetEncryptedNumber() *string {
	return s.EncryptedNumber
}

func (s *InvalidPhoneNumberFilterResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *InvalidPhoneNumberFilterResponseBodyData) GetOriginalNumber() *string {
	return s.OriginalNumber
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetCode(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.Code = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetEncryptedNumber(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.EncryptedNumber = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetExpireTime(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetOriginalNumber(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.OriginalNumber = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
