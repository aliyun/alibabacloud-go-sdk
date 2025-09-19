// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomBlockRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockIp(v string) *ModifyCustomBlockRecordRequest
	GetBlockIp() *string
	SetBound(v string) *ModifyCustomBlockRecordRequest
	GetBound() *string
	SetExpireTime(v int64) *ModifyCustomBlockRecordRequest
	GetExpireTime() *int64
	SetResourceOwnerId(v int64) *ModifyCustomBlockRecordRequest
	GetResourceOwnerId() *int64
	SetUuids(v string) *ModifyCustomBlockRecordRequest
	GetUuids() *string
}

type ModifyCustomBlockRecordRequest struct {
	// The IP address that you want to specify in the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.12.XX.XX
	BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
	// The traffic direction that you want to specify in the policy. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Bound *string `json:"Bound,omitempty" xml:"Bound,omitempty"`
	// The expiration time of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1940899881000
	ExpireTime      *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2516fe4f-adb6-45d1-87a7-90ce1213****,30746836-68d0-47f6-8b2d-c93150da****,7c3ac531-077b-46b8-8706-5c8d4e73****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ModifyCustomBlockRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomBlockRecordRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomBlockRecordRequest) GetBlockIp() *string {
	return s.BlockIp
}

func (s *ModifyCustomBlockRecordRequest) GetBound() *string {
	return s.Bound
}

func (s *ModifyCustomBlockRecordRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ModifyCustomBlockRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCustomBlockRecordRequest) GetUuids() *string {
	return s.Uuids
}

func (s *ModifyCustomBlockRecordRequest) SetBlockIp(v string) *ModifyCustomBlockRecordRequest {
	s.BlockIp = &v
	return s
}

func (s *ModifyCustomBlockRecordRequest) SetBound(v string) *ModifyCustomBlockRecordRequest {
	s.Bound = &v
	return s
}

func (s *ModifyCustomBlockRecordRequest) SetExpireTime(v int64) *ModifyCustomBlockRecordRequest {
	s.ExpireTime = &v
	return s
}

func (s *ModifyCustomBlockRecordRequest) SetResourceOwnerId(v int64) *ModifyCustomBlockRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCustomBlockRecordRequest) SetUuids(v string) *ModifyCustomBlockRecordRequest {
	s.Uuids = &v
	return s
}

func (s *ModifyCustomBlockRecordRequest) Validate() error {
	return dara.Validate(s)
}
