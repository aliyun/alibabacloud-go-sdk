// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppRagRequest
	GetAppId() *string
	SetStatus(v bool) *UpdateMmAppRagRequest
	GetStatus() *bool
	SetWorkspaceId(v string) *UpdateMmAppRagRequest
	GetWorkspaceId() *string
}

type UpdateMmAppRagRequest struct {
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

func (s UpdateMmAppRagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppRagRequest) GetStatus() *bool {
	return s.Status
}

func (s *UpdateMmAppRagRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppRagRequest) SetAppId(v string) *UpdateMmAppRagRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppRagRequest) SetStatus(v bool) *UpdateMmAppRagRequest {
	s.Status = &v
	return s
}

func (s *UpdateMmAppRagRequest) SetWorkspaceId(v string) *UpdateMmAppRagRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppRagRequest) Validate() error {
	return dara.Validate(s)
}
