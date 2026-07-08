// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMatchedResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *GetMatchedResourcesResponseBody
	GetData() interface{}
	SetMaxResults(v int32) *GetMatchedResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *GetMatchedResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetMatchedResourcesResponseBody
	GetRequestId() *string
}

type GetMatchedResourcesResponseBody struct {
	// 请求接口返回的数据。
	//
	// example:
	//
	// []
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// 分页参数：结果集的最大数量，默认值为 20。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 下一个查询开始 Token，NextToken 为空说明没有下一个。
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次请求的 ID。
	//
	// example:
	//
	// 26F62CED-1E0E-51AA-B8EB-BCD61C5B0C50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMatchedResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMatchedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetMatchedResourcesResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetMatchedResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetMatchedResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetMatchedResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMatchedResourcesResponseBody) SetData(v interface{}) *GetMatchedResourcesResponseBody {
	s.Data = v
	return s
}

func (s *GetMatchedResourcesResponseBody) SetMaxResults(v int32) *GetMatchedResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetMatchedResourcesResponseBody) SetNextToken(v string) *GetMatchedResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetMatchedResourcesResponseBody) SetRequestId(v string) *GetMatchedResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMatchedResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
