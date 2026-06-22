// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *DescribeImageCriteriaRequest
	GetValue() *string
}

type DescribeImageCriteriaRequest struct {
	// The fuzzy match value that you enter when querying images.
	//
	// > This parameter supports fuzzy matching by image ID, image tag, image instance ID, image repository name, image repository ID, image repository namespace, image region, image digest, or image repository type.
	//
	// example:
	//
	// 525
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeImageCriteriaRequest) SetValue(v string) *DescribeImageCriteriaRequest {
	s.Value = &v
	return s
}

func (s *DescribeImageCriteriaRequest) Validate() error {
	return dara.Validate(s)
}
