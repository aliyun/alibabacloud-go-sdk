// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListModelTemplatesResponseBodyData) *ListModelTemplatesResponseBody
	GetData() []*ListModelTemplatesResponseBodyData
	SetPageNumber(v int32) *ListModelTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListModelTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListModelTemplatesResponseBody
	GetTotalCount() *int32
}

type ListModelTemplatesResponseBody struct {
	// The list of returned data objects.
	Data []*ListModelTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The current page number of the query results.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of query results per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of query results.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelTemplatesResponseBody) GetData() []*ListModelTemplatesResponseBodyData {
	return s.Data
}

func (s *ListModelTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModelTemplatesResponseBody) SetData(v []*ListModelTemplatesResponseBodyData) *ListModelTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListModelTemplatesResponseBody) SetPageNumber(v int32) *ListModelTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModelTemplatesResponseBody) SetPageSize(v int32) *ListModelTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModelTemplatesResponseBody) SetRequestId(v string) *ListModelTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelTemplatesResponseBody) SetTotalCount(v int32) *ListModelTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelTemplatesResponseBodyData struct {
	// The Agent platform (such as ENTERPRISE or ENTERPRISE_JVS).
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The Agent provider name.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The model group configuration JSON object.
	//
	// example:
	//
	// {
	//
	// 	"defaults": {
	//
	// 		"model": {
	//
	// 			"primary": "bailian/qwen3.5-plus"
	//
	// 		}
	//
	// 	}
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The template group description.
	//
	// example:
	//
	// Test model group
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether models are configured in the group.
	//
	// example:
	//
	// true
	HasModel *bool `json:"HasModel,omitempty" xml:"HasModel,omitempty"`
	// The number of models in the model group, including referenced system provider models.
	ModelCount *int32 `json:"ModelCount,omitempty" xml:"ModelCount,omitempty"`
	// The model group ID.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// The template group name.
	//
	// example:
	//
	// model-template-001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The group authorization scope. Valid values:
	//
	// - ALL_USER: all users.
	//
	// - USER_MIXED: user mixed (a mix of user groups and users, only for Common groups).
	RefScope *string `json:"RefScope,omitempty" xml:"RefScope,omitempty"`
	// The number of authorized users in the group. Returned only when ListModelTemplates is called with refScope=USER_MIXED. Otherwise null.
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// The number of authorized user groups in the group. Returned only when ListModelTemplates is called with refScope=USER_MIXED. Otherwise null.
	UserGroupCount *int32 `json:"UserGroupCount,omitempty" xml:"UserGroupCount,omitempty"`
}

func (s ListModelTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelTemplatesResponseBodyData) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListModelTemplatesResponseBodyData) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListModelTemplatesResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *ListModelTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListModelTemplatesResponseBodyData) GetHasModel() *bool {
	return s.HasModel
}

func (s *ListModelTemplatesResponseBodyData) GetModelCount() *int32 {
	return s.ModelCount
}

func (s *ListModelTemplatesResponseBodyData) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListModelTemplatesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListModelTemplatesResponseBodyData) GetRefScope() *string {
	return s.RefScope
}

func (s *ListModelTemplatesResponseBodyData) GetUserCount() *int32 {
	return s.UserCount
}

func (s *ListModelTemplatesResponseBodyData) GetUserGroupCount() *int32 {
	return s.UserGroupCount
}

func (s *ListModelTemplatesResponseBodyData) SetAgentPlatform(v string) *ListModelTemplatesResponseBodyData {
	s.AgentPlatform = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetAgentProvider(v string) *ListModelTemplatesResponseBodyData {
	s.AgentProvider = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetConfig(v string) *ListModelTemplatesResponseBodyData {
	s.Config = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetDescription(v string) *ListModelTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetHasModel(v bool) *ListModelTemplatesResponseBodyData {
	s.HasModel = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetModelCount(v int32) *ListModelTemplatesResponseBodyData {
	s.ModelCount = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetModelTemplateId(v string) *ListModelTemplatesResponseBodyData {
	s.ModelTemplateId = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetName(v string) *ListModelTemplatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetRefScope(v string) *ListModelTemplatesResponseBodyData {
	s.RefScope = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetUserCount(v int32) *ListModelTemplatesResponseBodyData {
	s.UserCount = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetUserGroupCount(v int32) *ListModelTemplatesResponseBodyData {
	s.UserGroupCount = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
