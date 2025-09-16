// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterRunTimeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetClusterRunTimeInfoResponseBody
	GetRequestId() *string
	SetResult(v []*GetClusterRunTimeInfoResponseBodyResult) *GetClusterRunTimeInfoResponseBody
	GetResult() []*GetClusterRunTimeInfoResponseBodyResult
}

type GetClusterRunTimeInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result []*GetClusterRunTimeInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetClusterRunTimeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterRunTimeInfoResponseBody) GetResult() []*GetClusterRunTimeInfoResponseBodyResult {
	return s.Result
}

func (s *GetClusterRunTimeInfoResponseBody) SetRequestId(v string) *GetClusterRunTimeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBody) SetResult(v []*GetClusterRunTimeInfoResponseBodyResult) *GetClusterRunTimeInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResult struct {
	// The cluster name.
	//
	// example:
	//
	// vpc_hz_domain_1
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The information about the Searcher workers.
	DataNodes []*GetClusterRunTimeInfoResponseBodyResultDataNodes `json:"dataNodes,omitempty" xml:"dataNodes,omitempty" type:"Repeated"`
	// The information about the Query Result Searcher (QRS) workers.
	QueryNode *GetClusterRunTimeInfoResponseBodyResultQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
}

func (s GetClusterRunTimeInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResult) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetClusterRunTimeInfoResponseBodyResult) GetDataNodes() []*GetClusterRunTimeInfoResponseBodyResultDataNodes {
	return s.DataNodes
}

func (s *GetClusterRunTimeInfoResponseBodyResult) GetQueryNode() *GetClusterRunTimeInfoResponseBodyResultQueryNode {
	return s.QueryNode
}

func (s *GetClusterRunTimeInfoResponseBodyResult) SetClusterName(v string) *GetClusterRunTimeInfoResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResult) SetDataNodes(v []*GetClusterRunTimeInfoResponseBodyResultDataNodes) *GetClusterRunTimeInfoResponseBodyResult {
	s.DataNodes = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResult) SetQueryNode(v *GetClusterRunTimeInfoResponseBodyResultQueryNode) *GetClusterRunTimeInfoResponseBodyResult {
	s.QueryNode = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResultDataNodes struct {
	// The configuration status.
	ConfigStatusList []*GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList `json:"configStatusList,omitempty" xml:"configStatusList,omitempty" type:"Repeated"`
	// The data of the Searcher worker.
	DataStatusList []*GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList `json:"dataStatusList,omitempty" xml:"dataStatusList,omitempty" type:"Repeated"`
	// The service status of the QRS worker.
	ServiceStatus *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty" type:"Struct"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodes) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodes) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) GetConfigStatusList() []*GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	return s.ConfigStatusList
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) GetDataStatusList() []*GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	return s.DataStatusList
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) GetServiceStatus() *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	return s.ServiceStatus
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) SetConfigStatusList(v []*GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) *GetClusterRunTimeInfoResponseBodyResultDataNodes {
	s.ConfigStatusList = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) SetDataStatusList(v []*GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) *GetClusterRunTimeInfoResponseBodyResultDataNodes {
	s.DataStatusList = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) SetServiceStatus(v *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) *GetClusterRunTimeInfoResponseBodyResultDataNodes {
	s.ServiceStatus = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList struct {
	// The time when the configuration was last updated.
	//
	// example:
	//
	// ""
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	// The configuration progress. Unit: percentage.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of processed Searcher workers in the cluster.
	//
	// example:
	//
	// 1
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test_0704
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of Searcher workers in the cluster.
	//
	// example:
	//
	// 0
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) GetConfigUpdateTime() *string {
	return s.ConfigUpdateTime
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) GetDonePercent() *int32 {
	return s.DonePercent
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) GetDoneSize() *int32 {
	return s.DoneSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) GetName() *string {
	return s.Name
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetConfigUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.ConfigUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.TotalSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList struct {
	// The information about the advanced configuration.
	AdvanceConfigInfo *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo `json:"advanceConfigInfo,omitempty" xml:"advanceConfigInfo,omitempty" type:"Struct"`
	// The name of the worker that failed due to a deployment failure.
	DeployFailedWorker []*string `json:"deployFailedWorker,omitempty" xml:"deployFailedWorker,omitempty" type:"Repeated"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 2
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The configuration progress. Unit: percentage.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of processed QRS workers in the cluster.
	//
	// example:
	//
	// 100
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The error message.
	//
	// example:
	//
	// 0A3B1C48006A6C0905F6375F4821EB50
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The time when full data in the index was last updated.
	//
	// example:
	//
	// " "
	FullUpdateTime *string `json:"fullUpdateTime,omitempty" xml:"fullUpdateTime,omitempty"`
	// The time when the full index version was generated.
	//
	// example:
	//
	// 123423
	FullVersion *int64 `json:"fullVersion,omitempty" xml:"fullVersion,omitempty"`
	// The time when incremental data in the index was last updated.
	//
	// example:
	//
	// ""
	IncUpdateTime *string `json:"incUpdateTime,omitempty" xml:"incUpdateTime,omitempty"`
	// The time when the incremental index version was generated.
	//
	// example:
	//
	// 123423
	IncVersion *int64 `json:"incVersion,omitempty" xml:"incVersion,omitempty"`
	// The information about the index configuration.
	IndexConfigInfo *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo `json:"indexConfigInfo,omitempty" xml:"indexConfigInfo,omitempty" type:"Struct"`
	// The index size.
	//
	// example:
	//
	// 100
	IndexSize *int64 `json:"indexSize,omitempty" xml:"indexSize,omitempty"`
	// The name of the worker that failed due to insufficient disks.
	LackDiskWorker []*string `json:"lackDiskWorker,omitempty" xml:"lackDiskWorker,omitempty" type:"Repeated"`
	// The name of the worker that failed due to insufficient memory.
	LackMemWorker []*string `json:"lackMemWorker,omitempty" xml:"lackMemWorker,omitempty" type:"Repeated"`
	// The name of the QRS worker.
	//
	// example:
	//
	// ha-cn-c4d2rq7nt04_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of QRS workers in the cluster.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetAdvanceConfigInfo() *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo {
	return s.AdvanceConfigInfo
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetDeployFailedWorker() []*string {
	return s.DeployFailedWorker
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetDocSize() *int32 {
	return s.DocSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetDonePercent() *int32 {
	return s.DonePercent
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetDoneSize() *int32 {
	return s.DoneSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetFullUpdateTime() *string {
	return s.FullUpdateTime
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetFullVersion() *int64 {
	return s.FullVersion
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetIncUpdateTime() *string {
	return s.IncUpdateTime
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetIncVersion() *int64 {
	return s.IncVersion
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetIndexConfigInfo() *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo {
	return s.IndexConfigInfo
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetIndexSize() *int64 {
	return s.IndexSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetLackDiskWorker() []*string {
	return s.LackDiskWorker
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetLackMemWorker() []*string {
	return s.LackMemWorker
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetName() *string {
	return s.Name
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetAdvanceConfigInfo(v *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.AdvanceConfigInfo = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDeployFailedWorker(v []*string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DeployFailedWorker = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDocSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DocSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetErrorMsg(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.ErrorMsg = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetFullUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.FullUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetFullVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.FullVersion = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIncUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IncUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIncVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IncVersion = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIndexConfigInfo(v *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IndexConfigInfo = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIndexSize(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IndexSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetLackDiskWorker(v []*string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.LackDiskWorker = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetLackMemWorker(v []*string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.LackMemWorker = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.TotalSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo struct {
	// The name of the index configuration.
	//
	// example:
	//
	// index_meta_name
	ConfigMetaName *string `json:"configMetaName,omitempty" xml:"configMetaName,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) GetConfigMetaName() *string {
	return s.ConfigMetaName
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) GetVersion() *int64 {
	return s.Version
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) SetConfigMetaName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo {
	s.ConfigMetaName = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) SetVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo {
	s.Version = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo struct {
	// The name of the index configuration.
	//
	// example:
	//
	// index_meta_name
	ConfigMetaName *string `json:"configMetaName,omitempty" xml:"configMetaName,omitempty"`
	// The version of the index template.
	//
	// example:
	//
	// 1.0.0
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) GetConfigMetaName() *string {
	return s.ConfigMetaName
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) GetVersion() *int64 {
	return s.Version
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) SetConfigMetaName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo {
	s.ConfigMetaName = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) SetVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo {
	s.Version = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus struct {
	// The process progress of QRS workers in the cluster. Unit: percentage.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of processed QRS workers in the cluster.
	//
	// example:
	//
	// 100
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The name of the QRS worker.
	//
	// example:
	//
	// ha-cn-0ju2s170b03_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of QRS workers in the cluster.
	//
	// example:
	//
	// 100
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) GetDonePercent() *int32 {
	return s.DonePercent
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) GetDoneSize() *int32 {
	return s.DoneSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) GetName() *string {
	return s.Name
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.TotalSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResultQueryNode struct {
	// The configuration status.
	ConfigStatusList []*GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList `json:"configStatusList,omitempty" xml:"configStatusList,omitempty" type:"Repeated"`
	// The service status of the QRS worker.
	ServiceStatus *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty" type:"Struct"`
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNode) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNode) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNode) GetConfigStatusList() []*GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	return s.ConfigStatusList
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNode) GetServiceStatus() *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	return s.ServiceStatus
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNode) SetConfigStatusList(v []*GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) *GetClusterRunTimeInfoResponseBodyResultQueryNode {
	s.ConfigStatusList = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNode) SetServiceStatus(v *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) *GetClusterRunTimeInfoResponseBodyResultQueryNode {
	s.ServiceStatus = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNode) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList struct {
	// The time when the configuration was last updated.
	//
	// example:
	//
	// " "
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	// The process progress of QRS workers in the cluster. Unit: percentage.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of processed QRS workers in the cluster.
	//
	// example:
	//
	// 100
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// ha-cn-zvp2qr1sk01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of QRS workers in the cluster.
	//
	// example:
	//
	// 6
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) GetConfigUpdateTime() *string {
	return s.ConfigUpdateTime
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) GetDonePercent() *int32 {
	return s.DonePercent
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) GetDoneSize() *int32 {
	return s.DoneSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) GetName() *string {
	return s.Name
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetConfigUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.ConfigUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.TotalSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) Validate() error {
	return dara.Validate(s)
}

type GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus struct {
	// The process progress of QRS workers in the cluster. Unit: percentage.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of processed QRS workers in the cluster.
	//
	// example:
	//
	// 100
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// ha-cn-c4d2rq7nt04_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of QRS workers in the cluster.
	//
	// example:
	//
	// 100
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) GetDonePercent() *int32 {
	return s.DonePercent
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) GetDoneSize() *int32 {
	return s.DoneSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) GetName() *string {
	return s.Name
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.TotalSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) Validate() error {
	return dara.Validate(s)
}
