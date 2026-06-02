// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteWorkspaceCodeResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DeleteWorkspaceCodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteWorkspaceCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWorkspaceCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkspaceCodeResponseBody
	GetSuccess() *bool
}

type DeleteWorkspaceCodeResponseBody struct {
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// This record is being collected, please wait for a moment.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWorkspaceCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceCodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteWorkspaceCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteWorkspaceCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWorkspaceCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkspaceCodeResponseBody) SetErrorCode(v string) *DeleteWorkspaceCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteWorkspaceCodeResponseBody) SetHttpStatusCode(v int32) *DeleteWorkspaceCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteWorkspaceCodeResponseBody) SetMessage(v string) *DeleteWorkspaceCodeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWorkspaceCodeResponseBody) SetRequestId(v string) *DeleteWorkspaceCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceCodeResponseBody) SetSuccess(v bool) *DeleteWorkspaceCodeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkspaceCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
