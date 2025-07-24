// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiTemplates(v []*ApiTemplate) *ListApiTemplatesResponseBody
	GetApiTemplates() []*ApiTemplate
	SetMaxResults(v int32) *ListApiTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApiTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApiTemplatesResponseBody
	GetTotalCount() *int32
}

type ListApiTemplatesResponseBody struct {
	// Deprecated
	//
	// The array of API templates.
	ApiTemplates []*ApiTemplate `json:"ApiTemplates,omitempty" xml:"ApiTemplates,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 返回读取到的数据位置，空代表数据已经读取完毕。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 本次请求条件下的数据总量。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApiTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiTemplatesResponseBody) GetApiTemplates() []*ApiTemplate {
	return s.ApiTemplates
}

func (s *ListApiTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApiTemplatesResponseBody) SetApiTemplates(v []*ApiTemplate) *ListApiTemplatesResponseBody {
	s.ApiTemplates = v
	return s
}

func (s *ListApiTemplatesResponseBody) SetMaxResults(v int32) *ListApiTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApiTemplatesResponseBody) SetNextToken(v string) *ListApiTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApiTemplatesResponseBody) SetRequestId(v string) *ListApiTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiTemplatesResponseBody) SetTotalCount(v int32) *ListApiTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApiTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
