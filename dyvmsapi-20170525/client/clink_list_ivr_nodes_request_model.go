// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListIvrNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkListIvrNodesRequest
	GetEnterpriseId() *int64
	SetIvrName(v string) *ClinkListIvrNodesRequest
	GetIvrName() *string
	SetOwnerId(v int64) *ClinkListIvrNodesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkListIvrNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListIvrNodesRequest
	GetResourceOwnerId() *int64
}

type ClinkListIvrNodesRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 语音导航名称
	//
	// This parameter is required.
	//
	// example:
	//
	// IvrName
	IvrName              *string `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkListIvrNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrNodesRequest) GoString() string {
	return s.String()
}

func (s *ClinkListIvrNodesRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListIvrNodesRequest) GetIvrName() *string {
	return s.IvrName
}

func (s *ClinkListIvrNodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListIvrNodesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListIvrNodesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListIvrNodesRequest) SetEnterpriseId(v int64) *ClinkListIvrNodesRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListIvrNodesRequest) SetIvrName(v string) *ClinkListIvrNodesRequest {
	s.IvrName = &v
	return s
}

func (s *ClinkListIvrNodesRequest) SetOwnerId(v int64) *ClinkListIvrNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListIvrNodesRequest) SetResourceOwnerAccount(v string) *ClinkListIvrNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListIvrNodesRequest) SetResourceOwnerId(v int64) *ClinkListIvrNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListIvrNodesRequest) Validate() error {
	return dara.Validate(s)
}
