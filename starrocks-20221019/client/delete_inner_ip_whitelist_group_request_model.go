// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInnerIpWhitelistGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInnerIpWhitelistGroupId(v string) *DeleteInnerIpWhitelistGroupRequest
	GetInnerIpWhitelistGroupId() *string
	SetInstanceId(v string) *DeleteInnerIpWhitelistGroupRequest
	GetInstanceId() *string
}

type DeleteInnerIpWhitelistGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-g1
	InnerIpWhitelistGroupId *string `json:"InnerIpWhitelistGroupId,omitempty" xml:"InnerIpWhitelistGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-0104730e9de40215
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInnerIpWhitelistGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInnerIpWhitelistGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteInnerIpWhitelistGroupRequest) GetInnerIpWhitelistGroupId() *string {
	return s.InnerIpWhitelistGroupId
}

func (s *DeleteInnerIpWhitelistGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInnerIpWhitelistGroupRequest) SetInnerIpWhitelistGroupId(v string) *DeleteInnerIpWhitelistGroupRequest {
	s.InnerIpWhitelistGroupId = &v
	return s
}

func (s *DeleteInnerIpWhitelistGroupRequest) SetInstanceId(v string) *DeleteInnerIpWhitelistGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInnerIpWhitelistGroupRequest) Validate() error {
	return dara.Validate(s)
}
