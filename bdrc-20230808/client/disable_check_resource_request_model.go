// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCheckResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceArn(v string) *DisableCheckResourceRequest
	GetResourceArn() *string
}

type DisableCheckResourceRequest struct {
	// Unique resource identity
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s DisableCheckResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCheckResourceRequest) GoString() string {
	return s.String()
}

func (s *DisableCheckResourceRequest) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *DisableCheckResourceRequest) SetResourceArn(v string) *DisableCheckResourceRequest {
	s.ResourceArn = &v
	return s
}

func (s *DisableCheckResourceRequest) Validate() error {
	return dara.Validate(s)
}
