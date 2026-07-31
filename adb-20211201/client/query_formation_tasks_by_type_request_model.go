// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFormationTasksByTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *QueryFormationTasksByTypeRequest
	GetDBClusterId() *string
	SetRegionId(v string) *QueryFormationTasksByTypeRequest
	GetRegionId() *string
	SetTaskType(v string) *QueryFormationTasksByTypeRequest
	GetTaskType() *string
}

type QueryFormationTasksByTypeRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the details of all AnalyticDB for MySQL clusters in a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the supported regions and zones, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["CRAWLER"]
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryFormationTasksByTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationTasksByTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryFormationTasksByTypeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *QueryFormationTasksByTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryFormationTasksByTypeRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryFormationTasksByTypeRequest) SetDBClusterId(v string) *QueryFormationTasksByTypeRequest {
	s.DBClusterId = &v
	return s
}

func (s *QueryFormationTasksByTypeRequest) SetRegionId(v string) *QueryFormationTasksByTypeRequest {
	s.RegionId = &v
	return s
}

func (s *QueryFormationTasksByTypeRequest) SetTaskType(v string) *QueryFormationTasksByTypeRequest {
	s.TaskType = &v
	return s
}

func (s *QueryFormationTasksByTypeRequest) Validate() error {
	return dara.Validate(s)
}
