// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListWorkersResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListWorkersResponseBodyItems) *ListWorkersResponseBody
	GetItems() []*ListWorkersResponseBodyItems
	SetMaxResults(v int32) *ListWorkersResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListWorkersResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWorkersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListWorkersResponseBody
	GetTotalCount() *int64
}

type ListWorkersResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*ListWorkersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults     *int32                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWorkersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkersResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWorkersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListWorkersResponseBody) GetItems() []*ListWorkersResponseBodyItems {
	return s.Items
}

func (s *ListWorkersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWorkersResponseBody) SetCode(v string) *ListWorkersResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkersResponseBody) SetHttpStatusCode(v int32) *ListWorkersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWorkersResponseBody) SetItems(v []*ListWorkersResponseBodyItems) *ListWorkersResponseBody {
	s.Items = v
	return s
}

func (s *ListWorkersResponseBody) SetMaxResults(v int32) *ListWorkersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkersResponseBody) SetMessage(v string) *ListWorkersResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkersResponseBody) SetNextToken(v string) *ListWorkersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkersResponseBody) SetRequestId(v string) *ListWorkersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkersResponseBody) SetSuccess(v bool) *ListWorkersResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkersResponseBody) SetTotalCount(v int64) *ListWorkersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkersResponseBody) Validate() error {
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

type ListWorkersResponseBodyItems struct {
	AgentType   *string                               `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	DeployType  *string                               `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Groups      []*ListWorkersResponseBodyItemsGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	InstanceId  *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Template    *ListWorkersResponseBodyItemsTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	VersionCode *string                               `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s ListWorkersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkersResponseBodyItems) GetAgentType() *string {
	return s.AgentType
}

func (s *ListWorkersResponseBodyItems) GetDeployType() *string {
	return s.DeployType
}

func (s *ListWorkersResponseBodyItems) GetGroups() []*ListWorkersResponseBodyItemsGroups {
	return s.Groups
}

func (s *ListWorkersResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkersResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListWorkersResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListWorkersResponseBodyItems) GetTemplate() *ListWorkersResponseBodyItemsTemplate {
	return s.Template
}

func (s *ListWorkersResponseBodyItems) GetVersionCode() *string {
	return s.VersionCode
}

func (s *ListWorkersResponseBodyItems) SetAgentType(v string) *ListWorkersResponseBodyItems {
	s.AgentType = &v
	return s
}

func (s *ListWorkersResponseBodyItems) SetDeployType(v string) *ListWorkersResponseBodyItems {
	s.DeployType = &v
	return s
}

func (s *ListWorkersResponseBodyItems) SetGroups(v []*ListWorkersResponseBodyItemsGroups) *ListWorkersResponseBodyItems {
	s.Groups = v
	return s
}

func (s *ListWorkersResponseBodyItems) SetInstanceId(v string) *ListWorkersResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListWorkersResponseBodyItems) SetName(v string) *ListWorkersResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListWorkersResponseBodyItems) SetStatus(v string) *ListWorkersResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListWorkersResponseBodyItems) SetTemplate(v *ListWorkersResponseBodyItemsTemplate) *ListWorkersResponseBodyItems {
	s.Template = v
	return s
}

func (s *ListWorkersResponseBodyItems) SetVersionCode(v string) *ListWorkersResponseBodyItems {
	s.VersionCode = &v
	return s
}

func (s *ListWorkersResponseBodyItems) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkersResponseBodyItemsGroups struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkersResponseBodyItemsGroups) String() string {
	return dara.Prettify(s)
}

func (s ListWorkersResponseBodyItemsGroups) GoString() string {
	return s.String()
}

func (s *ListWorkersResponseBodyItemsGroups) GetName() *string {
	return s.Name
}

func (s *ListWorkersResponseBodyItemsGroups) GetRole() *string {
	return s.Role
}

func (s *ListWorkersResponseBodyItemsGroups) GetType() *string {
	return s.Type
}

func (s *ListWorkersResponseBodyItemsGroups) SetName(v string) *ListWorkersResponseBodyItemsGroups {
	s.Name = &v
	return s
}

func (s *ListWorkersResponseBodyItemsGroups) SetRole(v string) *ListWorkersResponseBodyItemsGroups {
	s.Role = &v
	return s
}

func (s *ListWorkersResponseBodyItemsGroups) SetType(v string) *ListWorkersResponseBodyItemsGroups {
	s.Type = &v
	return s
}

func (s *ListWorkersResponseBodyItemsGroups) Validate() error {
	return dara.Validate(s)
}

type ListWorkersResponseBodyItemsTemplate struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListWorkersResponseBodyItemsTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListWorkersResponseBodyItemsTemplate) GoString() string {
	return s.String()
}

func (s *ListWorkersResponseBodyItemsTemplate) GetLabel() *string {
	return s.Label
}

func (s *ListWorkersResponseBodyItemsTemplate) GetName() *string {
	return s.Name
}

func (s *ListWorkersResponseBodyItemsTemplate) GetVersion() *string {
	return s.Version
}

func (s *ListWorkersResponseBodyItemsTemplate) SetLabel(v string) *ListWorkersResponseBodyItemsTemplate {
	s.Label = &v
	return s
}

func (s *ListWorkersResponseBodyItemsTemplate) SetName(v string) *ListWorkersResponseBodyItemsTemplate {
	s.Name = &v
	return s
}

func (s *ListWorkersResponseBodyItemsTemplate) SetVersion(v string) *ListWorkersResponseBodyItemsTemplate {
	s.Version = &v
	return s
}

func (s *ListWorkersResponseBodyItemsTemplate) Validate() error {
	return dara.Validate(s)
}
