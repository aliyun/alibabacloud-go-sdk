// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataIdListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *SearchFormDataIdListRequest
	GetAppType() *string
	SetCreateFromTimeGMT(v string) *SearchFormDataIdListRequest
	GetCreateFromTimeGMT() *string
	SetCreateToTimeGMT(v string) *SearchFormDataIdListRequest
	GetCreateToTimeGMT() *string
	SetFormUuid(v string) *SearchFormDataIdListRequest
	GetFormUuid() *string
	SetLanguage(v string) *SearchFormDataIdListRequest
	GetLanguage() *string
	SetModifiedFromTimeGMT(v string) *SearchFormDataIdListRequest
	GetModifiedFromTimeGMT() *string
	SetModifiedToTimeGMT(v string) *SearchFormDataIdListRequest
	GetModifiedToTimeGMT() *string
	SetOriginatorId(v string) *SearchFormDataIdListRequest
	GetOriginatorId() *string
	SetPageNumber(v int32) *SearchFormDataIdListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchFormDataIdListRequest
	GetPageSize() *int32
	SetSearchFieldJson(v string) *SearchFormDataIdListRequest
	GetSearchFieldJson() *string
	SetSystemToken(v string) *SearchFormDataIdListRequest
	GetSystemToken() *string
}

type SearchFormDataIdListRequest struct {
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
	// zh-CN
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {}
	SearchFieldJson *string `json:"SearchFieldJson,omitempty" xml:"SearchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s SearchFormDataIdListRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataIdListRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListRequest) GetAppType() *string {
	return s.AppType
}

func (s *SearchFormDataIdListRequest) GetCreateFromTimeGMT() *string {
	return s.CreateFromTimeGMT
}

func (s *SearchFormDataIdListRequest) GetCreateToTimeGMT() *string {
	return s.CreateToTimeGMT
}

func (s *SearchFormDataIdListRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *SearchFormDataIdListRequest) GetLanguage() *string {
	return s.Language
}

func (s *SearchFormDataIdListRequest) GetModifiedFromTimeGMT() *string {
	return s.ModifiedFromTimeGMT
}

func (s *SearchFormDataIdListRequest) GetModifiedToTimeGMT() *string {
	return s.ModifiedToTimeGMT
}

func (s *SearchFormDataIdListRequest) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *SearchFormDataIdListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchFormDataIdListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchFormDataIdListRequest) GetSearchFieldJson() *string {
	return s.SearchFieldJson
}

func (s *SearchFormDataIdListRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *SearchFormDataIdListRequest) SetAppType(v string) *SearchFormDataIdListRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetCreateFromTimeGMT(v string) *SearchFormDataIdListRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetCreateToTimeGMT(v string) *SearchFormDataIdListRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetFormUuid(v string) *SearchFormDataIdListRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetLanguage(v string) *SearchFormDataIdListRequest {
	s.Language = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataIdListRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetModifiedToTimeGMT(v string) *SearchFormDataIdListRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetOriginatorId(v string) *SearchFormDataIdListRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetPageNumber(v int32) *SearchFormDataIdListRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetPageSize(v int32) *SearchFormDataIdListRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetSearchFieldJson(v string) *SearchFormDataIdListRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetSystemToken(v string) *SearchFormDataIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataIdListRequest) Validate() error {
	return dara.Validate(s)
}
