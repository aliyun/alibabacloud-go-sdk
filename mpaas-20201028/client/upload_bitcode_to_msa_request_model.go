// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadBitcodeToMsaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UploadBitcodeToMsaRequest
	GetAppId() *string
	SetBitcode(v string) *UploadBitcodeToMsaRequest
	GetBitcode() *string
	SetCodeVersion(v string) *UploadBitcodeToMsaRequest
	GetCodeVersion() *string
	SetLicense(v string) *UploadBitcodeToMsaRequest
	GetLicense() *string
	SetTenantId(v string) *UploadBitcodeToMsaRequest
	GetTenantId() *string
	SetType(v string) *UploadBitcodeToMsaRequest
	GetType() *string
	SetWorkspaceId(v string) *UploadBitcodeToMsaRequest
	GetWorkspaceId() *string
}

type UploadBitcodeToMsaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 3sAXCwAAAAAUAAAACHoAAP
	Bitcode *string `json:"Bitcode,omitempty" xml:"Bitcode,omitempty"`
	// example:
	//
	// xcode14
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// example:
	//
	// {}
	License *string `json:"License,omitempty" xml:"License,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UploadBitcodeToMsaRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadBitcodeToMsaRequest) GoString() string {
	return s.String()
}

func (s *UploadBitcodeToMsaRequest) GetAppId() *string {
	return s.AppId
}

func (s *UploadBitcodeToMsaRequest) GetBitcode() *string {
	return s.Bitcode
}

func (s *UploadBitcodeToMsaRequest) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *UploadBitcodeToMsaRequest) GetLicense() *string {
	return s.License
}

func (s *UploadBitcodeToMsaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadBitcodeToMsaRequest) GetType() *string {
	return s.Type
}

func (s *UploadBitcodeToMsaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UploadBitcodeToMsaRequest) SetAppId(v string) *UploadBitcodeToMsaRequest {
	s.AppId = &v
	return s
}

func (s *UploadBitcodeToMsaRequest) SetBitcode(v string) *UploadBitcodeToMsaRequest {
	s.Bitcode = &v
	return s
}

func (s *UploadBitcodeToMsaRequest) SetCodeVersion(v string) *UploadBitcodeToMsaRequest {
	s.CodeVersion = &v
	return s
}

func (s *UploadBitcodeToMsaRequest) SetLicense(v string) *UploadBitcodeToMsaRequest {
	s.License = &v
	return s
}

func (s *UploadBitcodeToMsaRequest) SetTenantId(v string) *UploadBitcodeToMsaRequest {
	s.TenantId = &v
	return s
}

func (s *UploadBitcodeToMsaRequest) SetType(v string) *UploadBitcodeToMsaRequest {
	s.Type = &v
	return s
}

func (s *UploadBitcodeToMsaRequest) SetWorkspaceId(v string) *UploadBitcodeToMsaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UploadBitcodeToMsaRequest) Validate() error {
	return dara.Validate(s)
}
