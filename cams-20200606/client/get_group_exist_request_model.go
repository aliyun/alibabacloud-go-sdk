// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupExistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *GetGroupExistRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *GetGroupExistRequest
	GetBizExtend() map[string]interface{}
	SetGroupName(v string) *GetGroupExistRequest
	GetGroupName() *string
	SetOwnerId(v int64) *GetGroupExistRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetGroupExistRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetGroupExistRequest
	GetResourceOwnerId() *int64
}

type GetGroupExistRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aaa
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetGroupExistRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupExistRequest) GoString() string {
	return s.String()
}

func (s *GetGroupExistRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *GetGroupExistRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *GetGroupExistRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGroupExistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetGroupExistRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetGroupExistRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetGroupExistRequest) SetBizCode(v string) *GetGroupExistRequest {
	s.BizCode = &v
	return s
}

func (s *GetGroupExistRequest) SetBizExtend(v map[string]interface{}) *GetGroupExistRequest {
	s.BizExtend = v
	return s
}

func (s *GetGroupExistRequest) SetGroupName(v string) *GetGroupExistRequest {
	s.GroupName = &v
	return s
}

func (s *GetGroupExistRequest) SetOwnerId(v int64) *GetGroupExistRequest {
	s.OwnerId = &v
	return s
}

func (s *GetGroupExistRequest) SetResourceOwnerAccount(v string) *GetGroupExistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetGroupExistRequest) SetResourceOwnerId(v int64) *GetGroupExistRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetGroupExistRequest) Validate() error {
	return dara.Validate(s)
}
