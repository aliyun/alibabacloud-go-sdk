// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ListAllGroupsRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *ListAllGroupsRequest
	GetBizExtend() map[string]interface{}
	SetOwnerId(v int64) *ListAllGroupsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAllGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAllGroupsRequest
	GetResourceOwnerId() *int64
}

type ListAllGroupsRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtend            map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	OwnerId              *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAllGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAllGroupsRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ListAllGroupsRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *ListAllGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAllGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAllGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAllGroupsRequest) SetBizCode(v string) *ListAllGroupsRequest {
	s.BizCode = &v
	return s
}

func (s *ListAllGroupsRequest) SetBizExtend(v map[string]interface{}) *ListAllGroupsRequest {
	s.BizExtend = v
	return s
}

func (s *ListAllGroupsRequest) SetOwnerId(v int64) *ListAllGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAllGroupsRequest) SetResourceOwnerAccount(v string) *ListAllGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAllGroupsRequest) SetResourceOwnerId(v int64) *ListAllGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAllGroupsRequest) Validate() error {
	return dara.Validate(s)
}
