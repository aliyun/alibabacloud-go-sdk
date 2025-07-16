// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DescribeCustomLogConfigRequest
	GetConfigId() *string
}

type DescribeCustomLogConfigRequest struct {
	// The ID of the custom configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s DescribeCustomLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomLogConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCustomLogConfigRequest) SetConfigId(v string) *DescribeCustomLogConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DescribeCustomLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
