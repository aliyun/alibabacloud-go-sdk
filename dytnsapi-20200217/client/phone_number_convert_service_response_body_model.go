// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberConvertServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PhoneNumberConvertServiceResponseBody
	GetCode() *string
	SetData(v []*PhoneNumberConvertServiceResponseBodyData) *PhoneNumberConvertServiceResponseBody
	GetData() []*PhoneNumberConvertServiceResponseBodyData
	SetMessage(v string) *PhoneNumberConvertServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *PhoneNumberConvertServiceResponseBody
	GetRequestId() *string
}

type PhoneNumberConvertServiceResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*PhoneNumberConvertServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberConvertServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberConvertServiceResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberConvertServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *PhoneNumberConvertServiceResponseBody) GetData() []*PhoneNumberConvertServiceResponseBodyData {
	return s.Data
}

func (s *PhoneNumberConvertServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PhoneNumberConvertServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PhoneNumberConvertServiceResponseBody) SetCode(v string) *PhoneNumberConvertServiceResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBody) SetData(v []*PhoneNumberConvertServiceResponseBodyData) *PhoneNumberConvertServiceResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberConvertServiceResponseBody) SetMessage(v string) *PhoneNumberConvertServiceResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBody) SetRequestId(v string) *PhoneNumberConvertServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBody) Validate() error {
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

type PhoneNumberConvertServiceResponseBodyData struct {
	ConverResult *bool   `json:"ConverResult,omitempty" xml:"ConverResult,omitempty"`
	Number       *string `json:"Number,omitempty" xml:"Number,omitempty"`
	NumberMd5    *string `json:"NumberMd5,omitempty" xml:"NumberMd5,omitempty"`
	NumberSha256 *string `json:"NumberSha256,omitempty" xml:"NumberSha256,omitempty"`
}

func (s PhoneNumberConvertServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberConvertServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberConvertServiceResponseBodyData) GetConverResult() *bool {
	return s.ConverResult
}

func (s *PhoneNumberConvertServiceResponseBodyData) GetNumber() *string {
	return s.Number
}

func (s *PhoneNumberConvertServiceResponseBodyData) GetNumberMd5() *string {
	return s.NumberMd5
}

func (s *PhoneNumberConvertServiceResponseBodyData) GetNumberSha256() *string {
	return s.NumberSha256
}

func (s *PhoneNumberConvertServiceResponseBodyData) SetConverResult(v bool) *PhoneNumberConvertServiceResponseBodyData {
	s.ConverResult = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBodyData) SetNumber(v string) *PhoneNumberConvertServiceResponseBodyData {
	s.Number = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBodyData) SetNumberMd5(v string) *PhoneNumberConvertServiceResponseBodyData {
	s.NumberMd5 = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBodyData) SetNumberSha256(v string) *PhoneNumberConvertServiceResponseBodyData {
	s.NumberSha256 = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
