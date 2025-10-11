// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationAttributeRequest
	GetApplicationId() *string
}

type DescribeApplicationAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s DescribeApplicationAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationAttributeRequest) SetApplicationId(v string) *DescribeApplicationAttributeRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationAttributeRequest) Validate() error {
	return dara.Validate(s)
}
