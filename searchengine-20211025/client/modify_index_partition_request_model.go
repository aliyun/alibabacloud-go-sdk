// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexPartitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceName(v string) *ModifyIndexPartitionRequest
	GetDataSourceName() *string
	SetDomainName(v string) *ModifyIndexPartitionRequest
	GetDomainName() *string
	SetGeneration(v int64) *ModifyIndexPartitionRequest
	GetGeneration() *int64
	SetIndexInfos(v []*ModifyIndexPartitionRequestIndexInfos) *ModifyIndexPartitionRequest
	GetIndexInfos() []*ModifyIndexPartitionRequestIndexInfos
}

type ModifyIndexPartitionRequest struct {
	// The name of the data source.
	//
	// example:
	//
	// test1
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The data center.
	//
	// example:
	//
	// pre_domain_1
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The primary key.
	//
	// example:
	//
	// 1633293829
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The index information.
	IndexInfos []*ModifyIndexPartitionRequestIndexInfos `json:"indexInfos,omitempty" xml:"indexInfos,omitempty" type:"Repeated"`
}

func (s ModifyIndexPartitionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexPartitionRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ModifyIndexPartitionRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyIndexPartitionRequest) GetGeneration() *int64 {
	return s.Generation
}

func (s *ModifyIndexPartitionRequest) GetIndexInfos() []*ModifyIndexPartitionRequestIndexInfos {
	return s.IndexInfos
}

func (s *ModifyIndexPartitionRequest) SetDataSourceName(v string) *ModifyIndexPartitionRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyIndexPartitionRequest) SetDomainName(v string) *ModifyIndexPartitionRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyIndexPartitionRequest) SetGeneration(v int64) *ModifyIndexPartitionRequest {
	s.Generation = &v
	return s
}

func (s *ModifyIndexPartitionRequest) SetIndexInfos(v []*ModifyIndexPartitionRequestIndexInfos) *ModifyIndexPartitionRequest {
	s.IndexInfos = v
	return s
}

func (s *ModifyIndexPartitionRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyIndexPartitionRequestIndexInfos struct {
	// The index name.
	//
	// example:
	//
	// atest2
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The concurrency. Default value: 1.
	//
	// example:
	//
	// 1
	ParallelNum *int32 `json:"parallelNum,omitempty" xml:"parallelNum,omitempty"`
	// The number of shards.
	//
	// example:
	//
	// 3
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
}

func (s ModifyIndexPartitionRequestIndexInfos) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexPartitionRequestIndexInfos) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionRequestIndexInfos) GetIndexName() *string {
	return s.IndexName
}

func (s *ModifyIndexPartitionRequestIndexInfos) GetParallelNum() *int32 {
	return s.ParallelNum
}

func (s *ModifyIndexPartitionRequestIndexInfos) GetPartitionCount() *int32 {
	return s.PartitionCount
}

func (s *ModifyIndexPartitionRequestIndexInfos) SetIndexName(v string) *ModifyIndexPartitionRequestIndexInfos {
	s.IndexName = &v
	return s
}

func (s *ModifyIndexPartitionRequestIndexInfos) SetParallelNum(v int32) *ModifyIndexPartitionRequestIndexInfos {
	s.ParallelNum = &v
	return s
}

func (s *ModifyIndexPartitionRequestIndexInfos) SetPartitionCount(v int32) *ModifyIndexPartitionRequestIndexInfos {
	s.PartitionCount = &v
	return s
}

func (s *ModifyIndexPartitionRequestIndexInfos) Validate() error {
	return dara.Validate(s)
}
