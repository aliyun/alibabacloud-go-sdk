// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationParametersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationParametersShrinkRequest
	GetApplicationId() *string
	SetComponentIdListShrink(v string) *DescribeApplicationParametersShrinkRequest
	GetComponentIdListShrink() *string
}

type DescribeApplicationParametersShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId         *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ComponentIdListShrink *string `json:"ComponentIdList,omitempty" xml:"ComponentIdList,omitempty"`
}

func (s DescribeApplicationParametersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationParametersShrinkRequest) GetComponentIdListShrink() *string {
	return s.ComponentIdListShrink
}

func (s *DescribeApplicationParametersShrinkRequest) SetApplicationId(v string) *DescribeApplicationParametersShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationParametersShrinkRequest) SetComponentIdListShrink(v string) *DescribeApplicationParametersShrinkRequest {
	s.ComponentIdListShrink = &v
	return s
}

func (s *DescribeApplicationParametersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
