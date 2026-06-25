// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemPropertyTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeSystemPropertyTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSystemPropertyTemplatesRequest
	GetNextToken() *string
	SetTemplateIds(v []*string) *DescribeSystemPropertyTemplatesRequest
	GetTemplateIds() []*string
	SetTemplateName(v string) *DescribeSystemPropertyTemplatesRequest
	GetTemplateName() *string
}

type DescribeSystemPropertyTemplatesRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The starting position for reading. If this parameter is empty, reading starts from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6k****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The property template ID.
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeSystemPropertyTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemPropertyTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemPropertyTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSystemPropertyTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSystemPropertyTemplatesRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *DescribeSystemPropertyTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeSystemPropertyTemplatesRequest) SetMaxResults(v int32) *DescribeSystemPropertyTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesRequest) SetNextToken(v string) *DescribeSystemPropertyTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesRequest) SetTemplateIds(v []*string) *DescribeSystemPropertyTemplatesRequest {
	s.TemplateIds = v
	return s
}

func (s *DescribeSystemPropertyTemplatesRequest) SetTemplateName(v string) *DescribeSystemPropertyTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
