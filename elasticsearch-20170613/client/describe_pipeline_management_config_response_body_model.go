// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePipelineManagementConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePipelineManagementConfigResponseBody
	GetRequestId() *string
	SetResult(v *DescribePipelineManagementConfigResponseBodyResult) *DescribePipelineManagementConfigResponseBody
	GetResult() *DescribePipelineManagementConfigResponseBodyResult
}

type DescribePipelineManagementConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *DescribePipelineManagementConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribePipelineManagementConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineManagementConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePipelineManagementConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePipelineManagementConfigResponseBody) GetResult() *DescribePipelineManagementConfigResponseBodyResult {
	return s.Result
}

func (s *DescribePipelineManagementConfigResponseBody) SetRequestId(v string) *DescribePipelineManagementConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePipelineManagementConfigResponseBody) SetResult(v *DescribePipelineManagementConfigResponseBodyResult) *DescribePipelineManagementConfigResponseBody {
	s.Result = v
	return s
}

func (s *DescribePipelineManagementConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePipelineManagementConfigResponseBodyResult struct {
	// The access addresses of the Elasticsearch cluster. Specify each address in the `http://Endpoint of the Elasticsearch cluster:Port number` format.
	//
	// example:
	//
	// ["http://es-cn-n6w1o1x0w001c****.elasticsearch.aliyuncs.com:9200"]
	Endpoints *string `json:"endpoints,omitempty" xml:"endpoints,omitempty"`
	// The ID of the Elasticsearch cluster.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****
	EsInstanceId *string   `json:"esInstanceId,omitempty" xml:"esInstanceId,omitempty"`
	PipelineIds  []*string `json:"pipelineIds,omitempty" xml:"pipelineIds,omitempty" type:"Repeated"`
	// The pipeline management method. Valid values: Kibana and MULTIPLE_PIPELINE.
	//
	// example:
	//
	// MULTIPLE_PIPELINE
	PipelineManagementType *string `json:"pipelineManagementType,omitempty" xml:"pipelineManagementType,omitempty"`
	// The username that is used to access the Elasticsearch cluster.
	//
	// example:
	//
	// elastic
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DescribePipelineManagementConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineManagementConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribePipelineManagementConfigResponseBodyResult) GetEndpoints() *string {
	return s.Endpoints
}

func (s *DescribePipelineManagementConfigResponseBodyResult) GetEsInstanceId() *string {
	return s.EsInstanceId
}

func (s *DescribePipelineManagementConfigResponseBodyResult) GetPipelineIds() []*string {
	return s.PipelineIds
}

func (s *DescribePipelineManagementConfigResponseBodyResult) GetPipelineManagementType() *string {
	return s.PipelineManagementType
}

func (s *DescribePipelineManagementConfigResponseBodyResult) GetUserName() *string {
	return s.UserName
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetEndpoints(v string) *DescribePipelineManagementConfigResponseBodyResult {
	s.Endpoints = &v
	return s
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetEsInstanceId(v string) *DescribePipelineManagementConfigResponseBodyResult {
	s.EsInstanceId = &v
	return s
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetPipelineIds(v []*string) *DescribePipelineManagementConfigResponseBodyResult {
	s.PipelineIds = v
	return s
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetPipelineManagementType(v string) *DescribePipelineManagementConfigResponseBodyResult {
	s.PipelineManagementType = &v
	return s
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetUserName(v string) *DescribePipelineManagementConfigResponseBodyResult {
	s.UserName = &v
	return s
}

func (s *DescribePipelineManagementConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
