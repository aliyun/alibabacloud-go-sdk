// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomBlockRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockIp(v string) *DisableCustomBlockRecordRequest
	GetBlockIp() *string
	SetBound(v string) *DisableCustomBlockRecordRequest
	GetBound() *string
	SetResourceOwnerId(v int64) *DisableCustomBlockRecordRequest
	GetResourceOwnerId() *int64
}

type DisableCustomBlockRecordRequest struct {
	// The IP address that is specified in the policy.
	//
	// > You can call the [DescribeCustomBlockRecords](~~DescribeCustomBlockRecords~~) operation to query the IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 62.233.XX.XX
	BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
	// The traffic direction that is specified in the policy. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Bound           *string `json:"Bound,omitempty" xml:"Bound,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DisableCustomBlockRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomBlockRecordRequest) GoString() string {
	return s.String()
}

func (s *DisableCustomBlockRecordRequest) GetBlockIp() *string {
	return s.BlockIp
}

func (s *DisableCustomBlockRecordRequest) GetBound() *string {
	return s.Bound
}

func (s *DisableCustomBlockRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableCustomBlockRecordRequest) SetBlockIp(v string) *DisableCustomBlockRecordRequest {
	s.BlockIp = &v
	return s
}

func (s *DisableCustomBlockRecordRequest) SetBound(v string) *DisableCustomBlockRecordRequest {
	s.Bound = &v
	return s
}

func (s *DisableCustomBlockRecordRequest) SetResourceOwnerId(v int64) *DisableCustomBlockRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableCustomBlockRecordRequest) Validate() error {
	return dara.Validate(s)
}
