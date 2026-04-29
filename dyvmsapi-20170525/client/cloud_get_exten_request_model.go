// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetExtenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudGetExtenRequest
	GetEnterpriseId() *int64
	SetExten(v string) *CloudGetExtenRequest
	GetExten() *string
	SetOwnerId(v int64) *CloudGetExtenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudGetExtenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudGetExtenRequest
	GetResourceOwnerId() *int64
}

type CloudGetExtenRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 分机号，3-11位
	//
	// This parameter is required.
	//
	// example:
	//
	// 9000
	Exten                *string `json:"Exten,omitempty" xml:"Exten,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudGetExtenRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudGetExtenRequest) GoString() string {
	return s.String()
}

func (s *CloudGetExtenRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetExtenRequest) GetExten() *string {
	return s.Exten
}

func (s *CloudGetExtenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudGetExtenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudGetExtenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudGetExtenRequest) SetEnterpriseId(v int64) *CloudGetExtenRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetExtenRequest) SetExten(v string) *CloudGetExtenRequest {
	s.Exten = &v
	return s
}

func (s *CloudGetExtenRequest) SetOwnerId(v int64) *CloudGetExtenRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudGetExtenRequest) SetResourceOwnerAccount(v string) *CloudGetExtenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudGetExtenRequest) SetResourceOwnerId(v int64) *CloudGetExtenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudGetExtenRequest) Validate() error {
	return dara.Validate(s)
}
