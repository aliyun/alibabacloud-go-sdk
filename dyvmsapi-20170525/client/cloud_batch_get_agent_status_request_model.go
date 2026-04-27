// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudBatchGetAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCnos(v string) *CloudBatchGetAgentStatusRequest
	GetCnos() *string
	SetEnterpriseId(v int64) *CloudBatchGetAgentStatusRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudBatchGetAgentStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudBatchGetAgentStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudBatchGetAgentStatusRequest
	GetResourceOwnerId() *int64
}

type CloudBatchGetAgentStatusRequest struct {
	// 座席号列表；座席号之间用`,`分隔，如：cnos=2000,20001 最多支持100个座席
	//
	// This parameter is required.
	//
	// example:
	//
	// 2000,2001
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
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

func (s CloudBatchGetAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchGetAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *CloudBatchGetAgentStatusRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudBatchGetAgentStatusRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudBatchGetAgentStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudBatchGetAgentStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudBatchGetAgentStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudBatchGetAgentStatusRequest) SetCnos(v string) *CloudBatchGetAgentStatusRequest {
	s.Cnos = &v
	return s
}

func (s *CloudBatchGetAgentStatusRequest) SetEnterpriseId(v int64) *CloudBatchGetAgentStatusRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudBatchGetAgentStatusRequest) SetOwnerId(v int64) *CloudBatchGetAgentStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudBatchGetAgentStatusRequest) SetResourceOwnerAccount(v string) *CloudBatchGetAgentStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudBatchGetAgentStatusRequest) SetResourceOwnerId(v int64) *CloudBatchGetAgentStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudBatchGetAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}
