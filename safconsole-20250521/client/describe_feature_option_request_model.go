// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFeatureOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureTemplate(v string) *DescribeFeatureOptionRequest
	GetFeatureTemplate() *string
}

type DescribeFeatureOptionRequest struct {
	// Feature template.
	//
	// This parameter is required.
	//
	// example:
	//
	// FINANCE_51
	FeatureTemplate *string `json:"FeatureTemplate,omitempty" xml:"FeatureTemplate,omitempty"`
}

func (s DescribeFeatureOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureOptionRequest) GoString() string {
	return s.String()
}

func (s *DescribeFeatureOptionRequest) GetFeatureTemplate() *string {
	return s.FeatureTemplate
}

func (s *DescribeFeatureOptionRequest) SetFeatureTemplate(v string) *DescribeFeatureOptionRequest {
	s.FeatureTemplate = &v
	return s
}

func (s *DescribeFeatureOptionRequest) Validate() error {
	return dara.Validate(s)
}
