// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDyplsOSSInfoForUploadFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *GetDyplsOSSInfoForUploadFileRequest
	GetBizType() *string
	SetOwnerId(v int64) *GetDyplsOSSInfoForUploadFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetDyplsOSSInfoForUploadFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetDyplsOSSInfoForUploadFileRequest
	GetResourceOwnerId() *int64
}

type GetDyplsOSSInfoForUploadFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// phone_card
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetDyplsOSSInfoForUploadFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDyplsOSSInfoForUploadFileRequest) GoString() string {
	return s.String()
}

func (s *GetDyplsOSSInfoForUploadFileRequest) GetBizType() *string {
	return s.BizType
}

func (s *GetDyplsOSSInfoForUploadFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetDyplsOSSInfoForUploadFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetDyplsOSSInfoForUploadFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetDyplsOSSInfoForUploadFileRequest) SetBizType(v string) *GetDyplsOSSInfoForUploadFileRequest {
	s.BizType = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileRequest) SetOwnerId(v int64) *GetDyplsOSSInfoForUploadFileRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileRequest) SetResourceOwnerAccount(v string) *GetDyplsOSSInfoForUploadFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileRequest) SetResourceOwnerId(v int64) *GetDyplsOSSInfoForUploadFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileRequest) Validate() error {
	return dara.Validate(s)
}
