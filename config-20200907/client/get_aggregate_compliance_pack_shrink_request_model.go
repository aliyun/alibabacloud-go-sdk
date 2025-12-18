// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateCompliancePackShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateCompliancePackShrinkRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *GetAggregateCompliancePackShrinkRequest
	GetCompliancePackId() *string
	SetTagShrink(v string) *GetAggregateCompliancePackShrinkRequest
	GetTagShrink() *string
}

type GetAggregateCompliancePackShrinkRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// Deprecated
	//
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetAggregateCompliancePackShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateCompliancePackShrinkRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateCompliancePackShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *GetAggregateCompliancePackShrinkRequest) SetAggregatorId(v string) *GetAggregateCompliancePackShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateCompliancePackShrinkRequest) SetCompliancePackId(v string) *GetAggregateCompliancePackShrinkRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateCompliancePackShrinkRequest) SetTagShrink(v string) *GetAggregateCompliancePackShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *GetAggregateCompliancePackShrinkRequest) Validate() error {
	return dara.Validate(s)
}
