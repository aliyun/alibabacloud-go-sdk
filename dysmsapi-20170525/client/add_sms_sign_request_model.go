// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *AddSmsSignRequest
	GetOwnerId() *int64
	SetRemark(v string) *AddSmsSignRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *AddSmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddSmsSignRequest
	GetResourceOwnerId() *int64
	SetSignFileList(v []*AddSmsSignRequestSignFileList) *AddSmsSignRequest
	GetSignFileList() []*AddSmsSignRequestSignFileList
	SetSignName(v string) *AddSmsSignRequest
	GetSignName() *string
	SetSignSource(v int32) *AddSmsSignRequest
	GetSignSource() *int32
	SetSignType(v int32) *AddSmsSignRequest
	GetSignType() *int32
}

type AddSmsSignRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the signature application. The description cannot exceed 200 characters in length. The description is one of the reference information for signature review. We recommend that you describe the use scenarios of your services in detail, and provide information that can verify the services, such as a website URL, a domain name with an ICP filing, an app download URL, an official account name, or a mini program name. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// This is the abbreviation of our company.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature files.
	//
	// This parameter is required.
	SignFileList []*AddSmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
	// The name of the signature.
	//
	// >
	//
	// 	- The signature name is not case-sensitive. For example, [Alibaba Cloud Communication] and [alibaba cloud communication] are considered as the same name.
	//
	// 	- If your verification code signature and general-purpose signature have the same name, the system uses the general-purpose signature to send messages by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The source of the signature. Valid values:
	//
	// 	- **0**: the full name or abbreviation of an enterprise or institution
	//
	// 	- **1**: the full name or abbreviation of a website that has obtained an ICP filing from the Ministry of Industry and Information Technology (MIIT) of China
	//
	// 	- **2**: the full name or abbreviation of an app
	//
	// 	- **3**: the full name or abbreviation of an official account or mini-program
	//
	// 	- **4**: the full name or abbreviation of an e-commerce store
	//
	// 	- **5**: the full name or abbreviation of a trademark
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// The type of the signature. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: general-purpose
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
}

func (s AddSmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSmsSignRequest) GoString() string {
	return s.String()
}

func (s *AddSmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddSmsSignRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddSmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddSmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddSmsSignRequest) GetSignFileList() []*AddSmsSignRequestSignFileList {
	return s.SignFileList
}

func (s *AddSmsSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *AddSmsSignRequest) GetSignSource() *int32 {
	return s.SignSource
}

func (s *AddSmsSignRequest) GetSignType() *int32 {
	return s.SignType
}

func (s *AddSmsSignRequest) SetOwnerId(v int64) *AddSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *AddSmsSignRequest) SetRemark(v string) *AddSmsSignRequest {
	s.Remark = &v
	return s
}

func (s *AddSmsSignRequest) SetResourceOwnerAccount(v string) *AddSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSmsSignRequest) SetResourceOwnerId(v int64) *AddSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddSmsSignRequest) SetSignFileList(v []*AddSmsSignRequestSignFileList) *AddSmsSignRequest {
	s.SignFileList = v
	return s
}

func (s *AddSmsSignRequest) SetSignName(v string) *AddSmsSignRequest {
	s.SignName = &v
	return s
}

func (s *AddSmsSignRequest) SetSignSource(v int32) *AddSmsSignRequest {
	s.SignSource = &v
	return s
}

func (s *AddSmsSignRequest) SetSignType(v int32) *AddSmsSignRequest {
	s.SignType = &v
	return s
}

func (s *AddSmsSignRequest) Validate() error {
	if s.SignFileList != nil {
		for _, item := range s.SignFileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddSmsSignRequestSignFileList struct {
	// The Base64-encoded string of the qualification document. An image cannot exceed 2 MB in size. In some scenarios, you must upload supporting documents to apply for signatures. For more information, see [SMS signature specifications](https://help.aliyun.com/document_detail/108076.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// R0lGODlhHAAmAKIHAKqqqsvLy0hISObm5vf394uL****
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	// The format of the qualification document. You can upload multiple images. Images in JPG, PNG, GIF, or JPEG format are supported.
	//
	// In some scenarios, you must upload supporting documents to apply for signatures. For more information, see [SMS signature specifications](https://help.aliyun.com/document_detail/108076.html).
	//
	// > If you apply for a signature for other users or if the signature source is the name of an enterprise or public institution, you must upload a certificate and a letter of authorization. For more information, see [Certificate](https://help.aliyun.com/document_detail/108076.html) and [Letter of authorization](https://help.aliyun.com/document_detail/56741.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// jpg
	FileSuffix *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
}

func (s AddSmsSignRequestSignFileList) String() string {
	return dara.Prettify(s)
}

func (s AddSmsSignRequestSignFileList) GoString() string {
	return s.String()
}

func (s *AddSmsSignRequestSignFileList) GetFileContents() *string {
	return s.FileContents
}

func (s *AddSmsSignRequestSignFileList) GetFileSuffix() *string {
	return s.FileSuffix
}

func (s *AddSmsSignRequestSignFileList) SetFileContents(v string) *AddSmsSignRequestSignFileList {
	s.FileContents = &v
	return s
}

func (s *AddSmsSignRequestSignFileList) SetFileSuffix(v string) *AddSmsSignRequestSignFileList {
	s.FileSuffix = &v
	return s
}

func (s *AddSmsSignRequestSignFileList) Validate() error {
	return dara.Validate(s)
}
