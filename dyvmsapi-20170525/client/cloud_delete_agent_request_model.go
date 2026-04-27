// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudDeleteAgentRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudDeleteAgentRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudDeleteAgentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudDeleteAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteAgentRequest
	GetResourceOwnerId() *int64
}

type CloudDeleteAgentRequest struct {
	// 座席工号；取值3-10位整数
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
	// 7000002
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudDeleteAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudDeleteAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteAgentRequest) SetCno(v string) *CloudDeleteAgentRequest {
	s.Cno = &v
	return s
}

func (s *CloudDeleteAgentRequest) SetEnterpriseId(v int64) *CloudDeleteAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteAgentRequest) SetOwnerId(v int64) *CloudDeleteAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteAgentRequest) SetResourceOwnerAccount(v string) *CloudDeleteAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteAgentRequest) SetResourceOwnerId(v int64) *CloudDeleteAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteAgentRequest) Validate() error {
	return dara.Validate(s)
}
