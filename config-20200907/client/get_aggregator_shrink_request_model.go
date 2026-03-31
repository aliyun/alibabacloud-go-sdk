// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregatorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregatorShrinkRequest
	GetAggregatorId() *string
	SetTagShrink(v string) *GetAggregatorShrinkRequest
	GetTagShrink() *string
}

type GetAggregatorShrinkRequest struct {
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-88ea626622af0055****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// Deprecated
	//
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetAggregatorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatorShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAggregatorShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregatorShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *GetAggregatorShrinkRequest) SetAggregatorId(v string) *GetAggregatorShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregatorShrinkRequest) SetTagShrink(v string) *GetAggregatorShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *GetAggregatorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
