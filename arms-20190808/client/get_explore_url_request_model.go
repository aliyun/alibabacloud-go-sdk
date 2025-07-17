// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExploreUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetExploreUrlRequest
	GetClusterId() *string
	SetExpression(v string) *GetExploreUrlRequest
	GetExpression() *string
	SetRegionId(v string) *GetExploreUrlRequest
	GetRegionId() *string
	SetType(v string) *GetExploreUrlRequest
	GetType() *string
}

type GetExploreUrlRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c9cc4c5e220f8461f9d71b6ec6e******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The query statement that corresponds to the data source.
	//
	// example:
	//
	// {app="buy2"}
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the Grafana data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// prometheus
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExploreUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExploreUrlRequest) GoString() string {
	return s.String()
}

func (s *GetExploreUrlRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetExploreUrlRequest) GetExpression() *string {
	return s.Expression
}

func (s *GetExploreUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetExploreUrlRequest) GetType() *string {
	return s.Type
}

func (s *GetExploreUrlRequest) SetClusterId(v string) *GetExploreUrlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetExploreUrlRequest) SetExpression(v string) *GetExploreUrlRequest {
	s.Expression = &v
	return s
}

func (s *GetExploreUrlRequest) SetRegionId(v string) *GetExploreUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetExploreUrlRequest) SetType(v string) *GetExploreUrlRequest {
	s.Type = &v
	return s
}

func (s *GetExploreUrlRequest) Validate() error {
	return dara.Validate(s)
}
