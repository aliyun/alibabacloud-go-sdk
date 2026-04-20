// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppMemoryRequest
	GetAppId() *string
	SetStatus(v bool) *UpdateMmAppMemoryRequest
	GetStatus() *bool
	SetWorkspaceId(v string) *UpdateMmAppMemoryRequest
	GetWorkspaceId() *string
}

type UpdateMmAppMemoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_c617650fbd9e49e8916189e23c62
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-6uhm7nfev4k8pwcz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMmAppMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppMemoryRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppMemoryRequest) GetStatus() *bool {
	return s.Status
}

func (s *UpdateMmAppMemoryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppMemoryRequest) SetAppId(v string) *UpdateMmAppMemoryRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppMemoryRequest) SetStatus(v bool) *UpdateMmAppMemoryRequest {
	s.Status = &v
	return s
}

func (s *UpdateMmAppMemoryRequest) SetWorkspaceId(v string) *UpdateMmAppMemoryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppMemoryRequest) Validate() error {
	return dara.Validate(s)
}
