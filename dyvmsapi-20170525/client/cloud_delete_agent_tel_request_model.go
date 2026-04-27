// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteAgentTelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudDeleteAgentTelRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudDeleteAgentTelRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudDeleteAgentTelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudDeleteAgentTelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteAgentTelRequest
	GetResourceOwnerId() *int64
	SetTel(v string) *CloudDeleteAgentTelRequest
	GetTel() *string
}

type CloudDeleteAgentTelRequest struct {
	// 座席工号；3-10位数字
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
	// 座席绑定电话；固话带区号，手机不加0；固话带分机的以\\"-\\"分隔
	//
	// This parameter is required.
	//
	// example:
	//
	// 22223333
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
}

func (s CloudDeleteAgentTelRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentTelRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentTelRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudDeleteAgentTelRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteAgentTelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteAgentTelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteAgentTelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteAgentTelRequest) GetTel() *string {
	return s.Tel
}

func (s *CloudDeleteAgentTelRequest) SetCno(v string) *CloudDeleteAgentTelRequest {
	s.Cno = &v
	return s
}

func (s *CloudDeleteAgentTelRequest) SetEnterpriseId(v int64) *CloudDeleteAgentTelRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteAgentTelRequest) SetOwnerId(v int64) *CloudDeleteAgentTelRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteAgentTelRequest) SetResourceOwnerAccount(v string) *CloudDeleteAgentTelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteAgentTelRequest) SetResourceOwnerId(v int64) *CloudDeleteAgentTelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteAgentTelRequest) SetTel(v string) *CloudDeleteAgentTelRequest {
	s.Tel = &v
	return s
}

func (s *CloudDeleteAgentTelRequest) Validate() error {
	return dara.Validate(s)
}
