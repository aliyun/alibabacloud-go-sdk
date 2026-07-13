// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTeamResponseBody
	GetCode() *string
	SetData(v *UpdateTeamResponseBodyData) *UpdateTeamResponseBody
	GetData() *UpdateTeamResponseBodyData
	SetHttpStatusCode(v int32) *UpdateTeamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTeamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTeamResponseBody
	GetSuccess() *bool
}

type UpdateTeamResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateTeamResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTeamResponseBody) GetData() *UpdateTeamResponseBodyData {
	return s.Data
}

func (s *UpdateTeamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTeamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTeamResponseBody) SetCode(v string) *UpdateTeamResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTeamResponseBody) SetData(v *UpdateTeamResponseBodyData) *UpdateTeamResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTeamResponseBody) SetHttpStatusCode(v int32) *UpdateTeamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTeamResponseBody) SetMessage(v string) *UpdateTeamResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTeamResponseBody) SetRequestId(v string) *UpdateTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTeamResponseBody) SetSuccess(v bool) *UpdateTeamResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTeamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTeamResponseBodyData struct {
	AdminName   *string                                  `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	CreatedAt   *string                                  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LeaderName  *string                                  `json:"LeaderName,omitempty" xml:"LeaderName,omitempty"`
	Name        *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	TeamMembers []*UpdateTeamResponseBodyDataTeamMembers `json:"TeamMembers,omitempty" xml:"TeamMembers,omitempty" type:"Repeated"`
	UpdatedAt   *string                                  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	WorkerNames []*string                                `json:"WorkerNames,omitempty" xml:"WorkerNames,omitempty" type:"Repeated"`
}

func (s UpdateTeamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponseBodyData) GetAdminName() *string {
	return s.AdminName
}

func (s *UpdateTeamResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateTeamResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateTeamResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateTeamResponseBodyData) GetLeaderName() *string {
	return s.LeaderName
}

func (s *UpdateTeamResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateTeamResponseBodyData) GetTeamMembers() []*UpdateTeamResponseBodyDataTeamMembers {
	return s.TeamMembers
}

func (s *UpdateTeamResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateTeamResponseBodyData) GetWorkerNames() []*string {
	return s.WorkerNames
}

func (s *UpdateTeamResponseBodyData) SetAdminName(v string) *UpdateTeamResponseBodyData {
	s.AdminName = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetCreatedAt(v string) *UpdateTeamResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetDescription(v string) *UpdateTeamResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetInstanceId(v string) *UpdateTeamResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetLeaderName(v string) *UpdateTeamResponseBodyData {
	s.LeaderName = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetName(v string) *UpdateTeamResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetTeamMembers(v []*UpdateTeamResponseBodyDataTeamMembers) *UpdateTeamResponseBodyData {
	s.TeamMembers = v
	return s
}

func (s *UpdateTeamResponseBodyData) SetUpdatedAt(v string) *UpdateTeamResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateTeamResponseBodyData) SetWorkerNames(v []*string) *UpdateTeamResponseBodyData {
	s.WorkerNames = v
	return s
}

func (s *UpdateTeamResponseBodyData) Validate() error {
	if s.TeamMembers != nil {
		for _, item := range s.TeamMembers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTeamResponseBodyDataTeamMembers struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateTeamResponseBodyDataTeamMembers) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamResponseBodyDataTeamMembers) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponseBodyDataTeamMembers) GetName() *string {
	return s.Name
}

func (s *UpdateTeamResponseBodyDataTeamMembers) SetName(v string) *UpdateTeamResponseBodyDataTeamMembers {
	s.Name = &v
	return s
}

func (s *UpdateTeamResponseBodyDataTeamMembers) Validate() error {
	return dara.Validate(s)
}
