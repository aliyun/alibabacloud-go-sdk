// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFixedNoAReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetANoWhiteGroupId(v int64) *CreateFixedNoAReportRequest
	GetANoWhiteGroupId() *int64
	SetCustName(v string) *CreateFixedNoAReportRequest
	GetCustName() *string
	SetCustPhoneNo(v string) *CreateFixedNoAReportRequest
	GetCustPhoneNo() *string
	SetCustType(v int64) *CreateFixedNoAReportRequest
	GetCustType() *int64
	SetDocumentNumber(v string) *CreateFixedNoAReportRequest
	GetDocumentNumber() *string
	SetDocumentType(v int64) *CreateFixedNoAReportRequest
	GetDocumentType() *int64
	SetFixedLineWorkId(v string) *CreateFixedNoAReportRequest
	GetFixedLineWorkId() *string
	SetFixedNoA(v string) *CreateFixedNoAReportRequest
	GetFixedNoA() *string
	SetIdCardAlivePhoto(v string) *CreateFixedNoAReportRequest
	GetIdCardAlivePhoto() *string
	SetIdCardBackPhoto(v string) *CreateFixedNoAReportRequest
	GetIdCardBackPhoto() *string
	SetIdCardFrontPhoto(v string) *CreateFixedNoAReportRequest
	GetIdCardFrontPhoto() *string
	SetOwnerId(v int64) *CreateFixedNoAReportRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateFixedNoAReportRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateFixedNoAReportRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFixedNoAReportRequest
	GetResourceOwnerId() *int64
}

type CreateFixedNoAReportRequest struct {
	// ID of the group to which Number A belongs. On the [Qualification Management](https://dyplsnext.console.aliyun.com/qualification) Page, go to the A-number whitelist section of the relevant qualification and click View Details.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ANoWhiteGroupId *int64 `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	// Name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 赵**
	CustName *string `json:"CustName,omitempty" xml:"CustName,omitempty"`
	// Phone number of the legal representative or agent.
	//
	// >If this qualification includes a Mobile X number, this field is Required; otherwise, the Review will Fail.
	//
	// example:
	//
	// 130*****8888
	CustPhoneNo *string `json:"CustPhoneNo,omitempty" xml:"CustPhoneNo,omitempty"`
	// Fixed-line customer type. Valid values:
	//
	// - 1: Legal representative.
	//
	// - 2: Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustType *int64 `json:"CustType,omitempty" xml:"CustType,omitempty"`
	// Certificate number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 370*********
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// Certificate type. Enter 1 for ID card.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DocumentType *int64 `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// Ticket number from the offline service hall when applying for a fixed-line number.
	//
	// example:
	//
	// 45***12
	FixedLineWorkId *string `json:"FixedLineWorkId,omitempty" xml:"FixedLineWorkId,omitempty"`
	// Number A (fixed-line number).
	//
	// This parameter is required.
	//
	// example:
	//
	// 89*****1234
	FixedNoA *string `json:"FixedNoA,omitempty" xml:"FixedNoA,omitempty"`
	// OSS path of the front-facing portrait photo.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/test1719383196033.jpg
	IdCardAlivePhoto *string `json:"IdCardAlivePhoto,omitempty" xml:"IdCardAlivePhoto,omitempty"`
	// OSS path of the photo of the back (portrait side) of the ID card.
	//
	// example:
	//
	// 123456/test1719383196032.jpg
	IdCardBackPhoto *string `json:"IdCardBackPhoto,omitempty" xml:"IdCardBackPhoto,omitempty"`
	// OSS path of the photo of the front (national emblem side) of the ID card.
	//
	// example:
	//
	// 123456/test1719383196031.jpg
	IdCardFrontPhoto *string `json:"IdCardFrontPhoto,omitempty" xml:"IdCardFrontPhoto,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Remarks (user-defined business notes).
	//
	// example:
	//
	// ***报备
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateFixedNoAReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFixedNoAReportRequest) GoString() string {
	return s.String()
}

func (s *CreateFixedNoAReportRequest) GetANoWhiteGroupId() *int64 {
	return s.ANoWhiteGroupId
}

func (s *CreateFixedNoAReportRequest) GetCustName() *string {
	return s.CustName
}

func (s *CreateFixedNoAReportRequest) GetCustPhoneNo() *string {
	return s.CustPhoneNo
}

func (s *CreateFixedNoAReportRequest) GetCustType() *int64 {
	return s.CustType
}

func (s *CreateFixedNoAReportRequest) GetDocumentNumber() *string {
	return s.DocumentNumber
}

func (s *CreateFixedNoAReportRequest) GetDocumentType() *int64 {
	return s.DocumentType
}

func (s *CreateFixedNoAReportRequest) GetFixedLineWorkId() *string {
	return s.FixedLineWorkId
}

func (s *CreateFixedNoAReportRequest) GetFixedNoA() *string {
	return s.FixedNoA
}

func (s *CreateFixedNoAReportRequest) GetIdCardAlivePhoto() *string {
	return s.IdCardAlivePhoto
}

func (s *CreateFixedNoAReportRequest) GetIdCardBackPhoto() *string {
	return s.IdCardBackPhoto
}

func (s *CreateFixedNoAReportRequest) GetIdCardFrontPhoto() *string {
	return s.IdCardFrontPhoto
}

func (s *CreateFixedNoAReportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFixedNoAReportRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateFixedNoAReportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFixedNoAReportRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFixedNoAReportRequest) SetANoWhiteGroupId(v int64) *CreateFixedNoAReportRequest {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetCustName(v string) *CreateFixedNoAReportRequest {
	s.CustName = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetCustPhoneNo(v string) *CreateFixedNoAReportRequest {
	s.CustPhoneNo = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetCustType(v int64) *CreateFixedNoAReportRequest {
	s.CustType = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetDocumentNumber(v string) *CreateFixedNoAReportRequest {
	s.DocumentNumber = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetDocumentType(v int64) *CreateFixedNoAReportRequest {
	s.DocumentType = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetFixedLineWorkId(v string) *CreateFixedNoAReportRequest {
	s.FixedLineWorkId = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetFixedNoA(v string) *CreateFixedNoAReportRequest {
	s.FixedNoA = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetIdCardAlivePhoto(v string) *CreateFixedNoAReportRequest {
	s.IdCardAlivePhoto = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetIdCardBackPhoto(v string) *CreateFixedNoAReportRequest {
	s.IdCardBackPhoto = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetIdCardFrontPhoto(v string) *CreateFixedNoAReportRequest {
	s.IdCardFrontPhoto = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetOwnerId(v int64) *CreateFixedNoAReportRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetRemark(v string) *CreateFixedNoAReportRequest {
	s.Remark = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetResourceOwnerAccount(v string) *CreateFixedNoAReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFixedNoAReportRequest) SetResourceOwnerId(v int64) *CreateFixedNoAReportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFixedNoAReportRequest) Validate() error {
	return dara.Validate(s)
}
