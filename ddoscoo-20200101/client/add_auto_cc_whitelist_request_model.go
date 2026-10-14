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
	// **[Deprecated]*	- This parameter is deprecated.
	//
	// > This parameter specifies the validity period of blacklist IP addresses. Whitelist IP addresses are permanently valid by default. You do not need to set this parameter.
	//
	// example:
	//
	// 0
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance for which you want to add whitelist IP addresses.
	//
	// > You can invoke [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS Pro or Anti-DDoS Premium instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configuration of the whitelist IP addresses to add. This parameter is a string that is converted from a JSON array. Each element in the JSON array is a whitelist IP struct that contains the following field:
	//
	// - **src**: String | Required | The IP address to add to the whitelist.
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
