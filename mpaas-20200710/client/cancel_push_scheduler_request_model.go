// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPushSchedulerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CancelPushSchedulerRequest
	GetAppId() *string
	SetTenantId(v string) *CancelPushSchedulerRequest
	GetTenantId() *string
	SetType(v int32) *CancelPushSchedulerRequest
	GetType() *int32
	SetUniqueIds(v string) *CancelPushSchedulerRequest
	GetUniqueIds() *string
	SetWorkspaceId(v string) *CancelPushSchedulerRequest
	GetWorkspaceId() *string
}

type CancelPushSchedulerRequest struct {
	// This parameter is required.
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	UniqueIds *string `json:"UniqueIds,omitempty" xml:"UniqueIds,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CancelPushSchedulerRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPushSchedulerRequest) GoString() string {
	return s.String()
}

func (s *CancelPushSchedulerRequest) GetAppId() *string {
	return s.AppId
}

func (s *CancelPushSchedulerRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CancelPushSchedulerRequest) GetType() *int32 {
	return s.Type
}

func (s *CancelPushSchedulerRequest) GetUniqueIds() *string {
	return s.UniqueIds
}

func (s *CancelPushSchedulerRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CancelPushSchedulerRequest) SetAppId(v string) *CancelPushSchedulerRequest {
	s.AppId = &v
	return s
}

func (s *CancelPushSchedulerRequest) SetTenantId(v string) *CancelPushSchedulerRequest {
	s.TenantId = &v
	return s
}

func (s *CancelPushSchedulerRequest) SetType(v int32) *CancelPushSchedulerRequest {
	s.Type = &v
	return s
}

func (s *CancelPushSchedulerRequest) SetUniqueIds(v string) *CancelPushSchedulerRequest {
	s.UniqueIds = &v
	return s
}

func (s *CancelPushSchedulerRequest) SetWorkspaceId(v string) *CancelPushSchedulerRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CancelPushSchedulerRequest) Validate() error {
	return dara.Validate(s)
}
