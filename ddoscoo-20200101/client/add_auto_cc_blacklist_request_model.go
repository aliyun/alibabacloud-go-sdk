// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAutoCcBlacklistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlacklist(v string) *AddAutoCcBlacklistRequest
	GetBlacklist() *string
	SetExpireTime(v int32) *AddAutoCcBlacklistRequest
	GetExpireTime() *int32
	SetInstanceId(v string) *AddAutoCcBlacklistRequest
	GetInstanceId() *string
}

type AddAutoCcBlacklistRequest struct {
	// The IP addresses that you want to manage. This parameter is a JSON string. The string contains the following field:
	//
	// 	- **src**: the IP address. This field is required and must be of the STRING type.
	//
	// >  You can manually add up to 2,000 IP addresses to the IP address blacklist. Separate multiple IP addresses with spaces or line breaks.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"src":"198.51.XX.XX"},{"src":"198.52.XX.XX"}]
	Blacklist *string `json:"Blacklist,omitempty" xml:"Blacklist,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 300
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddAutoCcBlacklistRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAutoCcBlacklistRequest) GoString() string {
	return s.String()
}

func (s *AddAutoCcBlacklistRequest) GetBlacklist() *string {
	return s.Blacklist
}

func (s *AddAutoCcBlacklistRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *AddAutoCcBlacklistRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddAutoCcBlacklistRequest) SetBlacklist(v string) *AddAutoCcBlacklistRequest {
	s.Blacklist = &v
	return s
}

func (s *AddAutoCcBlacklistRequest) SetExpireTime(v int32) *AddAutoCcBlacklistRequest {
	s.ExpireTime = &v
	return s
}

func (s *AddAutoCcBlacklistRequest) SetInstanceId(v string) *AddAutoCcBlacklistRequest {
	s.InstanceId = &v
	return s
}

func (s *AddAutoCcBlacklistRequest) Validate() error {
	return dara.Validate(s)
}
