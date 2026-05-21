// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoloWebLoginSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetHoloWebLoginSettingResponseBodyData) *GetHoloWebLoginSettingResponseBody
	GetData() *GetHoloWebLoginSettingResponseBodyData
	SetErrorCode(v string) *GetHoloWebLoginSettingResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetHoloWebLoginSettingResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GetHoloWebLoginSettingResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *GetHoloWebLoginSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHoloWebLoginSettingResponseBody
	GetSuccess() *bool
}

type GetHoloWebLoginSettingResponseBody struct {
	Data *GetHoloWebLoginSettingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHoloWebLoginSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHoloWebLoginSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetHoloWebLoginSettingResponseBody) GetData() *GetHoloWebLoginSettingResponseBodyData {
	return s.Data
}

func (s *GetHoloWebLoginSettingResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetHoloWebLoginSettingResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetHoloWebLoginSettingResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetHoloWebLoginSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHoloWebLoginSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHoloWebLoginSettingResponseBody) SetData(v *GetHoloWebLoginSettingResponseBodyData) *GetHoloWebLoginSettingResponseBody {
	s.Data = v
	return s
}

func (s *GetHoloWebLoginSettingResponseBody) SetErrorCode(v string) *GetHoloWebLoginSettingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetHoloWebLoginSettingResponseBody) SetErrorMessage(v string) *GetHoloWebLoginSettingResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetHoloWebLoginSettingResponseBody) SetHttpStatusCode(v string) *GetHoloWebLoginSettingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHoloWebLoginSettingResponseBody) SetRequestId(v string) *GetHoloWebLoginSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHoloWebLoginSettingResponseBody) SetSuccess(v bool) *GetHoloWebLoginSettingResponseBody {
	s.Success = &v
	return s
}

func (s *GetHoloWebLoginSettingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHoloWebLoginSettingResponseBodyData struct {
	// example:
	//
	// true
	AllowExternalAccountsLogin *bool `json:"AllowExternalAccountsLogin,omitempty" xml:"AllowExternalAccountsLogin,omitempty"`
}

func (s GetHoloWebLoginSettingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHoloWebLoginSettingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHoloWebLoginSettingResponseBodyData) GetAllowExternalAccountsLogin() *bool {
	return s.AllowExternalAccountsLogin
}

func (s *GetHoloWebLoginSettingResponseBodyData) SetAllowExternalAccountsLogin(v bool) *GetHoloWebLoginSettingResponseBodyData {
	s.AllowExternalAccountsLogin = &v
	return s
}

func (s *GetHoloWebLoginSettingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
