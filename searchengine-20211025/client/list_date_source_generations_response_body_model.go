// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDateSourceGenerationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDateSourceGenerationsResponseBody
	GetRequestId() *string
	SetResult(v []*ListDateSourceGenerationsResponseBodyResult) *ListDateSourceGenerationsResponseBody
	GetResult() []*ListDateSourceGenerationsResponseBodyResult
}

type ListDateSourceGenerationsResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListDateSourceGenerationsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDateSourceGenerationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDateSourceGenerationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDateSourceGenerationsResponseBody) GetResult() []*ListDateSourceGenerationsResponseBodyResult {
	return s.Result
}

func (s *ListDateSourceGenerationsResponseBody) SetRequestId(v string) *ListDateSourceGenerationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBody) SetResult(v []*ListDateSourceGenerationsResponseBodyResult) *ListDateSourceGenerationsResponseBody {
	s.Result = v
	return s
}

func (s *ListDateSourceGenerationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDateSourceGenerationsResponseBodyResult struct {
	// The ID of the offline deployment.
	//
	// example:
	//
	// 122
	BuildDeployId *int32 `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The timestamp that was generated when the index building was started.
	//
	// example:
	//
	// 1626143673
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The path of the dumped index in the Apsara File Storage for HDFS file system.
	//
	// example:
	//
	// ""
	DataDumpRoot *string `json:"dataDumpRoot,omitempty" xml:"dataDumpRoot,omitempty"`
	// The ID of the full index version.
	//
	// example:
	//
	// 1626143930
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The shards of the index version. The value is a key-value pair in which the key indicates the index name and the value indicates the number of shards. The number of value shards.
	Partition map[string]*int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The status of the index version.
	//
	// example:
	//
	// STOPPED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The start timestamp from which incremental data is retrieved.
	//
	// example:
	//
	// 1626143673
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ListDateSourceGenerationsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDateSourceGenerationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsResponseBodyResult) GetBuildDeployId() *int32 {
	return s.BuildDeployId
}

func (s *ListDateSourceGenerationsResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDateSourceGenerationsResponseBodyResult) GetDataDumpRoot() *string {
	return s.DataDumpRoot
}

func (s *ListDateSourceGenerationsResponseBodyResult) GetGeneration() *int64 {
	return s.Generation
}

func (s *ListDateSourceGenerationsResponseBodyResult) GetPartition() map[string]*int32 {
	return s.Partition
}

func (s *ListDateSourceGenerationsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListDateSourceGenerationsResponseBodyResult) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetBuildDeployId(v int32) *ListDateSourceGenerationsResponseBodyResult {
	s.BuildDeployId = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetCreateTime(v int64) *ListDateSourceGenerationsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetDataDumpRoot(v string) *ListDateSourceGenerationsResponseBodyResult {
	s.DataDumpRoot = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetGeneration(v int64) *ListDateSourceGenerationsResponseBodyResult {
	s.Generation = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetPartition(v map[string]*int32) *ListDateSourceGenerationsResponseBodyResult {
	s.Partition = v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetStatus(v string) *ListDateSourceGenerationsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetTimestamp(v int64) *ListDateSourceGenerationsResponseBodyResult {
	s.Timestamp = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
