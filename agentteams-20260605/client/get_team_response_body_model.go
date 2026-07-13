// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTeamResponseBody
	GetCode() *string
	SetData(v *GetTeamResponseBodyData) *GetTeamResponseBody
	GetData() *GetTeamResponseBodyData
	SetHttpStatusCode(v int32) *GetTeamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTeamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTeamResponseBody
	GetSuccess() *bool
}

type GetTeamResponseBody struct {
	Code           *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetTeamResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBody) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTeamResponseBody) GetData() *GetTeamResponseBodyData {
	return s.Data
}

func (s *GetTeamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTeamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTeamResponseBody) SetCode(v string) *GetTeamResponseBody {
	s.Code = &v
	return s
}

func (s *GetTeamResponseBody) SetData(v *GetTeamResponseBodyData) *GetTeamResponseBody {
	s.Data = v
	return s
}

func (s *GetTeamResponseBody) SetHttpStatusCode(v int32) *GetTeamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTeamResponseBody) SetMessage(v string) *GetTeamResponseBody {
	s.Message = &v
	return s
}

func (s *GetTeamResponseBody) SetRequestId(v string) *GetTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTeamResponseBody) SetSuccess(v bool) *GetTeamResponseBody {
	s.Success = &v
	return s
}

func (s *GetTeamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTeamResponseBodyData struct {
	AdminName   *string                               `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	CreatedAt   *string                               `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description *string                               `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LeaderName  *string                               `json:"LeaderName,omitempty" xml:"LeaderName,omitempty"`
	Name        *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Rooms       []*GetTeamResponseBodyDataRooms       `json:"Rooms,omitempty" xml:"Rooms,omitempty" type:"Repeated"`
	Status      *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	TeamMembers []*GetTeamResponseBodyDataTeamMembers `json:"TeamMembers,omitempty" xml:"TeamMembers,omitempty" type:"Repeated"`
	UpdatedAt   *string                               `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	WorkerNames []*string                             `json:"WorkerNames,omitempty" xml:"WorkerNames,omitempty" type:"Repeated"`
}

func (s GetTeamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyData) GetAdminName() *string {
	return s.AdminName
}

func (s *GetTeamResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetTeamResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetTeamResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTeamResponseBodyData) GetLeaderName() *string {
	return s.LeaderName
}

func (s *GetTeamResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTeamResponseBodyData) GetRooms() []*GetTeamResponseBodyDataRooms {
	return s.Rooms
}

func (s *GetTeamResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTeamResponseBodyData) GetTeamMembers() []*GetTeamResponseBodyDataTeamMembers {
	return s.TeamMembers
}

func (s *GetTeamResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetTeamResponseBodyData) GetWorkerNames() []*string {
	return s.WorkerNames
}

func (s *GetTeamResponseBodyData) SetAdminName(v string) *GetTeamResponseBodyData {
	s.AdminName = &v
	return s
}

func (s *GetTeamResponseBodyData) SetCreatedAt(v string) *GetTeamResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetTeamResponseBodyData) SetDescription(v string) *GetTeamResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTeamResponseBodyData) SetInstanceId(v string) *GetTeamResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetTeamResponseBodyData) SetLeaderName(v string) *GetTeamResponseBodyData {
	s.LeaderName = &v
	return s
}

func (s *GetTeamResponseBodyData) SetName(v string) *GetTeamResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTeamResponseBodyData) SetRooms(v []*GetTeamResponseBodyDataRooms) *GetTeamResponseBodyData {
	s.Rooms = v
	return s
}

func (s *GetTeamResponseBodyData) SetStatus(v string) *GetTeamResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTeamResponseBodyData) SetTeamMembers(v []*GetTeamResponseBodyDataTeamMembers) *GetTeamResponseBodyData {
	s.TeamMembers = v
	return s
}

func (s *GetTeamResponseBodyData) SetUpdatedAt(v string) *GetTeamResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetTeamResponseBodyData) SetWorkerNames(v []*string) *GetTeamResponseBodyData {
	s.WorkerNames = v
	return s
}

func (s *GetTeamResponseBodyData) Validate() error {
	if s.Rooms != nil {
		for _, item := range s.Rooms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type GetTeamResponseBodyDataRooms struct {
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTeamResponseBodyDataRooms) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBodyDataRooms) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyDataRooms) GetRoomId() *string {
	return s.RoomId
}

func (s *GetTeamResponseBodyDataRooms) GetType() *string {
	return s.Type
}

func (s *GetTeamResponseBodyDataRooms) SetRoomId(v string) *GetTeamResponseBodyDataRooms {
	s.RoomId = &v
	return s
}

func (s *GetTeamResponseBodyDataRooms) SetType(v string) *GetTeamResponseBodyDataRooms {
	s.Type = &v
	return s
}

func (s *GetTeamResponseBodyDataRooms) Validate() error {
	return dara.Validate(s)
}

type GetTeamResponseBodyDataTeamMembers struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTeamResponseBodyDataTeamMembers) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBodyDataTeamMembers) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyDataTeamMembers) GetName() *string {
	return s.Name
}

func (s *GetTeamResponseBodyDataTeamMembers) SetName(v string) *GetTeamResponseBodyDataTeamMembers {
	s.Name = &v
	return s
}

func (s *GetTeamResponseBodyDataTeamMembers) Validate() error {
	return dara.Validate(s)
}
