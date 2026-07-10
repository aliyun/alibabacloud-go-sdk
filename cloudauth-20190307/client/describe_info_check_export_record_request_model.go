// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInfoCheckExportRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeInfoCheckExportRecordRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *DescribeInfoCheckExportRecordRequest
	GetEndDate() *string
	SetPageSize(v int32) *DescribeInfoCheckExportRecordRequest
	GetPageSize() *int32
	SetProductType(v string) *DescribeInfoCheckExportRecordRequest
	GetProductType() *string
	SetStartDate(v string) *DescribeInfoCheckExportRecordRequest
	GetStartDate() *string
}

type DescribeInfoCheckExportRecordRequest struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query.
	//
	// example:
	//
	// 2025-10-11 21:24:48
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. Valid values:
	//
	// - **ID_CARD_2_META**: ID card two-element verification.
	//
	// - **ID_PERIOD**: ID card validity period verification.
	//
	// - **MOBILE_ONLINE_LENGTH**: mobile number online duration.
	//
	// - **MOBILE_ONLINE_STATUS**: mobile number online status.
	//
	// - **MOBILE_3_META_SIMPLE**: mobile number three-element verification (simple edition).
	//
	// - **MOBILE_3_META**: mobile number three-element verification (detailed edition).
	//
	// - **MOBILE_2_META**: mobile number two-element verification.
	//
	// - **BANK_CARD_N_META**: bank card verification (detailed edition).
	//
	// - **MOBILE_DETECT**: phone number detection.
	//
	// - **VEHICLE_N_META**: vehicle element verification (enhanced edition).
	//
	// - **VEHICLE_PENTA_INFO**: vehicle five-element information recognition.
	//
	// - **VEHICLE_LICENSE_INFO**: vehicle information recognition.
	//
	// - **VEHICLE_INSURE_DATE**: vehicle insurance date query.
	//
	// - **VEHICLE_CHECK**: vehicle element verification.
	//
	// example:
	//
	// ID_CARD_2_META
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2025-10-11 21:24:48
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeInfoCheckExportRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInfoCheckExportRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeInfoCheckExportRecordRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInfoCheckExportRecordRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeInfoCheckExportRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInfoCheckExportRecordRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInfoCheckExportRecordRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeInfoCheckExportRecordRequest) SetCurrentPage(v int32) *DescribeInfoCheckExportRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) SetEndDate(v string) *DescribeInfoCheckExportRecordRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) SetPageSize(v int32) *DescribeInfoCheckExportRecordRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) SetProductType(v string) *DescribeInfoCheckExportRecordRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) SetStartDate(v string) *DescribeInfoCheckExportRecordRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) Validate() error {
	return dara.Validate(s)
}
