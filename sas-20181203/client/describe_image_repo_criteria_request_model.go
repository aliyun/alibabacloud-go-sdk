// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRepoCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *DescribeImageRepoCriteriaRequest
	GetValue() *string
}

type DescribeImageRepoCriteriaRequest struct {
	// The search value for image repositories.
	//
	// > This parameter supports fuzzy match for image IDs, tags, image instance IDs, repository names, repository namespaces, repository IDs, repository regions, digests, and repository types.
	//
	// example:
	//
	// 2.0.2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageRepoCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeImageRepoCriteriaRequest) SetValue(v string) *DescribeImageRepoCriteriaRequest {
	s.Value = &v
	return s
}

func (s *DescribeImageRepoCriteriaRequest) Validate() error {
	return dara.Validate(s)
}
