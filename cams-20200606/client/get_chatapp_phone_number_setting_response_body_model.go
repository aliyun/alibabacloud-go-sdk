// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappPhoneNumberSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetChatappPhoneNumberSettingResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetChatappPhoneNumberSettingResponseBody
	GetCode() *string
	SetMessage(v string) *GetChatappPhoneNumberSettingResponseBody
	GetMessage() *string
	SetModel(v *GetChatappPhoneNumberSettingResponseBodyModel) *GetChatappPhoneNumberSettingResponseBody
	GetModel() *GetChatappPhoneNumberSettingResponseBodyModel
	SetRequestId(v string) *GetChatappPhoneNumberSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetChatappPhoneNumberSettingResponseBody
	GetSuccess() *bool
}

type GetChatappPhoneNumberSettingResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetChatappPhoneNumberSettingResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 42602478-F7C1-58D2-AFFE-88F7A18***6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetChatappPhoneNumberSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatappPhoneNumberSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberSettingResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetChatappPhoneNumberSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatappPhoneNumberSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatappPhoneNumberSettingResponseBody) GetModel() *GetChatappPhoneNumberSettingResponseBodyModel {
	return s.Model
}

func (s *GetChatappPhoneNumberSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatappPhoneNumberSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetChatappPhoneNumberSettingResponseBody) SetAccessDeniedDetail(v string) *GetChatappPhoneNumberSettingResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappPhoneNumberSettingResponseBody) SetCode(v string) *GetChatappPhoneNumberSettingResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappPhoneNumberSettingResponseBody) SetMessage(v string) *GetChatappPhoneNumberSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappPhoneNumberSettingResponseBody) SetModel(v *GetChatappPhoneNumberSettingResponseBodyModel) *GetChatappPhoneNumberSettingResponseBody {
	s.Model = v
	return s
}

func (s *GetChatappPhoneNumberSettingResponseBody) SetRequestId(v string) *GetChatappPhoneNumberSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatappPhoneNumberSettingResponseBody) SetSuccess(v bool) *GetChatappPhoneNumberSettingResponseBody {
	s.Success = &v
	return s
}

func (s *GetChatappPhoneNumberSettingResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChatappPhoneNumberSettingResponseBodyModel struct {
	// example:
	//
	// Y
	MarketingPaused *string `json:"MarketingPaused,omitempty" xml:"MarketingPaused,omitempty"`
	// example:
	//
	// 86138****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetChatappPhoneNumberSettingResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetChatappPhoneNumberSettingResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberSettingResponseBodyModel) GetMarketingPaused() *string {
	return s.MarketingPaused
}

func (s *GetChatappPhoneNumberSettingResponseBodyModel) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetChatappPhoneNumberSettingResponseBodyModel) SetMarketingPaused(v string) *GetChatappPhoneNumberSettingResponseBodyModel {
	s.MarketingPaused = &v
	return s
}

func (s *GetChatappPhoneNumberSettingResponseBodyModel) SetPhoneNumber(v string) *GetChatappPhoneNumberSettingResponseBodyModel {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappPhoneNumberSettingResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
