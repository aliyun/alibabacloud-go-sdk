// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisAssociateTagOptionFromResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *DisAssociateTagOptionFromResourceRequest
	GetResourceId() *string
	SetTagOptionId(v string) *DisAssociateTagOptionFromResourceRequest
	GetTagOptionId() *string
}

type DisAssociateTagOptionFromResourceRequest struct {
	// The ID of the resource with which the tag option is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp14katy3d****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the tag option.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-bp1u6mdf3d****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
}

func (s DisAssociateTagOptionFromResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DisAssociateTagOptionFromResourceRequest) GoString() string {
	return s.String()
}

func (s *DisAssociateTagOptionFromResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DisAssociateTagOptionFromResourceRequest) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *DisAssociateTagOptionFromResourceRequest) SetResourceId(v string) *DisAssociateTagOptionFromResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *DisAssociateTagOptionFromResourceRequest) SetTagOptionId(v string) *DisAssociateTagOptionFromResourceRequest {
	s.TagOptionId = &v
	return s
}

func (s *DisAssociateTagOptionFromResourceRequest) Validate() error {
	return dara.Validate(s)
}
