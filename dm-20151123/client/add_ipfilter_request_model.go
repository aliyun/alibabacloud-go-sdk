// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpfilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpAddress(v string) *AddIpfilterRequest
	GetIpAddress() *string
	SetOwnerId(v int64) *AddIpfilterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddIpfilterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddIpfilterRequest
	GetResourceOwnerId() *int64
}

type AddIpfilterRequest struct {
	// IP Address/IP Range/IP Segment
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx.xxx.xxx.xxx
	//
	// xxx.xxx.xxx.xxx-xxx.xxx.xxx.xxx
	//
	// xxx.xxx.xxx.xxx/xxx
	IpAddress            *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddIpfilterRequest) String() string {
	return dara.Prettify(s)
}

func (s AddIpfilterRequest) GoString() string {
	return s.String()
}

func (s *AddIpfilterRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *AddIpfilterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddIpfilterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddIpfilterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddIpfilterRequest) SetIpAddress(v string) *AddIpfilterRequest {
	s.IpAddress = &v
	return s
}

func (s *AddIpfilterRequest) SetOwnerId(v int64) *AddIpfilterRequest {
	s.OwnerId = &v
	return s
}

func (s *AddIpfilterRequest) SetResourceOwnerAccount(v string) *AddIpfilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddIpfilterRequest) SetResourceOwnerId(v int64) *AddIpfilterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddIpfilterRequest) Validate() error {
	return dara.Validate(s)
}
