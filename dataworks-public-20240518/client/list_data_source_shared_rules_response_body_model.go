// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceSharedRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceSharedRules(v []*ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) *ListDataSourceSharedRulesResponseBody
	GetDataSourceSharedRules() []*ListDataSourceSharedRulesResponseBodyDataSourceSharedRules
	SetRequestId(v string) *ListDataSourceSharedRulesResponseBody
	GetRequestId() *string
}

type ListDataSourceSharedRulesResponseBody struct {
	// The sharing rules of the data source.
	DataSourceSharedRules []*ListDataSourceSharedRulesResponseBodyDataSourceSharedRules `json:"DataSourceSharedRules,omitempty" xml:"DataSourceSharedRules,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourceSharedRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceSharedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesResponseBody) GetDataSourceSharedRules() []*ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	return s.DataSourceSharedRules
}

func (s *ListDataSourceSharedRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceSharedRulesResponseBody) SetDataSourceSharedRules(v []*ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) *ListDataSourceSharedRulesResponseBody {
	s.DataSourceSharedRules = v
	return s
}

func (s *ListDataSourceSharedRulesResponseBody) SetRequestId(v string) *ListDataSourceSharedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataSourceSharedRulesResponseBodyDataSourceSharedRules struct {
	// The time when the rule was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724379762000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who creates the rule.
	//
	// example:
	//
	// 1
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The data source ID. You can call the [ListDataSources](https://help.aliyun.com/document_detail/211431.html) operation to query the ID.
	//
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The environment to which the target data source belongs. The values are as follows:
	//
	// - Dev: the development environment.
	//
	// - Prod: the production environment.
	//
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data source in the destination workspace.
	//
	// example:
	//
	// targetProject.datasource
	SharedDataSourceName *string `json:"SharedDataSourceName,omitempty" xml:"SharedDataSourceName,omitempty"`
	// The user in the workspace to which the data source is shared. If the data source is shared to the entire workspace, this parameter is left empty.
	//
	// example:
	//
	// 1
	SharedUser *string `json:"SharedUser,omitempty" xml:"SharedUser,omitempty"`
	// The ID of the workspace with which the data source is associated.
	//
	// example:
	//
	// 1
	SourceProjectId *int64 `json:"SourceProjectId,omitempty" xml:"SourceProjectId,omitempty"`
	// The ID of the workspace to which the data source is shared.
	//
	// example:
	//
	// 1
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GetId() *int64 {
	return s.Id
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GetSharedDataSourceName() *string {
	return s.SharedDataSourceName
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GetSharedUser() *string {
	return s.SharedUser
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GetSourceProjectId() *int64 {
	return s.SourceProjectId
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetCreateTime(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetCreateUser(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.CreateUser = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetDataSourceId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.DataSourceId = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetEnvType(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.EnvType = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.Id = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetSharedDataSourceName(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.SharedDataSourceName = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetSharedUser(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.SharedUser = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetSourceProjectId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.SourceProjectId = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetTargetProjectId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.TargetProjectId = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) Validate() error {
	return dara.Validate(s)
}
