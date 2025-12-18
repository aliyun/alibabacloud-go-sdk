// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *GetCompliancePackShrinkRequest
	GetCompliancePackId() *string
	SetTagShrink(v string) *GetCompliancePackShrinkRequest
	GetTagShrink() *string
}

type GetCompliancePackShrinkRequest struct {
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-a8a8626622af0082****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// Deprecated
	//
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetCompliancePackShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCompliancePackShrinkRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetCompliancePackShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *GetCompliancePackShrinkRequest) SetCompliancePackId(v string) *GetCompliancePackShrinkRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetCompliancePackShrinkRequest) SetTagShrink(v string) *GetCompliancePackShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *GetCompliancePackShrinkRequest) Validate() error {
	return dara.Validate(s)
}
