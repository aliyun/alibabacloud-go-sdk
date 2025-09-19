// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBruteForceRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockIp(v string) *DisableBruteForceRecordRequest
	GetBlockIp() *string
	SetBound(v string) *DisableBruteForceRecordRequest
	GetBound() *string
	SetId(v int64) *DisableBruteForceRecordRequest
	GetId() *int64
	SetPort(v string) *DisableBruteForceRecordRequest
	GetPort() *string
	SetResourceOwnerId(v int64) *DisableBruteForceRecordRequest
	GetResourceOwnerId() *int64
	SetUuid(v string) *DisableBruteForceRecordRequest
	GetUuid() *string
}

type DisableBruteForceRecordRequest struct {
	// The IP address that you want to specify in the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8.210.XX.XX
	BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
	// The traffic direction that you want to specify in the policy. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// example:
	//
	// in
	Bound *string `json:"Bound,omitempty" xml:"Bound,omitempty"`
	// The ID of the IP address blocking policy.
	//
	// > You can call the [DescribeBruteForceRecords](~~DescribeBruteForceRecords~~) operation to query the policy ID.
	//
	// example:
	//
	// 114166XX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The port number.
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The UUID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbb9aa80-a8d1-443c-9ff0-2c36cd39****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DisableBruteForceRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableBruteForceRecordRequest) GoString() string {
	return s.String()
}

func (s *DisableBruteForceRecordRequest) GetBlockIp() *string {
	return s.BlockIp
}

func (s *DisableBruteForceRecordRequest) GetBound() *string {
	return s.Bound
}

func (s *DisableBruteForceRecordRequest) GetId() *int64 {
	return s.Id
}

func (s *DisableBruteForceRecordRequest) GetPort() *string {
	return s.Port
}

func (s *DisableBruteForceRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableBruteForceRecordRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DisableBruteForceRecordRequest) SetBlockIp(v string) *DisableBruteForceRecordRequest {
	s.BlockIp = &v
	return s
}

func (s *DisableBruteForceRecordRequest) SetBound(v string) *DisableBruteForceRecordRequest {
	s.Bound = &v
	return s
}

func (s *DisableBruteForceRecordRequest) SetId(v int64) *DisableBruteForceRecordRequest {
	s.Id = &v
	return s
}

func (s *DisableBruteForceRecordRequest) SetPort(v string) *DisableBruteForceRecordRequest {
	s.Port = &v
	return s
}

func (s *DisableBruteForceRecordRequest) SetResourceOwnerId(v int64) *DisableBruteForceRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableBruteForceRecordRequest) SetUuid(v string) *DisableBruteForceRecordRequest {
	s.Uuid = &v
	return s
}

func (s *DisableBruteForceRecordRequest) Validate() error {
	return dara.Validate(s)
}
