// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualificationOssInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *GetQualificationOssInfoRequest
	GetBizType() *string
	SetOwnerId(v int64) *GetQualificationOssInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetQualificationOssInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetQualificationOssInfoRequest
	GetResourceOwnerId() *int64
}

type GetQualificationOssInfoRequest struct {
	// 业务，非空
	//
	// This parameter is required.
	//
	// example:
	//
	// dysms
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetQualificationOssInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualificationOssInfoRequest) GoString() string {
	return s.String()
}

func (s *GetQualificationOssInfoRequest) GetBizType() *string {
	return s.BizType
}

func (s *GetQualificationOssInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetQualificationOssInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetQualificationOssInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetQualificationOssInfoRequest) SetBizType(v string) *GetQualificationOssInfoRequest {
	s.BizType = &v
	return s
}

func (s *GetQualificationOssInfoRequest) SetOwnerId(v int64) *GetQualificationOssInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetQualificationOssInfoRequest) SetResourceOwnerAccount(v string) *GetQualificationOssInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetQualificationOssInfoRequest) SetResourceOwnerId(v int64) *GetQualificationOssInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetQualificationOssInfoRequest) Validate() error {
	return dara.Validate(s)
}
