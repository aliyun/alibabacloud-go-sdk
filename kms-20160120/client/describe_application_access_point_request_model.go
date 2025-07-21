// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeApplicationAccessPointRequest
	GetName() *string
}

type DescribeApplicationAccessPointRequest struct {
	// The name of the AAP that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAccessPointRequest) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAccessPointRequest) SetName(v string) *DescribeApplicationAccessPointRequest {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
