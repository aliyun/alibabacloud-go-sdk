// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetClusterResponseBody
	GetRequestId() *string
	SetResult(v *GetClusterResponseBodyResult) *GetClusterResponseBody
	GetResult() *GetClusterResponseBodyResult
}

type GetClusterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The clusters.
	Result *GetClusterResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterResponseBody) GetResult() *GetClusterResponseBodyResult {
	return s.Result
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) SetResult(v *GetClusterResponseBodyResult) *GetClusterResponseBody {
	s.Result = v
	return s
}

func (s *GetClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyResult struct {
	// The configuration information.
	Config map[string]map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 2021-08-09 00:01:02
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2024-05-21 16:05:26
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The effective advanced configuration version.
	//
	// example:
	//
	// test_yyds_data1
	CurrentAdvanceConfigVersion *string `json:"currentAdvanceConfigVersion,omitempty" xml:"currentAdvanceConfigVersion,omitempty"`
	// The effective online configuration version.
	//
	// example:
	//
	// test_yyds_data1
	CurrentOnlineConfigVersion *string `json:"currentOnlineConfigVersion,omitempty" xml:"currentOnlineConfigVersion,omitempty"`
	// The specifications of Searcher workers.
	DataNode *GetClusterResponseBodyResultDataNode `json:"dataNode,omitempty" xml:"dataNode,omitempty" type:"Struct"`
	// The description of the cluster.
	//
	// example:
	//
	// fzz_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The latest advanced configuration version.
	//
	// example:
	//
	// test_yyds_data1
	LatestAdvanceConfigVersion *string `json:"latestAdvanceConfigVersion,omitempty" xml:"latestAdvanceConfigVersion,omitempty"`
	// The latest online configuration version.
	//
	// example:
	//
	// test_yyds_data1
	LatestOnlineConfigVersion *string `json:"latestOnlineConfigVersion,omitempty" xml:"latestOnlineConfigVersion,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The specifications of Query Result Searcher (QRS) workers.
	QueryNode *GetClusterResponseBodyResultQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
	// The creation status of the cluster. Valid values: NEW and PUBLISH. NEW indicates that the cluster is being created. PUBLISH indicates that the cluster is created.
	//
	// example:
	//
	// NEW
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetClusterResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyResult) GetConfig() map[string]map[string]interface{} {
	return s.Config
}

func (s *GetClusterResponseBodyResult) GetConfigUpdateTime() *string {
	return s.ConfigUpdateTime
}

func (s *GetClusterResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetClusterResponseBodyResult) GetCurrentAdvanceConfigVersion() *string {
	return s.CurrentAdvanceConfigVersion
}

func (s *GetClusterResponseBodyResult) GetCurrentOnlineConfigVersion() *string {
	return s.CurrentOnlineConfigVersion
}

func (s *GetClusterResponseBodyResult) GetDataNode() *GetClusterResponseBodyResultDataNode {
	return s.DataNode
}

func (s *GetClusterResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetClusterResponseBodyResult) GetLatestAdvanceConfigVersion() *string {
	return s.LatestAdvanceConfigVersion
}

func (s *GetClusterResponseBodyResult) GetLatestOnlineConfigVersion() *string {
	return s.LatestOnlineConfigVersion
}

func (s *GetClusterResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetClusterResponseBodyResult) GetQueryNode() *GetClusterResponseBodyResultQueryNode {
	return s.QueryNode
}

func (s *GetClusterResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetClusterResponseBodyResult) SetConfig(v map[string]map[string]interface{}) *GetClusterResponseBodyResult {
	s.Config = v
	return s
}

func (s *GetClusterResponseBodyResult) SetConfigUpdateTime(v string) *GetClusterResponseBodyResult {
	s.ConfigUpdateTime = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetCreateTime(v string) *GetClusterResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetCurrentAdvanceConfigVersion(v string) *GetClusterResponseBodyResult {
	s.CurrentAdvanceConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetCurrentOnlineConfigVersion(v string) *GetClusterResponseBodyResult {
	s.CurrentOnlineConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetDataNode(v *GetClusterResponseBodyResultDataNode) *GetClusterResponseBodyResult {
	s.DataNode = v
	return s
}

func (s *GetClusterResponseBodyResult) SetDescription(v string) *GetClusterResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetLatestAdvanceConfigVersion(v string) *GetClusterResponseBodyResult {
	s.LatestAdvanceConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetLatestOnlineConfigVersion(v string) *GetClusterResponseBodyResult {
	s.LatestOnlineConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetName(v string) *GetClusterResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetQueryNode(v *GetClusterResponseBodyResultQueryNode) *GetClusterResponseBodyResult {
	s.QueryNode = v
	return s
}

func (s *GetClusterResponseBodyResult) SetStatus(v string) *GetClusterResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyResultDataNode struct {
	// The name of the Searcher worker.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of replicas.
	//
	// example:
	//
	// 1
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The number of partitions.
	//
	// example:
	//
	// 2
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetClusterResponseBodyResultDataNode) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyResultDataNode) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyResultDataNode) GetName() *string {
	return s.Name
}

func (s *GetClusterResponseBodyResultDataNode) GetNumber() *int32 {
	return s.Number
}

func (s *GetClusterResponseBodyResultDataNode) GetPartition() *int32 {
	return s.Partition
}

func (s *GetClusterResponseBodyResultDataNode) SetName(v string) *GetClusterResponseBodyResultDataNode {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyResultDataNode) SetNumber(v int32) *GetClusterResponseBodyResultDataNode {
	s.Number = &v
	return s
}

func (s *GetClusterResponseBodyResultDataNode) SetPartition(v int32) *GetClusterResponseBodyResultDataNode {
	s.Partition = &v
	return s
}

func (s *GetClusterResponseBodyResultDataNode) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyResultQueryNode struct {
	// The name of the QRS worker.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The number of replicas.
	//
	// example:
	//
	// 2
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetClusterResponseBodyResultQueryNode) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyResultQueryNode) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyResultQueryNode) GetName() *string {
	return s.Name
}

func (s *GetClusterResponseBodyResultQueryNode) GetNumber() *int32 {
	return s.Number
}

func (s *GetClusterResponseBodyResultQueryNode) GetPartition() *int32 {
	return s.Partition
}

func (s *GetClusterResponseBodyResultQueryNode) SetName(v string) *GetClusterResponseBodyResultQueryNode {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyResultQueryNode) SetNumber(v int32) *GetClusterResponseBodyResultQueryNode {
	s.Number = &v
	return s
}

func (s *GetClusterResponseBodyResultQueryNode) SetPartition(v int32) *GetClusterResponseBodyResultQueryNode {
	s.Partition = &v
	return s
}

func (s *GetClusterResponseBodyResultQueryNode) Validate() error {
	return dara.Validate(s)
}
