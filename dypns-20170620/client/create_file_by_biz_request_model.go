// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileByBizRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateFileByBizRequest
	GetBizId() *string
	SetBizType(v string) *CreateFileByBizRequest
	GetBizType() *string
	SetFileType(v string) *CreateFileByBizRequest
	GetFileType() *string
	SetOwnerId(v int64) *CreateFileByBizRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateFileByBizRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateFileByBizRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFileByBizRequest
	GetResourceOwnerId() *int64
	SetSaveOssFileName(v string) *CreateFileByBizRequest
	GetSaveOssFileName() *string
	SetUserViewFileName(v string) *CreateFileByBizRequest
	GetUserViewFileName() *string
}

type CreateFileByBizRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	FileType             *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SaveOssFileName *string `json:"SaveOssFileName,omitempty" xml:"SaveOssFileName,omitempty"`
	// This parameter is required.
	UserViewFileName *string `json:"UserViewFileName,omitempty" xml:"UserViewFileName,omitempty"`
}

func (s CreateFileByBizRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileByBizRequest) GoString() string {
	return s.String()
}

func (s *CreateFileByBizRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateFileByBizRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateFileByBizRequest) GetFileType() *string {
	return s.FileType
}

func (s *CreateFileByBizRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFileByBizRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateFileByBizRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFileByBizRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFileByBizRequest) GetSaveOssFileName() *string {
	return s.SaveOssFileName
}

func (s *CreateFileByBizRequest) GetUserViewFileName() *string {
	return s.UserViewFileName
}

func (s *CreateFileByBizRequest) SetBizId(v string) *CreateFileByBizRequest {
	s.BizId = &v
	return s
}

func (s *CreateFileByBizRequest) SetBizType(v string) *CreateFileByBizRequest {
	s.BizType = &v
	return s
}

func (s *CreateFileByBizRequest) SetFileType(v string) *CreateFileByBizRequest {
	s.FileType = &v
	return s
}

func (s *CreateFileByBizRequest) SetOwnerId(v int64) *CreateFileByBizRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFileByBizRequest) SetRemark(v string) *CreateFileByBizRequest {
	s.Remark = &v
	return s
}

func (s *CreateFileByBizRequest) SetResourceOwnerAccount(v string) *CreateFileByBizRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFileByBizRequest) SetResourceOwnerId(v int64) *CreateFileByBizRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFileByBizRequest) SetSaveOssFileName(v string) *CreateFileByBizRequest {
	s.SaveOssFileName = &v
	return s
}

func (s *CreateFileByBizRequest) SetUserViewFileName(v string) *CreateFileByBizRequest {
	s.UserViewFileName = &v
	return s
}

func (s *CreateFileByBizRequest) Validate() error {
	return dara.Validate(s)
}
