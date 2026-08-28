// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationSSLRequest
	GetApplicationId() *string
}

type DescribeApplicationSSLRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s DescribeApplicationSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSSLRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationSSLRequest) SetApplicationId(v string) *DescribeApplicationSSLRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationSSLRequest) Validate() error {
	return dara.Validate(s)
}
