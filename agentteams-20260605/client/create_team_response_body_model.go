// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTeamResponseBody
	GetCode() *string
	SetData(v *CreateTeamResponseBodyData) *CreateTeamResponseBody
	GetData() *CreateTeamResponseBodyData
	SetHttpStatusCode(v int32) *CreateTeamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTeamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTeamResponseBody
	GetSuccess() *bool
}

type CreateTeamResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateTeamResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTeamResponseBody) GetData() *CreateTeamResponseBodyData {
	return s.Data
}

func (s *CreateTeamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTeamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTeamResponseBody) SetCode(v string) *CreateTeamResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTeamResponseBody) SetData(v *CreateTeamResponseBodyData) *CreateTeamResponseBody {
	s.Data = v
	return s
}

func (s *CreateTeamResponseBody) SetHttpStatusCode(v int32) *CreateTeamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTeamResponseBody) SetMessage(v string) *CreateTeamResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTeamResponseBody) SetRequestId(v string) *CreateTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTeamResponseBody) SetSuccess(v bool) *CreateTeamResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTeamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTeamResponseBodyData struct {
	AdminName   *string                                  `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	Description *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	TeamMembers []*CreateTeamResponseBodyDataTeamMembers `json:"TeamMembers,omitempty" xml:"TeamMembers,omitempty" type:"Repeated"`
}

func (s CreateTeamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTeamResponseBodyData) GetAdminName() *string {
	return s.AdminName
}

func (s *CreateTeamResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateTeamResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTeamResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateTeamResponseBodyData) GetTeamMembers() []*CreateTeamResponseBodyDataTeamMembers {
	return s.TeamMembers
}

func (s *CreateTeamResponseBodyData) SetAdminName(v string) *CreateTeamResponseBodyData {
	s.AdminName = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetDescription(v string) *CreateTeamResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetInstanceId(v string) *CreateTeamResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetName(v string) *CreateTeamResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTeamResponseBodyData) SetTeamMembers(v []*CreateTeamResponseBodyDataTeamMembers) *CreateTeamResponseBodyData {
	s.TeamMembers = v
	return s
}

func (s *CreateTeamResponseBodyData) Validate() error {
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

type CreateTeamResponseBodyDataTeamMembers struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateTeamResponseBodyDataTeamMembers) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamResponseBodyDataTeamMembers) GoString() string {
	return s.String()
}

func (s *CreateTeamResponseBodyDataTeamMembers) GetName() *string {
	return s.Name
}

func (s *CreateTeamResponseBodyDataTeamMembers) SetName(v string) *CreateTeamResponseBodyDataTeamMembers {
	s.Name = &v
	return s
}

func (s *CreateTeamResponseBodyDataTeamMembers) Validate() error {
	return dara.Validate(s)
}
