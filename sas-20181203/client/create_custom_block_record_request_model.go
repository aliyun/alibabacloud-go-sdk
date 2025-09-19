// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomBlockRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockIp(v string) *CreateCustomBlockRecordRequest
	GetBlockIp() *string
	SetBound(v string) *CreateCustomBlockRecordRequest
	GetBound() *string
	SetExpireTime(v int64) *CreateCustomBlockRecordRequest
	GetExpireTime() *int64
	SetResourceOwnerId(v int64) *CreateCustomBlockRecordRequest
	GetResourceOwnerId() *int64
	SetUuids(v string) *CreateCustomBlockRecordRequest
	GetUuids() *string
}

type CreateCustomBlockRecordRequest struct {
	// The IP address that you want to specify in the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.xx.xx
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
	// in
	Bound *string `json:"Bound,omitempty" xml:"Bound,omitempty"`
	// The expiration time of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1859094550000
	ExpireTime      *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The UUIDs of the servers. Separate multiple UUIDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 71c846d6-5c84-4714-acfc-58265bc3****,5013b5e8-1613-43a8-b4de-651db318****,df53f0ad-b3ba-4fe0-9ec7-f42a2ae2****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s CreateCustomBlockRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomBlockRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomBlockRecordRequest) GetBlockIp() *string {
	return s.BlockIp
}

func (s *CreateCustomBlockRecordRequest) GetBound() *string {
	return s.Bound
}

func (s *CreateCustomBlockRecordRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *CreateCustomBlockRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCustomBlockRecordRequest) GetUuids() *string {
	return s.Uuids
}

func (s *CreateCustomBlockRecordRequest) SetBlockIp(v string) *CreateCustomBlockRecordRequest {
	s.BlockIp = &v
	return s
}

func (s *CreateCustomBlockRecordRequest) SetBound(v string) *CreateCustomBlockRecordRequest {
	s.Bound = &v
	return s
}

func (s *CreateCustomBlockRecordRequest) SetExpireTime(v int64) *CreateCustomBlockRecordRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateCustomBlockRecordRequest) SetResourceOwnerId(v int64) *CreateCustomBlockRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCustomBlockRecordRequest) SetUuids(v string) *CreateCustomBlockRecordRequest {
	s.Uuids = &v
	return s
}

func (s *CreateCustomBlockRecordRequest) Validate() error {
	return dara.Validate(s)
}
