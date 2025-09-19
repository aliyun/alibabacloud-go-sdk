// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyUsageTopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribePropertyUsageTopRequest
	GetType() *string
}

type DescribePropertyUsageTopRequest struct {
	// The type of the asset fingerprint. Valid value:
	//
	// 	- **port**: port
	//
	// 	- **process**: process
	//
	// 	- **software**: software
	//
	// 	- **user**: account
	//
	// 	- **sca**: middleware
	//
	// This parameter is required.
	//
	// example:
	//
	// port
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePropertyUsageTopRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUsageTopRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyUsageTopRequest) GetType() *string {
	return s.Type
}

func (s *DescribePropertyUsageTopRequest) SetType(v string) *DescribePropertyUsageTopRequest {
	s.Type = &v
	return s
}

func (s *DescribePropertyUsageTopRequest) Validate() error {
	return dara.Validate(s)
}
