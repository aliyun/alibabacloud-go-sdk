// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkAgentStatusDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *ClinkAgentStatusDetailRequest
	GetCno() *string
	SetEnterpriseId(v int64) *ClinkAgentStatusDetailRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *ClinkAgentStatusDetailRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkAgentStatusDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkAgentStatusDetailRequest
	GetResourceOwnerId() *int64
}

type ClinkAgentStatusDetailRequest struct {
	// 座席工号
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

func (s ClinkAgentStatusDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusDetailRequest) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusDetailRequest) GetCno() *string {
	return s.Cno
}

func (s *ClinkAgentStatusDetailRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkAgentStatusDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkAgentStatusDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkAgentStatusDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkAgentStatusDetailRequest) SetCno(v string) *ClinkAgentStatusDetailRequest {
	s.Cno = &v
	return s
}

func (s *ClinkAgentStatusDetailRequest) SetEnterpriseId(v int64) *ClinkAgentStatusDetailRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkAgentStatusDetailRequest) SetOwnerId(v int64) *ClinkAgentStatusDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkAgentStatusDetailRequest) SetResourceOwnerAccount(v string) *ClinkAgentStatusDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkAgentStatusDetailRequest) SetResourceOwnerId(v int64) *ClinkAgentStatusDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkAgentStatusDetailRequest) Validate() error {
	return dara.Validate(s)
}
