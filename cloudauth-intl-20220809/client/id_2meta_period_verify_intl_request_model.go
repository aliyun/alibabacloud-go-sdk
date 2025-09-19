// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaPeriodVerifyIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocName(v string) *Id2MetaPeriodVerifyIntlRequest
	GetDocName() *string
	SetDocNo(v string) *Id2MetaPeriodVerifyIntlRequest
	GetDocNo() *string
	SetDocType(v string) *Id2MetaPeriodVerifyIntlRequest
	GetDocType() *string
	SetMerchantBizId(v string) *Id2MetaPeriodVerifyIntlRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *Id2MetaPeriodVerifyIntlRequest
	GetMerchantUserId() *string
	SetProductCode(v string) *Id2MetaPeriodVerifyIntlRequest
	GetProductCode() *string
	SetSceneCode(v string) *Id2MetaPeriodVerifyIntlRequest
	GetSceneCode() *string
	SetValidityEndDate(v string) *Id2MetaPeriodVerifyIntlRequest
	GetValidityEndDate() *string
	SetValidityStartDate(v string) *Id2MetaPeriodVerifyIntlRequest
	GetValidityStartDate() *string
}

type Id2MetaPeriodVerifyIntlRequest struct {
	// The user\\"s name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhang San
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// The user\\"s certificate number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 411xxxxxxxxxxx0001
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// The certificate type, which is uniquely identified by an 8-digit number.
	//
	// Currently, only second-generation resident ID cards from the Chinese mainland are supported. Set the value to the static field: **00000001**.
	//
	// For more information, see [Certificate types](https://www.alibabacloud.com/help/en/ekyc/latest/im1u641gyesiqmbg?spm=a2c63.p38356.0.i13#Hu5TG).
	//
	// This parameter is required.
	//
	// example:
	//
	// ​00000001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// A unique business identifier that you can customize. Use this identifier to locate and troubleshoot issues. The identifier can be up to 32 characters in length and can contain letters and digits. Make sure that the identifier is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// A custom user ID or another identifier for a specific user, such as a mobile number or email address. Desensitize the value of this field in advance, for example, by hashing the value.
	//
	// example:
	//
	// 1234567890
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The product solution to integrate. Set the value to **eKYC_Date_MIN**.
	//
	// This parameter is required.
	//
	// example:
	//
	// eKYC_Date_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// A custom authentication scenario ID. You can use this ID to query related records in the console. The ID can be up to 10 characters in length and can contain letters, digits, and underscores (_).
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The expiration date of the ID card\\"s validity period. The format is YYYYMMDD.
	//
	// > If the ID card is valid for a long term, enter **long-term*	- for this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20301001
	ValidityEndDate *string `json:"ValidityEndDate,omitempty" xml:"ValidityEndDate,omitempty"`
	// The start date of the validity period. The format is YYYYMMDD.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20201001
	ValidityStartDate *string `json:"ValidityStartDate,omitempty" xml:"ValidityStartDate,omitempty"`
}

func (s Id2MetaPeriodVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaPeriodVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyIntlRequest) GetDocName() *string {
	return s.DocName
}

func (s *Id2MetaPeriodVerifyIntlRequest) GetDocNo() *string {
	return s.DocNo
}

func (s *Id2MetaPeriodVerifyIntlRequest) GetDocType() *string {
	return s.DocType
}

func (s *Id2MetaPeriodVerifyIntlRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *Id2MetaPeriodVerifyIntlRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *Id2MetaPeriodVerifyIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *Id2MetaPeriodVerifyIntlRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *Id2MetaPeriodVerifyIntlRequest) GetValidityEndDate() *string {
	return s.ValidityEndDate
}

func (s *Id2MetaPeriodVerifyIntlRequest) GetValidityStartDate() *string {
	return s.ValidityStartDate
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetDocName(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.DocName = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetDocNo(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.DocNo = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetDocType(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.DocType = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetMerchantBizId(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.MerchantBizId = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetMerchantUserId(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.MerchantUserId = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetProductCode(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetSceneCode(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.SceneCode = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetValidityEndDate(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.ValidityEndDate = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) SetValidityStartDate(v string) *Id2MetaPeriodVerifyIntlRequest {
	s.ValidityStartDate = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
