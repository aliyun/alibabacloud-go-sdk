// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRebootStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuids(v string) *DescribeInstanceRebootStatusRequest
	GetUuids() *string
}

type DescribeInstanceRebootStatusRequest struct {
	// The UUIDs of the servers that you restart. Separate multiple UUIDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// d77f7802-4f0a-4221-ab02-4d999e****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeInstanceRebootStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRebootStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRebootStatusRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeInstanceRebootStatusRequest) SetUuids(v string) *DescribeInstanceRebootStatusRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeInstanceRebootStatusRequest) Validate() error {
	return dara.Validate(s)
}
