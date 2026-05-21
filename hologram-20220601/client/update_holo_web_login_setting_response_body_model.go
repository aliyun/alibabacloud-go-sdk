// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoloWebLoginSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateHoloWebLoginSettingResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateHoloWebLoginSettingResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateHoloWebLoginSettingResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UpdateHoloWebLoginSettingResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateHoloWebLoginSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHoloWebLoginSettingResponseBody
	GetSuccess() *bool
}

type UpdateHoloWebLoginSettingResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CB13FFDD-2DF8-5396-A848-2D6A31245B6D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHoloWebLoginSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoloWebLoginSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHoloWebLoginSettingResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateHoloWebLoginSettingResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateHoloWebLoginSettingResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateHoloWebLoginSettingResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateHoloWebLoginSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHoloWebLoginSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHoloWebLoginSettingResponseBody) SetData(v bool) *UpdateHoloWebLoginSettingResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateHoloWebLoginSettingResponseBody) SetErrorCode(v string) *UpdateHoloWebLoginSettingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateHoloWebLoginSettingResponseBody) SetErrorMessage(v string) *UpdateHoloWebLoginSettingResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateHoloWebLoginSettingResponseBody) SetHttpStatusCode(v string) *UpdateHoloWebLoginSettingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateHoloWebLoginSettingResponseBody) SetRequestId(v string) *UpdateHoloWebLoginSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHoloWebLoginSettingResponseBody) SetSuccess(v bool) *UpdateHoloWebLoginSettingResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHoloWebLoginSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
