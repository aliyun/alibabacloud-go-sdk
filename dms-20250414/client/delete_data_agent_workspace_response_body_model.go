// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteDataAgentWorkspaceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataAgentWorkspaceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataAgentWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataAgentWorkspaceResponseBody
	GetSuccess() *bool
}

type DeleteDataAgentWorkspaceResponseBody struct {
	// example:
	//
	// DMS-DA-40114
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// E0D2-*****-A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataAgentWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentWorkspaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataAgentWorkspaceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataAgentWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataAgentWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataAgentWorkspaceResponseBody) SetErrorCode(v string) *DeleteDataAgentWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataAgentWorkspaceResponseBody) SetErrorMessage(v string) *DeleteDataAgentWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataAgentWorkspaceResponseBody) SetRequestId(v string) *DeleteDataAgentWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataAgentWorkspaceResponseBody) SetSuccess(v bool) *DeleteDataAgentWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataAgentWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
