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
	// The description of the SMS signature scenario. The description cannot exceed 200 characters in length.
	//
	// This is reference information for signature review. Providing a complete application description helps reviewers understand your business scenario and improves review efficiency. Guidelines for filling in:
	//
	// - You can provide the use cases of a business that has been launched.
	//
	// - You can provide real-world SMS message examples to reflect your business scenarios.
	//
	// - You can provide the parameter values passed to variables and describe the business use cases and the reasons for choosing these variable attributes in detail.
	//
	// - You can provide the website links, registered domain names, or app store download links of your actual business.
	//
	// - For login scenarios, you can provide a test account and password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 当前的短信签名应用于双11大促推广营销
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of signature files.
	//
	// This parameter is required.
	SignFileList []*AddSmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
	// The signature name. The signature name must comply with the [Signature specifications](~~108076#section-0p8-qn8-mmy~~).
	//
	// > - Signature names are case-insensitive. For example, [Aliyun Communication] and [aliyun communication] are considered the same name.
	//
	// > - When your verification code signature and general-purpose signature have the same name, the system uses the general-purpose signature to send SMS messages by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The source of the signature. Valid values:
	//
	// -  **0**: Full name or abbreviation of an enterprise or public institution.
	//
	// -  **1**: Full name or abbreviation of a website registered with the Ministry of Industry and Information Technology (MIIT).
	//
	// -  **2**: Full name or abbreviation of an app.
	//
	// -  **3**: Full name or abbreviation of an official account or mini program.
	//
	// -  **4**: Full name or abbreviation of an e-commerce platform store name.
	//
	// -  **5**: Full name or abbreviation of a trademark name.
	//
	// For detailed descriptions of signature sources, see [Signature source](https://help.aliyun.com/en/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.0.4f9710fder2gR7#section-xup-k46-yi4).
	//
	// >This API does not support applying for signatures whose signature source is **Test or learning*	- or **Online trial**. If you need to apply for signatures with these two sources, go to the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/domestic/text/sign/add/qualification) to submit your application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// The type of the signature.
	//
	// - **0**: Verification code
	//
	// - **1**: General-purpose
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
	// The Base64-encoded string of the qualification certificate file for the signature. The image size cannot exceed 2 MB. In some scenarios, you need to upload a certificate file when you apply for a signature.
	//
	// For detailed specifications, see [File upload specifications](https://help.aliyun.com/document_detail/463316.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// R0lGODlhHAAmAKIHAKqqqsvLy0hISObm5vf394uL****
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	// The format of the signature certificate file. Multiple images can be uploaded. Currently, JPG, PNG, GIF, and JPEG formats are supported. In some scenarios, you need to upload a certificate file when you apply for a signature.
	//
	// > If the signature is for third-party use or if you are an individual-certified user whose self-use signature source is an enterprise or public institution name, you also need to upload a certificate file and a power of attorney. For more information, see [Certificate file](https://help.aliyun.com/document_detail/108076.html) and [Power of attorney](https://help.aliyun.com/document_detail/56741.html).
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
