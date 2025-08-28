// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotv2AllListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryRobotv2AllListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryRobotv2AllListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryRobotv2AllListRequest
	GetResourceOwnerId() *int64
}

type QueryRobotv2AllListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRobotv2AllListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotv2AllListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotv2AllListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryRobotv2AllListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryRobotv2AllListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryRobotv2AllListRequest) SetOwnerId(v int64) *QueryRobotv2AllListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotv2AllListRequest) SetResourceOwnerAccount(v string) *QueryRobotv2AllListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotv2AllListRequest) SetResourceOwnerId(v int64) *QueryRobotv2AllListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotv2AllListRequest) Validate() error {
	return dara.Validate(s)
}
