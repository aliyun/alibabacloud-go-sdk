// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTeamsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTeamsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListTeamsResponseBodyItems) *ListTeamsResponseBody
	GetItems() []*ListTeamsResponseBodyItems
	SetMaxResults(v int32) *ListTeamsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTeamsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTeamsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTeamsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTeamsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListTeamsResponseBody
	GetTotalCount() *int64
}

type ListTeamsResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*ListTeamsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults     *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTeamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTeamsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTeamsResponseBody) GetItems() []*ListTeamsResponseBodyItems {
	return s.Items
}

func (s *ListTeamsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTeamsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTeamsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTeamsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTeamsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTeamsResponseBody) SetCode(v string) *ListTeamsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTeamsResponseBody) SetHttpStatusCode(v int32) *ListTeamsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTeamsResponseBody) SetItems(v []*ListTeamsResponseBodyItems) *ListTeamsResponseBody {
	s.Items = v
	return s
}

func (s *ListTeamsResponseBody) SetMaxResults(v int32) *ListTeamsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTeamsResponseBody) SetMessage(v string) *ListTeamsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTeamsResponseBody) SetNextToken(v string) *ListTeamsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTeamsResponseBody) SetRequestId(v string) *ListTeamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTeamsResponseBody) SetSuccess(v bool) *ListTeamsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTeamsResponseBody) SetTotalCount(v int64) *ListTeamsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTeamsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTeamsResponseBodyItems struct {
	AdminName   *string                                  `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	CreatedAt   *string                                  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LeaderName  *string                                  `json:"LeaderName,omitempty" xml:"LeaderName,omitempty"`
	Name        *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TeamMembers []*ListTeamsResponseBodyItemsTeamMembers `json:"TeamMembers,omitempty" xml:"TeamMembers,omitempty" type:"Repeated"`
	WorkerNames []*string                                `json:"WorkerNames,omitempty" xml:"WorkerNames,omitempty" type:"Repeated"`
}

func (s ListTeamsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBodyItems) GetAdminName() *string {
	return s.AdminName
}

func (s *ListTeamsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListTeamsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListTeamsResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTeamsResponseBodyItems) GetLeaderName() *string {
	return s.LeaderName
}

func (s *ListTeamsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListTeamsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListTeamsResponseBodyItems) GetTeamMembers() []*ListTeamsResponseBodyItemsTeamMembers {
	return s.TeamMembers
}

func (s *ListTeamsResponseBodyItems) GetWorkerNames() []*string {
	return s.WorkerNames
}

func (s *ListTeamsResponseBodyItems) SetAdminName(v string) *ListTeamsResponseBodyItems {
	s.AdminName = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetCreatedAt(v string) *ListTeamsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetDescription(v string) *ListTeamsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetInstanceId(v string) *ListTeamsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetLeaderName(v string) *ListTeamsResponseBodyItems {
	s.LeaderName = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetName(v string) *ListTeamsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetStatus(v string) *ListTeamsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetTeamMembers(v []*ListTeamsResponseBodyItemsTeamMembers) *ListTeamsResponseBodyItems {
	s.TeamMembers = v
	return s
}

func (s *ListTeamsResponseBodyItems) SetWorkerNames(v []*string) *ListTeamsResponseBodyItems {
	s.WorkerNames = v
	return s
}

func (s *ListTeamsResponseBodyItems) Validate() error {
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

type ListTeamsResponseBodyItemsTeamMembers struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListTeamsResponseBodyItemsTeamMembers) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBodyItemsTeamMembers) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBodyItemsTeamMembers) GetName() *string {
	return s.Name
}

func (s *ListTeamsResponseBodyItemsTeamMembers) SetName(v string) *ListTeamsResponseBodyItemsTeamMembers {
	s.Name = &v
	return s
}

func (s *ListTeamsResponseBodyItemsTeamMembers) Validate() error {
	return dara.Validate(s)
}
