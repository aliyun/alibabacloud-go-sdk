// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataSecondGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *SearchFormDataSecondGenerationRequest
	GetAppType() *string
	SetCreateFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest
	GetCreateFromTimeGMT() *string
	SetCreateToTimeGMT(v string) *SearchFormDataSecondGenerationRequest
	GetCreateToTimeGMT() *string
	SetFormUuid(v string) *SearchFormDataSecondGenerationRequest
	GetFormUuid() *string
	SetModifiedFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest
	GetModifiedFromTimeGMT() *string
	SetModifiedToTimeGMT(v string) *SearchFormDataSecondGenerationRequest
	GetModifiedToTimeGMT() *string
	SetOrderConfigJson(v string) *SearchFormDataSecondGenerationRequest
	GetOrderConfigJson() *string
	SetOriginatorId(v string) *SearchFormDataSecondGenerationRequest
	GetOriginatorId() *string
	SetPageNumber(v int32) *SearchFormDataSecondGenerationRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchFormDataSecondGenerationRequest
	GetPageSize() *int32
	SetSearchCondition(v string) *SearchFormDataSecondGenerationRequest
	GetSearchCondition() *string
	SetSystemToken(v string) *SearchFormDataSecondGenerationRequest
	GetSystemToken() *string
}

type SearchFormDataSecondGenerationRequest struct {
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

func (s SearchFormDataSecondGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationRequest) GetAppType() *string {
	return s.AppType
}

func (s *SearchFormDataSecondGenerationRequest) GetCreateFromTimeGMT() *string {
	return s.CreateFromTimeGMT
}

func (s *SearchFormDataSecondGenerationRequest) GetCreateToTimeGMT() *string {
	return s.CreateToTimeGMT
}

func (s *SearchFormDataSecondGenerationRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *SearchFormDataSecondGenerationRequest) GetModifiedFromTimeGMT() *string {
	return s.ModifiedFromTimeGMT
}

func (s *SearchFormDataSecondGenerationRequest) GetModifiedToTimeGMT() *string {
	return s.ModifiedToTimeGMT
}

func (s *SearchFormDataSecondGenerationRequest) GetOrderConfigJson() *string {
	return s.OrderConfigJson
}

func (s *SearchFormDataSecondGenerationRequest) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *SearchFormDataSecondGenerationRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchFormDataSecondGenerationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchFormDataSecondGenerationRequest) GetSearchCondition() *string {
	return s.SearchCondition
}

func (s *SearchFormDataSecondGenerationRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *SearchFormDataSecondGenerationRequest) SetAppType(v string) *SearchFormDataSecondGenerationRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetCreateFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetCreateToTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetFormUuid(v string) *SearchFormDataSecondGenerationRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetModifiedToTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetOrderConfigJson(v string) *SearchFormDataSecondGenerationRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetOriginatorId(v string) *SearchFormDataSecondGenerationRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetPageNumber(v int32) *SearchFormDataSecondGenerationRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetPageSize(v int32) *SearchFormDataSecondGenerationRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetSearchCondition(v string) *SearchFormDataSecondGenerationRequest {
	s.SearchCondition = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetSystemToken(v string) *SearchFormDataSecondGenerationRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) Validate() error {
	return dara.Validate(s)
}
