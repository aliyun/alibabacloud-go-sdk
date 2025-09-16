// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
	SetResult(v []*ListClustersResponseBodyResult) *ListClustersResponseBody
	GetResult() []*ListClustersResponseBodyResult
}

type ListClustersResponseBody struct {
	// id of request
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The clusters.
	Result []*ListClustersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClustersResponseBody) GetResult() []*ListClustersResponseBodyResult {
	return s.Result
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetResult(v []*ListClustersResponseBodyResult) *ListClustersResponseBody {
	s.Result = v
	return s
}

func (s *ListClustersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyResult struct {
	// The configuration information.
	Config map[string]map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The time when the configuration was updated.
	//
	// example:
	//
	// " "
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
	// " "
	CurrentAdvanceConfigVersion *string `json:"currentAdvanceConfigVersion,omitempty" xml:"currentAdvanceConfigVersion,omitempty"`
	// The effective dictionary configuration version.
	//
	// example:
	//
	// ha-cn-pl32rf0****_offline_adv_v1
	CurrentOfflineDictConfigVersion *string `json:"currentOfflineDictConfigVersion,omitempty" xml:"currentOfflineDictConfigVersion,omitempty"`
	// The effective online configuration version.
	//
	// example:
	//
	// " "
	CurrentOnlineConfigVersion *string `json:"currentOnlineConfigVersion,omitempty" xml:"currentOnlineConfigVersion,omitempty"`
	// The effective query configuration version.
	//
	// example:
	//
	// ha-cn-pl32rf0****_offline_adv_v1
	CurrentOnlineQueryConfigVersion *string `json:"currentOnlineQueryConfigVersion,omitempty" xml:"currentOnlineQueryConfigVersion,omitempty"`
	// The information about Searcher workers.
	DataNode *ListClustersResponseBodyResultDataNode `json:"dataNode,omitempty" xml:"dataNode,omitempty" type:"Struct"`
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
	// " "
	LatestAdvanceConfigVersion *string `json:"latestAdvanceConfigVersion,omitempty" xml:"latestAdvanceConfigVersion,omitempty"`
	// The latest dictionary configuration version.
	//
	// example:
	//
	// ha-cn-pl32rf0****_offline_adv_v1
	LatestOfflineDictConfigVersion *string `json:"latestOfflineDictConfigVersion,omitempty" xml:"latestOfflineDictConfigVersion,omitempty"`
	// The latest online configuration version.
	//
	// example:
	//
	// " "
	LatestOnlineConfigVersion *string `json:"latestOnlineConfigVersion,omitempty" xml:"latestOnlineConfigVersion,omitempty"`
	// The latest query configuration version.
	//
	// example:
	//
	// ha-cn-pl32rf0****_offline_adv_v1
	LatestOnlineQueryConfigVersion *string `json:"latestOnlineQueryConfigVersion,omitempty" xml:"latestOnlineQueryConfigVersion,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// ha-cn-7pp2pcna701_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The information about Query Result Searcher (QRS) workers.
	QueryNode *ListClustersResponseBodyResultQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
	// The cluster status. Valid values: running: The cluster is running. starting: The cluster is being started. stopping: The cluster is being stopped. stopped: The cluster is stopped.
	//
	// example:
	//
	// "starting"
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListClustersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyResult) GetConfig() map[string]map[string]interface{} {
	return s.Config
}

func (s *ListClustersResponseBodyResult) GetConfigUpdateTime() *string {
	return s.ConfigUpdateTime
}

func (s *ListClustersResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClustersResponseBodyResult) GetCurrentAdvanceConfigVersion() *string {
	return s.CurrentAdvanceConfigVersion
}

func (s *ListClustersResponseBodyResult) GetCurrentOfflineDictConfigVersion() *string {
	return s.CurrentOfflineDictConfigVersion
}

func (s *ListClustersResponseBodyResult) GetCurrentOnlineConfigVersion() *string {
	return s.CurrentOnlineConfigVersion
}

func (s *ListClustersResponseBodyResult) GetCurrentOnlineQueryConfigVersion() *string {
	return s.CurrentOnlineQueryConfigVersion
}

func (s *ListClustersResponseBodyResult) GetDataNode() *ListClustersResponseBodyResultDataNode {
	return s.DataNode
}

func (s *ListClustersResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListClustersResponseBodyResult) GetLatestAdvanceConfigVersion() *string {
	return s.LatestAdvanceConfigVersion
}

func (s *ListClustersResponseBodyResult) GetLatestOfflineDictConfigVersion() *string {
	return s.LatestOfflineDictConfigVersion
}

func (s *ListClustersResponseBodyResult) GetLatestOnlineConfigVersion() *string {
	return s.LatestOnlineConfigVersion
}

func (s *ListClustersResponseBodyResult) GetLatestOnlineQueryConfigVersion() *string {
	return s.LatestOnlineQueryConfigVersion
}

func (s *ListClustersResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListClustersResponseBodyResult) GetQueryNode() *ListClustersResponseBodyResultQueryNode {
	return s.QueryNode
}

func (s *ListClustersResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListClustersResponseBodyResult) SetConfig(v map[string]map[string]interface{}) *ListClustersResponseBodyResult {
	s.Config = v
	return s
}

func (s *ListClustersResponseBodyResult) SetConfigUpdateTime(v string) *ListClustersResponseBodyResult {
	s.ConfigUpdateTime = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCreateTime(v string) *ListClustersResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentAdvanceConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentAdvanceConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentOfflineDictConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentOfflineDictConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentOnlineConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentOnlineConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentOnlineQueryConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentOnlineQueryConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetDataNode(v *ListClustersResponseBodyResultDataNode) *ListClustersResponseBodyResult {
	s.DataNode = v
	return s
}

func (s *ListClustersResponseBodyResult) SetDescription(v string) *ListClustersResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestAdvanceConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestAdvanceConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestOfflineDictConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestOfflineDictConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestOnlineConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestOnlineConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestOnlineQueryConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestOnlineQueryConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetName(v string) *ListClustersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetQueryNode(v *ListClustersResponseBodyResultQueryNode) *ListClustersResponseBodyResult {
	s.QueryNode = v
	return s
}

func (s *ListClustersResponseBodyResult) SetStatus(v string) *ListClustersResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyResultDataNode struct {
	// The name of the Searcher worker.
	//
	// example:
	//
	// ha-cn-8ed2k7brm05_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of Searcher workers.
	//
	// example:
	//
	// 1
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The ID of the partition that is stored on the Searcher worker.
	//
	// example:
	//
	// dt=20220216
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s ListClustersResponseBodyResultDataNode) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyResultDataNode) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyResultDataNode) GetName() *string {
	return s.Name
}

func (s *ListClustersResponseBodyResultDataNode) GetNumber() *int32 {
	return s.Number
}

func (s *ListClustersResponseBodyResultDataNode) GetPartition() *int32 {
	return s.Partition
}

func (s *ListClustersResponseBodyResultDataNode) SetName(v string) *ListClustersResponseBodyResultDataNode {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyResultDataNode) SetNumber(v int32) *ListClustersResponseBodyResultDataNode {
	s.Number = &v
	return s
}

func (s *ListClustersResponseBodyResultDataNode) SetPartition(v int32) *ListClustersResponseBodyResultDataNode {
	s.Partition = &v
	return s
}

func (s *ListClustersResponseBodyResultDataNode) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyResultQueryNode struct {
	// The name of the QRS worker.
	//
	// example:
	//
	// test_0704
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of QRS workers.
	//
	// example:
	//
	// 1
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The ID of the partition that is stored on the QRS worker.
	//
	// example:
	//
	// dt=20211216
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s ListClustersResponseBodyResultQueryNode) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyResultQueryNode) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyResultQueryNode) GetName() *string {
	return s.Name
}

func (s *ListClustersResponseBodyResultQueryNode) GetNumber() *int32 {
	return s.Number
}

func (s *ListClustersResponseBodyResultQueryNode) GetPartition() *int32 {
	return s.Partition
}

func (s *ListClustersResponseBodyResultQueryNode) SetName(v string) *ListClustersResponseBodyResultQueryNode {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyResultQueryNode) SetNumber(v int32) *ListClustersResponseBodyResultQueryNode {
	s.Number = &v
	return s
}

func (s *ListClustersResponseBodyResultQueryNode) SetPartition(v int32) *ListClustersResponseBodyResultQueryNode {
	s.Partition = &v
	return s
}

func (s *ListClustersResponseBodyResultQueryNode) Validate() error {
	return dara.Validate(s)
}
