// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoCcBlacklistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlacklist(v string) *DeleteAutoCcBlacklistRequest
	GetBlacklist() *string
	SetInstanceId(v string) *DeleteAutoCcBlacklistRequest
	GetInstanceId() *string
	SetQueryType(v string) *DeleteAutoCcBlacklistRequest
	GetQueryType() *string
}

type DeleteAutoCcBlacklistRequest struct {
	// The IP addresses that you want to manage. This parameter is a JSON string. The string contains the following fields:
	//
	// 	- **src**: the IP address. This field is required and must be of the STRING type.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"src":"198.51.XX.XX"},{"src":"198.52.XX.XX"}]
	Blacklist *string `json:"Blacklist,omitempty" xml:"Blacklist,omitempty"`
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
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
}

func (s DeleteAutoCcBlacklistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoCcBlacklistRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcBlacklistRequest) GetBlacklist() *string {
	return s.Blacklist
}

func (s *DeleteAutoCcBlacklistRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAutoCcBlacklistRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DeleteAutoCcBlacklistRequest) SetBlacklist(v string) *DeleteAutoCcBlacklistRequest {
	s.Blacklist = &v
	return s
}

func (s *DeleteAutoCcBlacklistRequest) SetInstanceId(v string) *DeleteAutoCcBlacklistRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAutoCcBlacklistRequest) SetQueryType(v string) *DeleteAutoCcBlacklistRequest {
	s.QueryType = &v
	return s
}

func (s *DeleteAutoCcBlacklistRequest) Validate() error {
	return dara.Validate(s)
}
