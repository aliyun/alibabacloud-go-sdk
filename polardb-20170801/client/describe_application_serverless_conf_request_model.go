// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationServerlessConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationServerlessConfRequest
	GetApplicationId() *string
}

type DescribeApplicationServerlessConfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s DescribeApplicationServerlessConfRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationServerlessConfRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationServerlessConfRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationServerlessConfRequest) SetApplicationId(v string) *DescribeApplicationServerlessConfRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationServerlessConfRequest) Validate() error {
	return dara.Validate(s)
}
