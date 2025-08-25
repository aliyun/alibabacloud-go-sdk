// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMpsSchedulerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CancelMpsSchedulerRequest
	GetAppId() *string
	SetType(v int32) *CancelMpsSchedulerRequest
	GetType() *int32
	SetUniqueIds(v string) *CancelMpsSchedulerRequest
	GetUniqueIds() *string
	SetWorkspaceId(v string) *CancelMpsSchedulerRequest
	GetWorkspaceId() *string
}

type CancelMpsSchedulerRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Type  *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	UniqueIds *string `json:"UniqueIds,omitempty" xml:"UniqueIds,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CancelMpsSchedulerRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelMpsSchedulerRequest) GoString() string {
	return s.String()
}

func (s *CancelMpsSchedulerRequest) GetAppId() *string {
	return s.AppId
}

func (s *CancelMpsSchedulerRequest) GetType() *int32 {
	return s.Type
}

func (s *CancelMpsSchedulerRequest) GetUniqueIds() *string {
	return s.UniqueIds
}

func (s *CancelMpsSchedulerRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CancelMpsSchedulerRequest) SetAppId(v string) *CancelMpsSchedulerRequest {
	s.AppId = &v
	return s
}

func (s *CancelMpsSchedulerRequest) SetType(v int32) *CancelMpsSchedulerRequest {
	s.Type = &v
	return s
}

func (s *CancelMpsSchedulerRequest) SetUniqueIds(v string) *CancelMpsSchedulerRequest {
	s.UniqueIds = &v
	return s
}

func (s *CancelMpsSchedulerRequest) SetWorkspaceId(v string) *CancelMpsSchedulerRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CancelMpsSchedulerRequest) Validate() error {
	return dara.Validate(s)
}
