// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDataTasksResponseBody
	GetRequestId() *string
	SetResult(v []*ListDataTasksResponseBodyResult) *ListDataTasksResponseBody
	GetResult() []*ListDataTasksResponseBodyResult
}

type ListDataTasksResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result []*ListDataTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDataTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataTasksResponseBody) GetResult() []*ListDataTasksResponseBodyResult {
	return s.Result
}

func (s *ListDataTasksResponseBody) SetRequestId(v string) *ListDataTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataTasksResponseBody) SetResult(v []*ListDataTasksResponseBodyResult) *ListDataTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListDataTasksResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataTasksResponseBodyResult struct {
	// The time when the site monitoring task was created.
	//
	// example:
	//
	// 2020-07-30 06:32:18
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The information of the target cluster.
	SinkCluster *ListDataTasksResponseBodyResultSinkCluster `json:"sinkCluster,omitempty" xml:"sinkCluster,omitempty" type:"Struct"`
	// The information about the source cluster.
	SourceCluster *ListDataTasksResponseBodyResultSourceCluster `json:"sourceCluster,omitempty" xml:"sourceCluster,omitempty" type:"Struct"`
	// The status of the task.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// et_cn_mfv1233r47272****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListDataTasksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDataTasksResponseBodyResult) GetSinkCluster() *ListDataTasksResponseBodyResultSinkCluster {
	return s.SinkCluster
}

func (s *ListDataTasksResponseBodyResult) GetSourceCluster() *ListDataTasksResponseBodyResultSourceCluster {
	return s.SourceCluster
}

func (s *ListDataTasksResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListDataTasksResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *ListDataTasksResponseBodyResult) SetCreateTime(v string) *ListDataTasksResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListDataTasksResponseBodyResult) SetSinkCluster(v *ListDataTasksResponseBodyResultSinkCluster) *ListDataTasksResponseBodyResult {
	s.SinkCluster = v
	return s
}

func (s *ListDataTasksResponseBodyResult) SetSourceCluster(v *ListDataTasksResponseBodyResultSourceCluster) *ListDataTasksResponseBodyResult {
	s.SourceCluster = v
	return s
}

func (s *ListDataTasksResponseBodyResult) SetStatus(v string) *ListDataTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataTasksResponseBodyResult) SetTaskId(v string) *ListDataTasksResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *ListDataTasksResponseBodyResult) Validate() error {
	if s.SinkCluster != nil {
		if err := s.SinkCluster.Validate(); err != nil {
			return err
		}
	}
	if s.SourceCluster != nil {
		if err := s.SourceCluster.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataTasksResponseBodyResultSinkCluster struct {
	// The type of the target cluster. Default value: elasticsearch.
	//
	// example:
	//
	// 1
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// The public network access address of the target cluster.
	//
	// example:
	//
	// http://192.168.xx.xx:4101
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The target index.
	//
	// example:
	//
	// product_info
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// The type of the destination index.
	//
	// example:
	//
	// _doc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The ID of the VPC to which the cluster belongs.
	//
	// example:
	//
	// vpc-2ze55voww95g82gak****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The instance ID or Server Load Balancer (SLB) ID of the current cluster.
	//
	// example:
	//
	// es-cn-09k1rnu3g0002****-worker
	VpcInstanceId *string `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
	// The access port number of the cluster.
	//
	// example:
	//
	// 9200
	VpcInstancePort *string `json:"vpcInstancePort,omitempty" xml:"vpcInstancePort,omitempty"`
}

func (s ListDataTasksResponseBodyResultSinkCluster) String() string {
	return dara.Prettify(s)
}

func (s ListDataTasksResponseBodyResultSinkCluster) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponseBodyResultSinkCluster) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListDataTasksResponseBodyResultSinkCluster) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListDataTasksResponseBodyResultSinkCluster) GetIndex() *string {
	return s.Index
}

func (s *ListDataTasksResponseBodyResultSinkCluster) GetType() *string {
	return s.Type
}

func (s *ListDataTasksResponseBodyResultSinkCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *ListDataTasksResponseBodyResultSinkCluster) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *ListDataTasksResponseBodyResultSinkCluster) GetVpcInstancePort() *string {
	return s.VpcInstancePort
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetDataSourceType(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.DataSourceType = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetEndpoint(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.Endpoint = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetIndex(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.Index = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetType(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.Type = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetVpcId(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.VpcId = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetVpcInstanceId(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.VpcInstanceId = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetVpcInstancePort(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.VpcInstancePort = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) Validate() error {
	return dara.Validate(s)
}

type ListDataTasksResponseBodyResultSourceCluster struct {
	// The type of the source cluster. Default value: elasticsearch.
	//
	// example:
	//
	// 1
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// The index whose data you want to migrate.
	//
	// example:
	//
	// product_info
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// The Mapping configuration of the cluster.
	//
	// example:
	//
	// {\\"_doc\\":{\\"properties\\":{\\"user\\":{\\"properties\\":{\\"last\\":{\\"type\\":\\"text\\",...}}}}}}
	Mapping *string `json:"mapping,omitempty" xml:"mapping,omitempty"`
	// The routing field to index the table. It is set to the primary key by default.
	//
	// example:
	//
	// _id
	Routing *string `json:"routing,omitempty" xml:"routing,omitempty"`
	// The Settings of the cluster.
	//
	// example:
	//
	// {\\n  \\"index\\": {\\n    \\"replication\\": {\\n}.....}}
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
	// The type of the destination index.
	//
	// example:
	//
	// _doc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataTasksResponseBodyResultSourceCluster) String() string {
	return dara.Prettify(s)
}

func (s ListDataTasksResponseBodyResultSourceCluster) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponseBodyResultSourceCluster) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListDataTasksResponseBodyResultSourceCluster) GetIndex() *string {
	return s.Index
}

func (s *ListDataTasksResponseBodyResultSourceCluster) GetMapping() *string {
	return s.Mapping
}

func (s *ListDataTasksResponseBodyResultSourceCluster) GetRouting() *string {
	return s.Routing
}

func (s *ListDataTasksResponseBodyResultSourceCluster) GetSettings() *string {
	return s.Settings
}

func (s *ListDataTasksResponseBodyResultSourceCluster) GetType() *string {
	return s.Type
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetDataSourceType(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.DataSourceType = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetIndex(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Index = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetMapping(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Mapping = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetRouting(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Routing = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetSettings(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Settings = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetType(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Type = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) Validate() error {
	return dara.Validate(s)
}
