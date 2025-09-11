// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ModifySmsSignRequest
	GetOwnerId() *int64
	SetRemark(v string) *ModifySmsSignRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *ModifySmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySmsSignRequest
	GetResourceOwnerId() *int64
	SetSignFileList(v []*ModifySmsSignRequestSignFileList) *ModifySmsSignRequest
	GetSignFileList() []*ModifySmsSignRequestSignFileList
	SetSignName(v string) *ModifySmsSignRequest
	GetSignName() *string
	SetSignSource(v int32) *ModifySmsSignRequest
	GetSignSource() *int32
	SetSignType(v int32) *ModifySmsSignRequest
	GetSignType() *int32
}

type ModifySmsSignRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The scenario description of your released services. Provide the information of your services, such as a website URL, a domain name with an ICP filing, an app download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
	//
	// > The description can be up to 200 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// This is the abbreviation of our company.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of signature files.
	//
	// This parameter is required.
	SignFileList []*ModifySmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
	// The signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The source of the signature. Valid values:
	//
	// 	- **0**: full name or abbreviation of an enterprise or institution.
	//
	// 	- **1**: full name or abbreviation of a website with Ministry of Industry and Information Technology (MIIT) filing.
	//
	// 	- **2**: full name or abbreviation of an app.
	//
	// 	- **3**: full name or abbreviation of a WeChat official account or applet.
	//
	// 	- **4**: full name or abbreviation of an e-commerce store.
	//
	// 	- **5**: full name or abbreviation of a trademark.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// The type of the signature. Valid values:
	//
	// 	- **0**: verification-code signature
	//
	// 	- **1**: general-purpose signature
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
}

func (s ModifySmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySmsSignRequest) GoString() string {
	return s.String()
}

func (s *ModifySmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySmsSignRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifySmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySmsSignRequest) GetSignFileList() []*ModifySmsSignRequestSignFileList {
	return s.SignFileList
}

func (s *ModifySmsSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *ModifySmsSignRequest) GetSignSource() *int32 {
	return s.SignSource
}

func (s *ModifySmsSignRequest) GetSignType() *int32 {
	return s.SignType
}

func (s *ModifySmsSignRequest) SetOwnerId(v int64) *ModifySmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySmsSignRequest) SetRemark(v string) *ModifySmsSignRequest {
	s.Remark = &v
	return s
}

func (s *ModifySmsSignRequest) SetResourceOwnerAccount(v string) *ModifySmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySmsSignRequest) SetResourceOwnerId(v int64) *ModifySmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySmsSignRequest) SetSignFileList(v []*ModifySmsSignRequestSignFileList) *ModifySmsSignRequest {
	s.SignFileList = v
	return s
}

func (s *ModifySmsSignRequest) SetSignName(v string) *ModifySmsSignRequest {
	s.SignName = &v
	return s
}

func (s *ModifySmsSignRequest) SetSignSource(v int32) *ModifySmsSignRequest {
	s.SignSource = &v
	return s
}

func (s *ModifySmsSignRequest) SetSignType(v int32) *ModifySmsSignRequest {
	s.SignType = &v
	return s
}

func (s *ModifySmsSignRequest) Validate() error {
	return dara.Validate(s)
}

type ModifySmsSignRequestSignFileList struct {
	// The base64-encoded string of the signed files. The size of the image cannot exceed 2 MB.
	//
	// In some scenarios, documents are required to prove your identity. For more information, see [Signature specifications](https://help.aliyun.com/document_detail/108076.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// R0lGODlhHAAmAKIHAKqqqsvLy0hISObm5vf394uLiwAA
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	// The format of the documents. You can upload multiple images. JPG, PNG, GIF, and JPEG are supported.
	//
	// In some scenarios, documents are required to prove your identity. For more information, see [Signature specifications](https://help.aliyun.com/document_detail/108076.html).
	//
	// > If the signature is used for other purposes or the signature source is an enterprise or public institution, you must upload some documents and an authorization letter. For more information, see [Documents](https://help.aliyun.com/document_detail/108076.html) and [Letter of authorization](https://help.aliyun.com/document_detail/56741.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// jpg
	FileSuffix *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
}

func (s ModifySmsSignRequestSignFileList) String() string {
	return dara.Prettify(s)
}

func (s ModifySmsSignRequestSignFileList) GoString() string {
	return s.String()
}

func (s *ModifySmsSignRequestSignFileList) GetFileContents() *string {
	return s.FileContents
}

func (s *ModifySmsSignRequestSignFileList) GetFileSuffix() *string {
	return s.FileSuffix
}

func (s *ModifySmsSignRequestSignFileList) SetFileContents(v string) *ModifySmsSignRequestSignFileList {
	s.FileContents = &v
	return s
}

func (s *ModifySmsSignRequestSignFileList) SetFileSuffix(v string) *ModifySmsSignRequestSignFileList {
	s.FileSuffix = &v
	return s
}

func (s *ModifySmsSignRequestSignFileList) Validate() error {
	return dara.Validate(s)
}
