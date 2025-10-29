// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableEsInstanceIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAvailableEsInstanceIdsResponseBody
	GetRequestId() *string
	SetResult(v []*ListAvailableEsInstanceIdsResponseBodyResult) *ListAvailableEsInstanceIdsResponseBody
	GetResult() []*ListAvailableEsInstanceIdsResponseBodyResult
}

type ListAvailableEsInstanceIdsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*ListAvailableEsInstanceIdsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAvailableEsInstanceIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableEsInstanceIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableEsInstanceIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableEsInstanceIdsResponseBody) GetResult() []*ListAvailableEsInstanceIdsResponseBodyResult {
	return s.Result
}

func (s *ListAvailableEsInstanceIdsResponseBody) SetRequestId(v string) *ListAvailableEsInstanceIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBody) SetResult(v []*ListAvailableEsInstanceIdsResponseBodyResult) *ListAvailableEsInstanceIdsResponseBody {
	s.Result = v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBody) Validate() error {
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

type ListAvailableEsInstanceIdsResponseBodyResult struct {
	// The name of the Elasticsearch cluster.
	//
	// example:
	//
	// instanceName
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The address that is used to access the Elasticsearch cluster over the Internet.
	//
	// example:
	//
	// http://es-cn-n6w1o1x0w001c****.elasticsearch.aliyuncs.com:9200
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The ID of the Elasticsearch cluster.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****
	EsInstanceId *string `json:"esInstanceId,omitempty" xml:"esInstanceId,omitempty"`
	// The address that is used to access the Kibana console of the Elasticsearch cluster over the Internet.
	//
	// example:
	//
	// https://es-cn-n6w1o1x0w001c****.kibana.elasticsearch.aliyuncs.com:5601
	KibanaEndpoint *string `json:"kibanaEndpoint,omitempty" xml:"kibanaEndpoint,omitempty"`
}

func (s ListAvailableEsInstanceIdsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableEsInstanceIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) GetEsInstanceId() *string {
	return s.EsInstanceId
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) GetKibanaEndpoint() *string {
	return s.KibanaEndpoint
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) SetDescription(v string) *ListAvailableEsInstanceIdsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) SetEndpoint(v string) *ListAvailableEsInstanceIdsResponseBodyResult {
	s.Endpoint = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) SetEsInstanceId(v string) *ListAvailableEsInstanceIdsResponseBodyResult {
	s.EsInstanceId = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) SetKibanaEndpoint(v string) *ListAvailableEsInstanceIdsResponseBodyResult {
	s.KibanaEndpoint = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
