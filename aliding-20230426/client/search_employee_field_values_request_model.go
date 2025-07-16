// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEmployeeFieldValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *SearchEmployeeFieldValuesRequest
	GetAppType() *string
	SetCreateFromTimeGMT(v string) *SearchEmployeeFieldValuesRequest
	GetCreateFromTimeGMT() *string
	SetCreateToTimeGMT(v string) *SearchEmployeeFieldValuesRequest
	GetCreateToTimeGMT() *string
	SetFormUuid(v string) *SearchEmployeeFieldValuesRequest
	GetFormUuid() *string
	SetLanguage(v string) *SearchEmployeeFieldValuesRequest
	GetLanguage() *string
	SetModifiedFromTimeGMT(v string) *SearchEmployeeFieldValuesRequest
	GetModifiedFromTimeGMT() *string
	SetModifiedToTimeGMT(v string) *SearchEmployeeFieldValuesRequest
	GetModifiedToTimeGMT() *string
	SetOriginatorId(v string) *SearchEmployeeFieldValuesRequest
	GetOriginatorId() *string
	SetSearchFieldJson(v string) *SearchEmployeeFieldValuesRequest
	GetSearchFieldJson() *string
	SetSystemToken(v string) *SearchEmployeeFieldValuesRequest
	GetSystemToken() *string
	SetTargetFieldJson(v string) *SearchEmployeeFieldValuesRequest
	GetTargetFieldJson() *string
}

type SearchEmployeeFieldValuesRequest struct {
	// example:
	//
	// APP_PBKxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateFromTimeGMT *string `json:"CreateFromTimeGMT,omitempty" xml:"CreateFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateToTimeGMT *string `json:"CreateToTimeGMT,omitempty" xml:"CreateToTimeGMT,omitempty"`
	// example:
	//
	// FORM-EF6Yxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 2021-05-01
	ModifiedFromTimeGMT *string `json:"ModifiedFromTimeGMT,omitempty" xml:"ModifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-09-10
	ModifiedToTimeGMT *string `json:"ModifiedToTimeGMT,omitempty" xml:"ModifiedToTimeGMT,omitempty"`
	// example:
	//
	// 112212221
	OriginatorId *string `json:"OriginatorId,omitempty" xml:"OriginatorId,omitempty"`
	// example:
	//
	// {\\"textField_annandfa\\":\\"1212\\"}
	SearchFieldJson *string `json:"SearchFieldJson,omitempty" xml:"SearchFieldJson,omitempty"`
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	// example:
	//
	// [\\"textField_xahdfna\\"]
	TargetFieldJson *string `json:"TargetFieldJson,omitempty" xml:"TargetFieldJson,omitempty"`
}

func (s SearchEmployeeFieldValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchEmployeeFieldValuesRequest) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesRequest) GetAppType() *string {
	return s.AppType
}

func (s *SearchEmployeeFieldValuesRequest) GetCreateFromTimeGMT() *string {
	return s.CreateFromTimeGMT
}

func (s *SearchEmployeeFieldValuesRequest) GetCreateToTimeGMT() *string {
	return s.CreateToTimeGMT
}

func (s *SearchEmployeeFieldValuesRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *SearchEmployeeFieldValuesRequest) GetLanguage() *string {
	return s.Language
}

func (s *SearchEmployeeFieldValuesRequest) GetModifiedFromTimeGMT() *string {
	return s.ModifiedFromTimeGMT
}

func (s *SearchEmployeeFieldValuesRequest) GetModifiedToTimeGMT() *string {
	return s.ModifiedToTimeGMT
}

func (s *SearchEmployeeFieldValuesRequest) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *SearchEmployeeFieldValuesRequest) GetSearchFieldJson() *string {
	return s.SearchFieldJson
}

func (s *SearchEmployeeFieldValuesRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *SearchEmployeeFieldValuesRequest) GetTargetFieldJson() *string {
	return s.TargetFieldJson
}

func (s *SearchEmployeeFieldValuesRequest) SetAppType(v string) *SearchEmployeeFieldValuesRequest {
	s.AppType = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetCreateFromTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetCreateToTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetFormUuid(v string) *SearchEmployeeFieldValuesRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetLanguage(v string) *SearchEmployeeFieldValuesRequest {
	s.Language = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetModifiedFromTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetModifiedToTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetOriginatorId(v string) *SearchEmployeeFieldValuesRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetSearchFieldJson(v string) *SearchEmployeeFieldValuesRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetSystemToken(v string) *SearchEmployeeFieldValuesRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetTargetFieldJson(v string) *SearchEmployeeFieldValuesRequest {
	s.TargetFieldJson = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) Validate() error {
	return dara.Validate(s)
}
