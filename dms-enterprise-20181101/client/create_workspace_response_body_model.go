// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkspaceResponseBody
	GetSuccess() *bool
	SetWorkspaceId(v int64) *CreateWorkspaceResponseBody
	GetWorkspaceId() *int64
}

type CreateWorkspaceResponseBody struct {
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 4AC23904-55DE-550B-9676-E8946F07****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkspaceResponseBody) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetSuccess(v bool) *CreateWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspaceId(v int64) *CreateWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
