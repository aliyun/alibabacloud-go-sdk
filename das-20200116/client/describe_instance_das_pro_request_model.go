// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDasProRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceDasProRequest
	GetInstanceId() *string
}

type DescribeInstanceDasProRequest struct {
	// The database instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceDasProRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDasProRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDasProRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceDasProRequest) SetInstanceId(v string) *DescribeInstanceDasProRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDasProRequest) Validate() error {
	return dara.Validate(s)
}
