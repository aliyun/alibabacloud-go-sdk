// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAutoCcWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTime(v int32) *AddAutoCcWhitelistRequest
	GetExpireTime() *int32
	SetInstanceId(v string) *AddAutoCcWhitelistRequest
	GetInstanceId() *string
	SetWhitelist(v string) *AddAutoCcWhitelistRequest
	GetWhitelist() *string
}

type AddAutoCcWhitelistRequest struct {
	// This parameter is deprecated.
	//
	// > This parameter indicates the validity period of the IP address blacklist. By default, the traffic from the IP addresses that you add to the whitelist is always allowed. You do not need to set this parameter.
	//
	// example:
	//
	// 0
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configuration of the IP addresses that you want to add to the whitelist. The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// 	- **src**: the IP address that you want to add. This parameter is required. Data type: string.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"src":"192.XX.XX.1"},{"src":"192.XX.XX.2"}]
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s AddAutoCcWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAutoCcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *AddAutoCcWhitelistRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *AddAutoCcWhitelistRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddAutoCcWhitelistRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *AddAutoCcWhitelistRequest) SetExpireTime(v int32) *AddAutoCcWhitelistRequest {
	s.ExpireTime = &v
	return s
}

func (s *AddAutoCcWhitelistRequest) SetInstanceId(v string) *AddAutoCcWhitelistRequest {
	s.InstanceId = &v
	return s
}

func (s *AddAutoCcWhitelistRequest) SetWhitelist(v string) *AddAutoCcWhitelistRequest {
	s.Whitelist = &v
	return s
}

func (s *AddAutoCcWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
