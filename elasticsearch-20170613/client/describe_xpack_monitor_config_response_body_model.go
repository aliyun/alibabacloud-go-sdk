// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeXpackMonitorConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeXpackMonitorConfigResponseBody
	GetRequestId() *string
	SetResult(v *DescribeXpackMonitorConfigResponseBodyResult) *DescribeXpackMonitorConfigResponseBody
	GetResult() *DescribeXpackMonitorConfigResponseBodyResult
}

type DescribeXpackMonitorConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *DescribeXpackMonitorConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeXpackMonitorConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeXpackMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeXpackMonitorConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeXpackMonitorConfigResponseBody) GetResult() *DescribeXpackMonitorConfigResponseBodyResult {
	return s.Result
}

func (s *DescribeXpackMonitorConfigResponseBody) SetRequestId(v string) *DescribeXpackMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBody) SetResult(v *DescribeXpackMonitorConfigResponseBodyResult) *DescribeXpackMonitorConfigResponseBody {
	s.Result = v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeXpackMonitorConfigResponseBodyResult struct {
	// Indicates whether the X-Pack Monitoring feature is enabled. Valid values:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	Enable    *bool     `json:"enable,omitempty" xml:"enable,omitempty"`
	Endpoints []*string `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	// The ID of the associated Elasticsearch cluster.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****
	EsInstanceId *string   `json:"esInstanceId,omitempty" xml:"esInstanceId,omitempty"`
	PipelineIds  []*string `json:"pipelineIds,omitempty" xml:"pipelineIds,omitempty" type:"Repeated"`
	// The username that is used to access the associated Elasticsearch cluster.
	//
	// example:
	//
	// elastic
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DescribeXpackMonitorConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeXpackMonitorConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) GetEndpoints() []*string {
	return s.Endpoints
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) GetEsInstanceId() *string {
	return s.EsInstanceId
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) GetPipelineIds() []*string {
	return s.PipelineIds
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) GetUserName() *string {
	return s.UserName
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetEnable(v bool) *DescribeXpackMonitorConfigResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetEndpoints(v []*string) *DescribeXpackMonitorConfigResponseBodyResult {
	s.Endpoints = v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetEsInstanceId(v string) *DescribeXpackMonitorConfigResponseBodyResult {
	s.EsInstanceId = &v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetPipelineIds(v []*string) *DescribeXpackMonitorConfigResponseBodyResult {
	s.PipelineIds = v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetUserName(v string) *DescribeXpackMonitorConfigResponseBodyResult {
	s.UserName = &v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
