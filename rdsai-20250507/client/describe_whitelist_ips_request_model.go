// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeWhitelistIpsRequest
	GetInstanceId() *string
}

type DescribeWhitelistIpsRequest struct {
	// example:
	//
	// rds_copilot****_public_cn-4****02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeWhitelistIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistIpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeWhitelistIpsRequest) SetInstanceId(v string) *DescribeWhitelistIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeWhitelistIpsRequest) Validate() error {
	return dara.Validate(s)
}
