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
	// This parameter is required.
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 411xxxxxxxxxxx0001
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ​00000001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 1234567890
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eKYC_Date_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20301001
	ValidityEndDate *string `json:"ValidityEndDate,omitempty" xml:"ValidityEndDate,omitempty"`
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
