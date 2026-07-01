// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsTrademarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CreateSmsTrademarkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateSmsTrademarkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsTrademarkRequest
	GetResourceOwnerId() *int64
	SetTrademarkApplicantName(v string) *CreateSmsTrademarkRequest
	GetTrademarkApplicantName() *string
	SetTrademarkEffExpDate(v string) *CreateSmsTrademarkRequest
	GetTrademarkEffExpDate() *string
	SetTrademarkName(v string) *CreateSmsTrademarkRequest
	GetTrademarkName() *string
	SetTrademarkPic(v string) *CreateSmsTrademarkRequest
	GetTrademarkPic() *string
	SetTrademarkRegistrationNumber(v string) *CreateSmsTrademarkRequest
	GetTrademarkRegistrationNumber() *string
}

type CreateSmsTrademarkRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the applicant. The value can be up to 50 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里巴巴（中国）有限公司
	TrademarkApplicantName *string `json:"TrademarkApplicantName,omitempty" xml:"TrademarkApplicantName,omitempty"`
	// The effective and expiration dates of the exclusive right.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2010-12-14~2030-12-13
	TrademarkEffExpDate *string `json:"TrademarkEffExpDate,omitempty" xml:"TrademarkEffExpDate,omitempty"`
	// The trademark name. The value can be up to 15 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云
	TrademarkName *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	// The fileKey of the trademark details screenshot.
	//
	// 1. How to query a trademark:
	//
	// - Log on to the China Trademark Network, click **Trademark Online Query**, and take a screenshot of the trademark details.
	//
	// - Accept the terms of use and enter the **Application/Registration Number*	- to query.
	//
	// - Click the **Application/Registration Number*	- to view the details.
	//
	// 2. Information about the trademark file uploaded to OSS. File upload requirements:
	//
	// - The name of the file to be uploaded cannot contain Chinese characters or special characters.
	//
	// - Only images in JPG, PNG, GIF, and JPEG formats are supported, and the image size cannot exceed 5 MB.
	//
	// - The screenshot must contain the complete URL.
	//
	// - The trademark image must be clear and identical to the **signature name**.
	//
	// - The **applicant name*	- must be identical to the name of the enterprise or institution associated with the signature.
	//
	// - The trademark status must be registered trademark.
	//
	// 3. To obtain the fileKey, see [Upload files to OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/test1719383196031.jpg
	TrademarkPic *string `json:"TrademarkPic,omitempty" xml:"TrademarkPic,omitempty"`
	// The trademark registration number. The value can be up to 15 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 766905
	TrademarkRegistrationNumber *string `json:"TrademarkRegistrationNumber,omitempty" xml:"TrademarkRegistrationNumber,omitempty"`
}

func (s CreateSmsTrademarkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsTrademarkRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsTrademarkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsTrademarkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsTrademarkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsTrademarkRequest) GetTrademarkApplicantName() *string {
	return s.TrademarkApplicantName
}

func (s *CreateSmsTrademarkRequest) GetTrademarkEffExpDate() *string {
	return s.TrademarkEffExpDate
}

func (s *CreateSmsTrademarkRequest) GetTrademarkName() *string {
	return s.TrademarkName
}

func (s *CreateSmsTrademarkRequest) GetTrademarkPic() *string {
	return s.TrademarkPic
}

func (s *CreateSmsTrademarkRequest) GetTrademarkRegistrationNumber() *string {
	return s.TrademarkRegistrationNumber
}

func (s *CreateSmsTrademarkRequest) SetOwnerId(v int64) *CreateSmsTrademarkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsTrademarkRequest) SetResourceOwnerAccount(v string) *CreateSmsTrademarkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsTrademarkRequest) SetResourceOwnerId(v int64) *CreateSmsTrademarkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsTrademarkRequest) SetTrademarkApplicantName(v string) *CreateSmsTrademarkRequest {
	s.TrademarkApplicantName = &v
	return s
}

func (s *CreateSmsTrademarkRequest) SetTrademarkEffExpDate(v string) *CreateSmsTrademarkRequest {
	s.TrademarkEffExpDate = &v
	return s
}

func (s *CreateSmsTrademarkRequest) SetTrademarkName(v string) *CreateSmsTrademarkRequest {
	s.TrademarkName = &v
	return s
}

func (s *CreateSmsTrademarkRequest) SetTrademarkPic(v string) *CreateSmsTrademarkRequest {
	s.TrademarkPic = &v
	return s
}

func (s *CreateSmsTrademarkRequest) SetTrademarkRegistrationNumber(v string) *CreateSmsTrademarkRequest {
	s.TrademarkRegistrationNumber = &v
	return s
}

func (s *CreateSmsTrademarkRequest) Validate() error {
	return dara.Validate(s)
}
