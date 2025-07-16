// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDatasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *SearchFormDatasRequest
	GetAppType() *string
	SetCreateFromTimeGMT(v string) *SearchFormDatasRequest
	GetCreateFromTimeGMT() *string
	SetCreateToTimeGMT(v string) *SearchFormDatasRequest
	GetCreateToTimeGMT() *string
	SetCurrentPage(v int32) *SearchFormDatasRequest
	GetCurrentPage() *int32
	SetDynamicOrder(v string) *SearchFormDatasRequest
	GetDynamicOrder() *string
	SetFormUuid(v string) *SearchFormDatasRequest
	GetFormUuid() *string
	SetLanguage(v string) *SearchFormDatasRequest
	GetLanguage() *string
	SetModifiedFromTimeGMT(v string) *SearchFormDatasRequest
	GetModifiedFromTimeGMT() *string
	SetModifiedToTimeGMT(v string) *SearchFormDatasRequest
	GetModifiedToTimeGMT() *string
	SetOriginatorId(v string) *SearchFormDatasRequest
	GetOriginatorId() *string
	SetPageSize(v int32) *SearchFormDatasRequest
	GetPageSize() *int32
	SetSearchFieldJson(v string) *SearchFormDatasRequest
	GetSearchFieldJson() *string
	SetSystemToken(v string) *SearchFormDatasRequest
	GetSystemToken() *string
}

type SearchFormDatasRequest struct {
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
	// example:
	//
	// 20
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	DynamicOrder *string `json:"DynamicOrder,omitempty" xml:"DynamicOrder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-xxxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// en_US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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
	// 012345
	OriginatorId *string `json:"OriginatorId,omitempty" xml:"OriginatorId,omitempty"`
	// example:
	//
	// 20
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchFieldJson *string `json:"SearchFieldJson,omitempty" xml:"SearchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s SearchFormDatasRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDatasRequest) GetAppType() *string {
	return s.AppType
}

func (s *SearchFormDatasRequest) GetCreateFromTimeGMT() *string {
	return s.CreateFromTimeGMT
}

func (s *SearchFormDatasRequest) GetCreateToTimeGMT() *string {
	return s.CreateToTimeGMT
}

func (s *SearchFormDatasRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *SearchFormDatasRequest) GetDynamicOrder() *string {
	return s.DynamicOrder
}

func (s *SearchFormDatasRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *SearchFormDatasRequest) GetLanguage() *string {
	return s.Language
}

func (s *SearchFormDatasRequest) GetModifiedFromTimeGMT() *string {
	return s.ModifiedFromTimeGMT
}

func (s *SearchFormDatasRequest) GetModifiedToTimeGMT() *string {
	return s.ModifiedToTimeGMT
}

func (s *SearchFormDatasRequest) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *SearchFormDatasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchFormDatasRequest) GetSearchFieldJson() *string {
	return s.SearchFieldJson
}

func (s *SearchFormDatasRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *SearchFormDatasRequest) SetAppType(v string) *SearchFormDatasRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateFromTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateToTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCurrentPage(v int32) *SearchFormDatasRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasRequest) SetDynamicOrder(v string) *SearchFormDatasRequest {
	s.DynamicOrder = &v
	return s
}

func (s *SearchFormDatasRequest) SetFormUuid(v string) *SearchFormDatasRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasRequest) SetLanguage(v string) *SearchFormDatasRequest {
	s.Language = &v
	return s
}

func (s *SearchFormDatasRequest) SetModifiedFromTimeGMT(v string) *SearchFormDatasRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetModifiedToTimeGMT(v string) *SearchFormDatasRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetOriginatorId(v string) *SearchFormDatasRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDatasRequest) SetPageSize(v int32) *SearchFormDatasRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDatasRequest) SetSearchFieldJson(v string) *SearchFormDatasRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDatasRequest) SetSystemToken(v string) *SearchFormDatasRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDatasRequest) Validate() error {
	return dara.Validate(s)
}
