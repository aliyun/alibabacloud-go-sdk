// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoCcWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteAutoCcWhitelistRequest
	GetInstanceId() *string
	SetWhitelist(v string) *DeleteAutoCcWhitelistRequest
	GetWhitelist() *string
}

type DeleteAutoCcWhitelistRequest struct {
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
	// The IP addresses that you want to manage. This parameter is a JSON string. This parameter is a JSON string. The string contains the following field:
	//
	// 	- **src**: the IP address. This field is required and must be of the string type.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"src":"1.1.1.1"},{"src":"2.2.2.2"}]
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s DeleteAutoCcWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoCcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcWhitelistRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAutoCcWhitelistRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *DeleteAutoCcWhitelistRequest) SetInstanceId(v string) *DeleteAutoCcWhitelistRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAutoCcWhitelistRequest) SetWhitelist(v string) *DeleteAutoCcWhitelistRequest {
	s.Whitelist = &v
	return s
}

func (s *DeleteAutoCcWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
