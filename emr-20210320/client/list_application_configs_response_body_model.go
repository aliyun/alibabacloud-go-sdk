// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v []*ListApplicationConfigsResponseBodyApplicationConfigs) *ListApplicationConfigsResponseBody
	GetApplicationConfigs() []*ListApplicationConfigsResponseBodyApplicationConfigs
	SetMaxResults(v int32) *ListApplicationConfigsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationConfigsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationConfigsResponseBody
	GetTotalCount() *int32
}

type ListApplicationConfigsResponseBody struct {
	// The application configurations.
	ApplicationConfigs []*ListApplicationConfigsResponseBodyApplicationConfigs `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number of the next page returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationConfigsResponseBody) GetApplicationConfigs() []*ListApplicationConfigsResponseBodyApplicationConfigs {
	return s.ApplicationConfigs
}

func (s *ListApplicationConfigsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationConfigsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationConfigsResponseBody) SetApplicationConfigs(v []*ListApplicationConfigsResponseBodyApplicationConfigs) *ListApplicationConfigsResponseBody {
	s.ApplicationConfigs = v
	return s
}

func (s *ListApplicationConfigsResponseBody) SetMaxResults(v int32) *ListApplicationConfigsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationConfigsResponseBody) SetNextToken(v string) *ListApplicationConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationConfigsResponseBody) SetRequestId(v string) *ListApplicationConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationConfigsResponseBody) SetTotalCount(v int32) *ListApplicationConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationConfigsResponseBodyApplicationConfigs struct {
	// The name of the application.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The status of the configuration value.
	//
	// example:
	//
	// EFFECT
	ConfigEffectState *string `json:"ConfigEffectState,omitempty" xml:"ConfigEffectState,omitempty"`
	// The name of the configuration file.
	//
	// example:
	//
	// hdfs-site.xml
	ConfigFileName *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	// The key of the configuration item.
	//
	// example:
	//
	// dfs.replication
	ConfigItemKey *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// 2
	ConfigItemValue *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1628589439114
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the configurations are custom.
	//
	// example:
	//
	// false
	Custom *bool `json:"Custom,omitempty" xml:"Custom,omitempty"`
	// The description.
	//
	// example:
	//
	// dfs.replication description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The initial value.
	//
	// example:
	//
	// 2
	InitValue *string `json:"InitValue,omitempty" xml:"InitValue,omitempty"`
	// The person who modified the configurations.
	//
	// example:
	//
	// 170906468757****
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// ng-d555335ced5****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// i-bp18y0ng3qqxog4mw****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1628589439114
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListApplicationConfigsResponseBodyApplicationConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationConfigsResponseBodyApplicationConfigs) GoString() string {
	return s.String()
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetConfigEffectState() *string {
	return s.ConfigEffectState
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetCustom() *bool {
	return s.Custom
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetInitValue() *string {
	return s.InitValue
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetModifier() *string {
	return s.Modifier
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetNodeId() *string {
	return s.NodeId
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetApplicationName(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetConfigEffectState(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.ConfigEffectState = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetConfigFileName(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.ConfigFileName = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetConfigItemKey(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.ConfigItemKey = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetConfigItemValue(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.ConfigItemValue = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetCreateTime(v int64) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetCustom(v bool) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.Custom = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetDescription(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.Description = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetInitValue(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.InitValue = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetModifier(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.Modifier = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetNodeGroupId(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.NodeGroupId = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetNodeId(v string) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.NodeId = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) SetUpdateTime(v int64) *ListApplicationConfigsResponseBodyApplicationConfigs {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationConfigsResponseBodyApplicationConfigs) Validate() error {
	return dara.Validate(s)
}
