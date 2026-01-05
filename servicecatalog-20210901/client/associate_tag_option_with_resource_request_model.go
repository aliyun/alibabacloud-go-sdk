// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTagOptionWithResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *AssociateTagOptionWithResourceRequest
	GetResourceId() *string
	SetTagOptionId(v string) *AssociateTagOptionWithResourceRequest
	GetTagOptionId() *string
}

type AssociateTagOptionWithResourceRequest struct {
	// The ID of the resource with which the tag option is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// port-bp15p96922****
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

func (s AssociateTagOptionWithResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateTagOptionWithResourceRequest) GoString() string {
	return s.String()
}

func (s *AssociateTagOptionWithResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AssociateTagOptionWithResourceRequest) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *AssociateTagOptionWithResourceRequest) SetResourceId(v string) *AssociateTagOptionWithResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *AssociateTagOptionWithResourceRequest) SetTagOptionId(v string) *AssociateTagOptionWithResourceRequest {
	s.TagOptionId = &v
	return s
}

func (s *AssociateTagOptionWithResourceRequest) Validate() error {
	return dara.Validate(s)
}
