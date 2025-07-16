// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataSecondGenerationNoTableFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetAppType() *string
	SetCreateFromTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetCreateFromTimeGMT() *string
	SetCreateToTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetCreateToTimeGMT() *string
	SetFormUuid(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetFormUuid() *string
	SetModifiedFromTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetModifiedFromTimeGMT() *string
	SetModifiedToTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetModifiedToTimeGMT() *string
	SetOrderConfigJson(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetOrderConfigJson() *string
	SetOriginatorId(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetOriginatorId() *string
	SetPageNumber(v int32) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetPageSize() *int32
	SetSearchCondition(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetSearchCondition() *string
	SetSystemToken(v string) *SearchFormDataSecondGenerationNoTableFieldRequest
	GetSystemToken() *string
}

type SearchFormDataSecondGenerationNoTableFieldRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0xxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// yyyy-MM-dd
	CreateFromTimeGMT *string `json:"CreateFromTimeGMT,omitempty" xml:"CreateFromTimeGMT,omitempty"`
	// example:
	//
	// yyyy-MM-dd
	CreateToTimeGMT *string `json:"CreateToTimeGMT,omitempty" xml:"CreateToTimeGMT,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-xxxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// yyyy-MM-dd
	ModifiedFromTimeGMT *string `json:"ModifiedFromTimeGMT,omitempty" xml:"ModifiedFromTimeGMT,omitempty"`
	// example:
	//
	// yyyy-MM-dd
	ModifiedToTimeGMT *string `json:"ModifiedToTimeGMT,omitempty" xml:"ModifiedToTimeGMT,omitempty"`
	// example:
	//
	// {}
	OrderConfigJson *string `json:"OrderConfigJson,omitempty" xml:"OrderConfigJson,omitempty"`
	// example:
	//
	// 012345
	OriginatorId *string `json:"OriginatorId,omitempty" xml:"OriginatorId,omitempty"`
	// example:
	//
	// 20
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {}
	SearchCondition *string `json:"SearchCondition,omitempty" xml:"SearchCondition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetAppType() *string {
	return s.AppType
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetCreateFromTimeGMT() *string {
	return s.CreateFromTimeGMT
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetCreateToTimeGMT() *string {
	return s.CreateToTimeGMT
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetModifiedFromTimeGMT() *string {
	return s.ModifiedFromTimeGMT
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetModifiedToTimeGMT() *string {
	return s.ModifiedToTimeGMT
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetOrderConfigJson() *string {
	return s.OrderConfigJson
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetSearchCondition() *string {
	return s.SearchCondition
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetAppType(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetCreateFromTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetCreateToTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetFormUuid(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetModifiedToTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetOrderConfigJson(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetOriginatorId(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetPageNumber(v int32) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetPageSize(v int32) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetSearchCondition(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.SearchCondition = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetSystemToken(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) Validate() error {
	return dara.Validate(s)
}
