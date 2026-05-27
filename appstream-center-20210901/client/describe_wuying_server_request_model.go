// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWuyingServerId(v string) *DescribeWuyingServerRequest
	GetWuyingServerId() *string
}

type DescribeWuyingServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aig-bp1234567890abcde
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s DescribeWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *DescribeWuyingServerRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *DescribeWuyingServerRequest) SetWuyingServerId(v string) *DescribeWuyingServerRequest {
	s.WuyingServerId = &v
	return s
}

func (s *DescribeWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}
