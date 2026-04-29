// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudDeleteAgentGroupRequest
	GetEnterpriseId() *int64
	SetGno(v string) *CloudDeleteAgentGroupRequest
	GetGno() *string
	SetOwnerId(v int64) *CloudDeleteAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudDeleteAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteAgentGroupRequest
	GetResourceOwnerId() *int64
}

type CloudDeleteAgentGroupRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 外呼组编号；说明：删除此gno对应的外呼组信息
	//
	// example:
	//
	// WH123
	Gno                  *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudDeleteAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentGroupRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteAgentGroupRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudDeleteAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteAgentGroupRequest) SetEnterpriseId(v int64) *CloudDeleteAgentGroupRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteAgentGroupRequest) SetGno(v string) *CloudDeleteAgentGroupRequest {
	s.Gno = &v
	return s
}

func (s *CloudDeleteAgentGroupRequest) SetOwnerId(v int64) *CloudDeleteAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteAgentGroupRequest) SetResourceOwnerAccount(v string) *CloudDeleteAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteAgentGroupRequest) SetResourceOwnerId(v int64) *CloudDeleteAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
