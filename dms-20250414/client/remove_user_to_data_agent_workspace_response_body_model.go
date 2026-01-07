// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserToDataAgentWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RemoveUserToDataAgentWorkspaceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RemoveUserToDataAgentWorkspaceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RemoveUserToDataAgentWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveUserToDataAgentWorkspaceResponseBody
	GetSuccess() *bool
}

type RemoveUserToDataAgentWorkspaceResponseBody struct {
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

func (s RemoveUserToDataAgentWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserToDataAgentWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserToDataAgentWorkspaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RemoveUserToDataAgentWorkspaceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RemoveUserToDataAgentWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserToDataAgentWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveUserToDataAgentWorkspaceResponseBody) SetErrorCode(v string) *RemoveUserToDataAgentWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceResponseBody) SetErrorMessage(v string) *RemoveUserToDataAgentWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceResponseBody) SetRequestId(v string) *RemoveUserToDataAgentWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceResponseBody) SetSuccess(v bool) *RemoveUserToDataAgentWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
