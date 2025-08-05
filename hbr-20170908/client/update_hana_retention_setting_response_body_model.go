// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaRetentionSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHanaRetentionSettingResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateHanaRetentionSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHanaRetentionSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHanaRetentionSettingResponseBody
	GetSuccess() *bool
}

type UpdateHanaRetentionSettingResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7BEE26EB-8EE3-57A0-A9DE-5FD700165DE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHanaRetentionSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaRetentionSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHanaRetentionSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHanaRetentionSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHanaRetentionSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHanaRetentionSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHanaRetentionSettingResponseBody) SetCode(v string) *UpdateHanaRetentionSettingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHanaRetentionSettingResponseBody) SetMessage(v string) *UpdateHanaRetentionSettingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHanaRetentionSettingResponseBody) SetRequestId(v string) *UpdateHanaRetentionSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHanaRetentionSettingResponseBody) SetSuccess(v bool) *UpdateHanaRetentionSettingResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHanaRetentionSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
