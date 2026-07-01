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
	// The description of the SMS signature application. The description cannot exceed 200 characters in length.
	//
	// The description is used as a reference for signature review. A complete description helps reviewers understand your business scenario and improves review efficiency. Guidelines:
	//
	// - Provide the use case of a service that is already online.
	//
	// - Provide an SMS example from a real scenario to illustrate your business scenario.
	//
	// - Provide the values passed for variables, and describe the business scenario in detail and the reason for choosing the variable attributes.
	//
	// - Provide the website URL of the actual service, a filed domain name, or an app store download link.
	//
	// - For logon scenarios, provide a test account and password.
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
	SignFileList []*ModifySmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
	// The signature name.
	//
	// > You can modify a signature that has been approved, but you cannot change its name. The modified signature must be reviewed and approved before it can be used. The original signature cannot be used until the review is complete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The signature source. Valid values:
	//
	// - **0**: full name or abbreviation of an enterprise or public institution.
	//
	// - **1**: full name or abbreviation of a website filed with the Ministry of Industry and Information Technology (MIIT).
	//
	// - **2**: full name or abbreviation of an app.
	//
	// - **3**: full name or abbreviation of an official account or mini program.
	//
	// - **4**: full name or abbreviation of a store on an e-commerce platform.
	//
	// - **5**: full name or abbreviation of a trademark.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// The signature type. Valid values:
	//
	// - **0**: verification code.
	//
	// - **1**: general.
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

type ModifySmsSignRequestSignFileList struct {
	// 签名的纸质证明文件经base64编码后的字符串。图片不超过2 MB。
	//
	// 个别场景下，申请签名需要上传证明文件。详细说明，请参见[短信签名规范](https://help.aliyun.com/document_detail/108076.html)。
	//
	// This parameter is required.
	//
	// example:
	//
	// R0lGODlhHAAmAKIHAKqqqsvLy0hISObm5vf394uLiwAA
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	// 签名的证明文件格式，支持上传多张图片。当前支持JPG、PNG、GIF或JPEG格式的图片。
	//
	// 个别场景下，申请签名需要上传证明文件。详细说明，请参见[短信签名规范](https://help.aliyun.com/document_detail/108076.html)。
	//
	// > 如果签名用途为他用或个人认证用户的自用签名来源为企事业单位名时，还需上传证明文件和委托授权书，详情请参见[证明文件](https://help.aliyun.com/document_detail/108076.html)和[授权委托书](https://help.aliyun.com/document_detail/56741.html)。
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
