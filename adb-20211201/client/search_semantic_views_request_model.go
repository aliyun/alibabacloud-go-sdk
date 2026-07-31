// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSemanticViewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *SearchSemanticViewsRequest
	GetDBClusterId() *string
	SetQueryText(v string) *SearchSemanticViewsRequest
	GetQueryText() *string
	SetTopK(v int32) *SearchSemanticViewsRequest
	GetTopK() *int32
}

type SearchSemanticViewsRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The search query.
	//
	// example:
	//
	// 销售额
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// The number of the most relevant semantic views to return.
	//
	// example:
	//
	// 3
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s SearchSemanticViewsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchSemanticViewsRequest) GoString() string {
	return s.String()
}

func (s *SearchSemanticViewsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SearchSemanticViewsRequest) GetQueryText() *string {
	return s.QueryText
}

func (s *SearchSemanticViewsRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *SearchSemanticViewsRequest) SetDBClusterId(v string) *SearchSemanticViewsRequest {
	s.DBClusterId = &v
	return s
}

func (s *SearchSemanticViewsRequest) SetQueryText(v string) *SearchSemanticViewsRequest {
	s.QueryText = &v
	return s
}

func (s *SearchSemanticViewsRequest) SetTopK(v int32) *SearchSemanticViewsRequest {
	s.TopK = &v
	return s
}

func (s *SearchSemanticViewsRequest) Validate() error {
	return dara.Validate(s)
}
