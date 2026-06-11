// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppTransitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppTransitionRequest
	GetAppId() *string
	SetStatus(v bool) *UpdateMmAppTransitionRequest
	GetStatus() *bool
	SetWorkspaceId(v string) *UpdateMmAppTransitionRequest
	GetWorkspaceId() *string
}

type UpdateMmAppTransitionRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMmAppTransitionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppTransitionRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppTransitionRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppTransitionRequest) GetStatus() *bool {
	return s.Status
}

func (s *UpdateMmAppTransitionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppTransitionRequest) SetAppId(v string) *UpdateMmAppTransitionRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppTransitionRequest) SetStatus(v bool) *UpdateMmAppTransitionRequest {
	s.Status = &v
	return s
}

func (s *UpdateMmAppTransitionRequest) SetWorkspaceId(v string) *UpdateMmAppTransitionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppTransitionRequest) Validate() error {
	return dara.Validate(s)
}
