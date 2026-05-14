// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *ClinkDescribeClientRequest
	GetCno() *string
	SetEnterpriseId(v int64) *ClinkDescribeClientRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *ClinkDescribeClientRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkDescribeClientRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDescribeClientRequest
	GetResourceOwnerId() *int64
}

type ClinkDescribeClientRequest struct {
	// 座席号
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkDescribeClientRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeClientRequest) GoString() string {
	return s.String()
}

func (s *ClinkDescribeClientRequest) GetCno() *string {
	return s.Cno
}

func (s *ClinkDescribeClientRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDescribeClientRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDescribeClientRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDescribeClientRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDescribeClientRequest) SetCno(v string) *ClinkDescribeClientRequest {
	s.Cno = &v
	return s
}

func (s *ClinkDescribeClientRequest) SetEnterpriseId(v int64) *ClinkDescribeClientRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDescribeClientRequest) SetOwnerId(v int64) *ClinkDescribeClientRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDescribeClientRequest) SetResourceOwnerAccount(v string) *ClinkDescribeClientRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDescribeClientRequest) SetResourceOwnerId(v int64) *ClinkDescribeClientRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDescribeClientRequest) Validate() error {
	return dara.Validate(s)
}
