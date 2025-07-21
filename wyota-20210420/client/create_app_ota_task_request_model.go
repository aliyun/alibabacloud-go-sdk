// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppOtaTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersionUid(v string) *CreateAppOtaTaskRequest
	GetAppVersionUid() *string
	SetChannel(v string) *CreateAppOtaTaskRequest
	GetChannel() *string
	SetClientIdList(v []*string) *CreateAppOtaTaskRequest
	GetClientIdList() []*string
	SetClientType(v int32) *CreateAppOtaTaskRequest
	GetClientType() *int32
	SetCreator(v string) *CreateAppOtaTaskRequest
	GetCreator() *string
	SetDescription(v string) *CreateAppOtaTaskRequest
	GetDescription() *string
	SetForceUpgrade(v int32) *CreateAppOtaTaskRequest
	GetForceUpgrade() *int32
	SetLabel(v string) *CreateAppOtaTaskRequest
	GetLabel() *string
	SetName(v string) *CreateAppOtaTaskRequest
	GetName() *string
	SetProject(v string) *CreateAppOtaTaskRequest
	GetProject() *string
	SetRegions(v []*string) *CreateAppOtaTaskRequest
	GetRegions() []*string
	SetStatus(v int32) *CreateAppOtaTaskRequest
	GetStatus() *int32
	SetTaskType(v int32) *CreateAppOtaTaskRequest
	GetTaskType() *int32
	SetTenantId(v string) *CreateAppOtaTaskRequest
	GetTenantId() *string
	SetTenantIdList(v []*string) *CreateAppOtaTaskRequest
	GetTenantIdList() []*string
}

type CreateAppOtaTaskRequest struct {
	AppVersionUid *string   `json:"AppVersionUid,omitempty" xml:"AppVersionUid,omitempty"`
	Channel       *string   `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ClientIdList  []*string `json:"ClientIdList,omitempty" xml:"ClientIdList,omitempty" type:"Repeated"`
	ClientType    *int32    `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Creator       *string   `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description   *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ForceUpgrade  *int32    `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	Label         *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Project       *string   `json:"Project,omitempty" xml:"Project,omitempty"`
	Regions       []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	Status        *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskType      *int32    `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TenantId      *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantIdList  []*string `json:"TenantIdList,omitempty" xml:"TenantIdList,omitempty" type:"Repeated"`
}

func (s CreateAppOtaTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAppOtaTaskRequest) GetAppVersionUid() *string {
	return s.AppVersionUid
}

func (s *CreateAppOtaTaskRequest) GetChannel() *string {
	return s.Channel
}

func (s *CreateAppOtaTaskRequest) GetClientIdList() []*string {
	return s.ClientIdList
}

func (s *CreateAppOtaTaskRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *CreateAppOtaTaskRequest) GetCreator() *string {
	return s.Creator
}

func (s *CreateAppOtaTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppOtaTaskRequest) GetForceUpgrade() *int32 {
	return s.ForceUpgrade
}

func (s *CreateAppOtaTaskRequest) GetLabel() *string {
	return s.Label
}

func (s *CreateAppOtaTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateAppOtaTaskRequest) GetProject() *string {
	return s.Project
}

func (s *CreateAppOtaTaskRequest) GetRegions() []*string {
	return s.Regions
}

func (s *CreateAppOtaTaskRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateAppOtaTaskRequest) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateAppOtaTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAppOtaTaskRequest) GetTenantIdList() []*string {
	return s.TenantIdList
}

func (s *CreateAppOtaTaskRequest) SetAppVersionUid(v string) *CreateAppOtaTaskRequest {
	s.AppVersionUid = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetChannel(v string) *CreateAppOtaTaskRequest {
	s.Channel = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetClientIdList(v []*string) *CreateAppOtaTaskRequest {
	s.ClientIdList = v
	return s
}

func (s *CreateAppOtaTaskRequest) SetClientType(v int32) *CreateAppOtaTaskRequest {
	s.ClientType = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetCreator(v string) *CreateAppOtaTaskRequest {
	s.Creator = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetDescription(v string) *CreateAppOtaTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetForceUpgrade(v int32) *CreateAppOtaTaskRequest {
	s.ForceUpgrade = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetLabel(v string) *CreateAppOtaTaskRequest {
	s.Label = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetName(v string) *CreateAppOtaTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetProject(v string) *CreateAppOtaTaskRequest {
	s.Project = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetRegions(v []*string) *CreateAppOtaTaskRequest {
	s.Regions = v
	return s
}

func (s *CreateAppOtaTaskRequest) SetStatus(v int32) *CreateAppOtaTaskRequest {
	s.Status = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetTaskType(v int32) *CreateAppOtaTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetTenantId(v string) *CreateAppOtaTaskRequest {
	s.TenantId = &v
	return s
}

func (s *CreateAppOtaTaskRequest) SetTenantIdList(v []*string) *CreateAppOtaTaskRequest {
	s.TenantIdList = v
	return s
}

func (s *CreateAppOtaTaskRequest) Validate() error {
	return dara.Validate(s)
}
