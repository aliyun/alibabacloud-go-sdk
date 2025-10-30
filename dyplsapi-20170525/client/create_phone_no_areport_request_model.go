// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhoneNoAReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetANoWhiteGroupId(v int64) *CreatePhoneNoAReportRequest
	GetANoWhiteGroupId() *int64
	SetCustName(v string) *CreatePhoneNoAReportRequest
	GetCustName() *string
	SetDocumentNumber(v string) *CreatePhoneNoAReportRequest
	GetDocumentNumber() *string
	SetDocumentType(v int64) *CreatePhoneNoAReportRequest
	GetDocumentType() *int64
	SetIdCardAlivePhoto(v string) *CreatePhoneNoAReportRequest
	GetIdCardAlivePhoto() *string
	SetIdCardBackPhoto(v string) *CreatePhoneNoAReportRequest
	GetIdCardBackPhoto() *string
	SetIdCardFrontPhoto(v string) *CreatePhoneNoAReportRequest
	GetIdCardFrontPhoto() *string
	SetOwnerId(v int64) *CreatePhoneNoAReportRequest
	GetOwnerId() *int64
	SetPhoneNoA(v string) *CreatePhoneNoAReportRequest
	GetPhoneNoA() *string
	SetRemark(v string) *CreatePhoneNoAReportRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreatePhoneNoAReportRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePhoneNoAReportRequest
	GetResourceOwnerId() *int64
}

type CreatePhoneNoAReportRequest struct {
	// 所属a号码组id
	//
	// This parameter is required.
	//
	// example:
	//
	// 19
	ANoWhiteGroupId *int64 `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	// 姓名
	//
	// This parameter is required.
	//
	// example:
	//
	// 赵**
	CustName *string `json:"CustName,omitempty" xml:"CustName,omitempty"`
	// 证件号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 370*********
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// 证件类型 填写1表示身份证
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DocumentType *int64 `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// 半身照oss路径地址
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/test1719383196033.jpg示例值
	IdCardAlivePhoto *string `json:"IdCardAlivePhoto,omitempty" xml:"IdCardAlivePhoto,omitempty"`
	// 身份证反面照片oss路径地址
	//
	// example:
	//
	// 123456/test1719383196032.jpg
	IdCardBackPhoto *string `json:"IdCardBackPhoto,omitempty" xml:"IdCardBackPhoto,omitempty"`
	// 身份证正面照片oss路径地址
	//
	// example:
	//
	// 123456/test1719383196031.jpg
	IdCardFrontPhoto *string `json:"IdCardFrontPhoto,omitempty" xml:"IdCardFrontPhoto,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 手机号
	//
	// This parameter is required.
	//
	// example:
	//
	// 130*****1234
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// 备注（客户自己的业务备注，可编辑）
	//
	// example:
	//
	// ***报备
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreatePhoneNoAReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePhoneNoAReportRequest) GoString() string {
	return s.String()
}

func (s *CreatePhoneNoAReportRequest) GetANoWhiteGroupId() *int64 {
	return s.ANoWhiteGroupId
}

func (s *CreatePhoneNoAReportRequest) GetCustName() *string {
	return s.CustName
}

func (s *CreatePhoneNoAReportRequest) GetDocumentNumber() *string {
	return s.DocumentNumber
}

func (s *CreatePhoneNoAReportRequest) GetDocumentType() *int64 {
	return s.DocumentType
}

func (s *CreatePhoneNoAReportRequest) GetIdCardAlivePhoto() *string {
	return s.IdCardAlivePhoto
}

func (s *CreatePhoneNoAReportRequest) GetIdCardBackPhoto() *string {
	return s.IdCardBackPhoto
}

func (s *CreatePhoneNoAReportRequest) GetIdCardFrontPhoto() *string {
	return s.IdCardFrontPhoto
}

func (s *CreatePhoneNoAReportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePhoneNoAReportRequest) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *CreatePhoneNoAReportRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreatePhoneNoAReportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePhoneNoAReportRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePhoneNoAReportRequest) SetANoWhiteGroupId(v int64) *CreatePhoneNoAReportRequest {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetCustName(v string) *CreatePhoneNoAReportRequest {
	s.CustName = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetDocumentNumber(v string) *CreatePhoneNoAReportRequest {
	s.DocumentNumber = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetDocumentType(v int64) *CreatePhoneNoAReportRequest {
	s.DocumentType = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetIdCardAlivePhoto(v string) *CreatePhoneNoAReportRequest {
	s.IdCardAlivePhoto = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetIdCardBackPhoto(v string) *CreatePhoneNoAReportRequest {
	s.IdCardBackPhoto = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetIdCardFrontPhoto(v string) *CreatePhoneNoAReportRequest {
	s.IdCardFrontPhoto = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetOwnerId(v int64) *CreatePhoneNoAReportRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetPhoneNoA(v string) *CreatePhoneNoAReportRequest {
	s.PhoneNoA = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetRemark(v string) *CreatePhoneNoAReportRequest {
	s.Remark = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetResourceOwnerAccount(v string) *CreatePhoneNoAReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) SetResourceOwnerId(v int64) *CreatePhoneNoAReportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePhoneNoAReportRequest) Validate() error {
	return dara.Validate(s)
}
