// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalSmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CreateDigitalSmsTemplateRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateDigitalSmsTemplateRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateDigitalSmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDigitalSmsTemplateRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *CreateDigitalSmsTemplateRequest
	GetSignName() *string
	SetTemplateContents(v []*CreateDigitalSmsTemplateRequestTemplateContents) *CreateDigitalSmsTemplateRequest
	GetTemplateContents() []*CreateDigitalSmsTemplateRequestTemplateContents
	SetTemplateName(v string) *CreateDigitalSmsTemplateRequest
	GetTemplateName() *string
}

type CreateDigitalSmsTemplateRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// This parameter is required.
	TemplateContents []*CreateDigitalSmsTemplateRequestTemplateContents `json:"TemplateContents,omitempty" xml:"TemplateContents,omitempty" type:"Repeated"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateDigitalSmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateDigitalSmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDigitalSmsTemplateRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateDigitalSmsTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDigitalSmsTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDigitalSmsTemplateRequest) GetSignName() *string {
	return s.SignName
}

func (s *CreateDigitalSmsTemplateRequest) GetTemplateContents() []*CreateDigitalSmsTemplateRequestTemplateContents {
	return s.TemplateContents
}

func (s *CreateDigitalSmsTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateDigitalSmsTemplateRequest) SetOwnerId(v int64) *CreateDigitalSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequest) SetRemark(v string) *CreateDigitalSmsTemplateRequest {
	s.Remark = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequest) SetResourceOwnerAccount(v string) *CreateDigitalSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequest) SetResourceOwnerId(v int64) *CreateDigitalSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequest) SetSignName(v string) *CreateDigitalSmsTemplateRequest {
	s.SignName = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequest) SetTemplateContents(v []*CreateDigitalSmsTemplateRequestTemplateContents) *CreateDigitalSmsTemplateRequest {
	s.TemplateContents = v
	return s
}

func (s *CreateDigitalSmsTemplateRequest) SetTemplateName(v string) *CreateDigitalSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequest) Validate() error {
	if s.TemplateContents != nil {
		for _, item := range s.TemplateContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDigitalSmsTemplateRequestTemplateContents struct {
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	// example:
	//
	// file-1
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 1840901
	FileSize *int32 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// txt
	FileSuffix *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
}

func (s CreateDigitalSmsTemplateRequestTemplateContents) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSmsTemplateRequestTemplateContents) GoString() string {
	return s.String()
}

func (s *CreateDigitalSmsTemplateRequestTemplateContents) GetFileContents() *string {
	return s.FileContents
}

func (s *CreateDigitalSmsTemplateRequestTemplateContents) GetFileName() *string {
	return s.FileName
}

func (s *CreateDigitalSmsTemplateRequestTemplateContents) GetFileSize() *int32 {
	return s.FileSize
}

func (s *CreateDigitalSmsTemplateRequestTemplateContents) GetFileSuffix() *string {
	return s.FileSuffix
}

func (s *CreateDigitalSmsTemplateRequestTemplateContents) SetFileContents(v string) *CreateDigitalSmsTemplateRequestTemplateContents {
	s.FileContents = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequestTemplateContents) SetFileName(v string) *CreateDigitalSmsTemplateRequestTemplateContents {
	s.FileName = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequestTemplateContents) SetFileSize(v int32) *CreateDigitalSmsTemplateRequestTemplateContents {
	s.FileSize = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequestTemplateContents) SetFileSuffix(v string) *CreateDigitalSmsTemplateRequestTemplateContents {
	s.FileSuffix = &v
	return s
}

func (s *CreateDigitalSmsTemplateRequestTemplateContents) Validate() error {
	return dara.Validate(s)
}
