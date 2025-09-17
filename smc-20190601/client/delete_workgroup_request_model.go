// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkgroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteWorkgroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteWorkgroupRequest
	GetResourceOwnerAccount() *string
	SetWorkgroupId(v string) *DeleteWorkgroupRequest
	GetWorkgroupId() *string
}

type DeleteWorkgroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the workgroup.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-***
	WorkgroupId *string `json:"WorkgroupId,omitempty" xml:"WorkgroupId,omitempty"`
}

func (s DeleteWorkgroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkgroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkgroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteWorkgroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteWorkgroupRequest) GetWorkgroupId() *string {
	return s.WorkgroupId
}

func (s *DeleteWorkgroupRequest) SetOwnerId(v int64) *DeleteWorkgroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteWorkgroupRequest) SetResourceOwnerAccount(v string) *DeleteWorkgroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteWorkgroupRequest) SetWorkgroupId(v string) *DeleteWorkgroupRequest {
	s.WorkgroupId = &v
	return s
}

func (s *DeleteWorkgroupRequest) Validate() error {
	return dara.Validate(s)
}
