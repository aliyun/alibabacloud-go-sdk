// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSInfoForUploadFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *GetOSSInfoForUploadFileRequest
	GetBizType() *string
	SetOwnerId(v int64) *GetOSSInfoForUploadFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetOSSInfoForUploadFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetOSSInfoForUploadFileRequest
	GetResourceOwnerId() *int64
}

type GetOSSInfoForUploadFileRequest struct {
	// Business type, default value is **fcMediaSms**.
	//
	// When creating signatures and templates, and uploading **additional materials**, this value is **fcMediaSms**.
	//
	// example:
	//
	// fcMediaSms
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetOSSInfoForUploadFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOSSInfoForUploadFileRequest) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForUploadFileRequest) GetBizType() *string {
	return s.BizType
}

func (s *GetOSSInfoForUploadFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetOSSInfoForUploadFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetOSSInfoForUploadFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetOSSInfoForUploadFileRequest) SetBizType(v string) *GetOSSInfoForUploadFileRequest {
	s.BizType = &v
	return s
}

func (s *GetOSSInfoForUploadFileRequest) SetOwnerId(v int64) *GetOSSInfoForUploadFileRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOSSInfoForUploadFileRequest) SetResourceOwnerAccount(v string) *GetOSSInfoForUploadFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetOSSInfoForUploadFileRequest) SetResourceOwnerId(v int64) *GetOSSInfoForUploadFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetOSSInfoForUploadFileRequest) Validate() error {
	return dara.Validate(s)
}
