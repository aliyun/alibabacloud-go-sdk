// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationParametersRequest
	GetApplicationId() *string
	SetComponentIdList(v []*string) *DescribeApplicationParametersRequest
	GetComponentIdList() []*string
}

type DescribeApplicationParametersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId   *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ComponentIdList []*string `json:"ComponentIdList,omitempty" xml:"ComponentIdList,omitempty" type:"Repeated"`
}

func (s DescribeApplicationParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationParametersRequest) GetComponentIdList() []*string {
	return s.ComponentIdList
}

func (s *DescribeApplicationParametersRequest) SetApplicationId(v string) *DescribeApplicationParametersRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationParametersRequest) SetComponentIdList(v []*string) *DescribeApplicationParametersRequest {
	s.ComponentIdList = v
	return s
}

func (s *DescribeApplicationParametersRequest) Validate() error {
	return dara.Validate(s)
}
