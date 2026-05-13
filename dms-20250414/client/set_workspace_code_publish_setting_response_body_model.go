// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkspaceCodePublishSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SetWorkspaceCodePublishSettingResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *SetWorkspaceCodePublishSettingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SetWorkspaceCodePublishSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetWorkspaceCodePublishSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetWorkspaceCodePublishSettingResponseBody
	GetSuccess() *bool
}

type SetWorkspaceCodePublishSettingResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Failed to publish，repo branch empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetWorkspaceCodePublishSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetWorkspaceCodePublishSettingResponseBody) GoString() string {
	return s.String()
}

func (s *SetWorkspaceCodePublishSettingResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SetWorkspaceCodePublishSettingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SetWorkspaceCodePublishSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetWorkspaceCodePublishSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetWorkspaceCodePublishSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetWorkspaceCodePublishSettingResponseBody) SetErrorCode(v string) *SetWorkspaceCodePublishSettingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetWorkspaceCodePublishSettingResponseBody) SetHttpStatusCode(v int32) *SetWorkspaceCodePublishSettingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SetWorkspaceCodePublishSettingResponseBody) SetMessage(v string) *SetWorkspaceCodePublishSettingResponseBody {
	s.Message = &v
	return s
}

func (s *SetWorkspaceCodePublishSettingResponseBody) SetRequestId(v string) *SetWorkspaceCodePublishSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWorkspaceCodePublishSettingResponseBody) SetSuccess(v bool) *SetWorkspaceCodePublishSettingResponseBody {
	s.Success = &v
	return s
}

func (s *SetWorkspaceCodePublishSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
