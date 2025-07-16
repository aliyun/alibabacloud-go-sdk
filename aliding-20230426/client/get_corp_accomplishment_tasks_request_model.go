// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpAccomplishmentTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppTypes(v string) *GetCorpAccomplishmentTasksRequest
	GetAppTypes() *string
	SetCorpId(v string) *GetCorpAccomplishmentTasksRequest
	GetCorpId() *string
	SetCreateFromTimeGMT(v int64) *GetCorpAccomplishmentTasksRequest
	GetCreateFromTimeGMT() *int64
	SetCreateToTimeGMT(v int64) *GetCorpAccomplishmentTasksRequest
	GetCreateToTimeGMT() *int64
	SetKeyword(v string) *GetCorpAccomplishmentTasksRequest
	GetKeyword() *string
	SetLanguage(v string) *GetCorpAccomplishmentTasksRequest
	GetLanguage() *string
	SetPageNumber(v int32) *GetCorpAccomplishmentTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetCorpAccomplishmentTasksRequest
	GetPageSize() *int32
	SetProcessCodes(v string) *GetCorpAccomplishmentTasksRequest
	GetProcessCodes() *string
	SetToken(v string) *GetCorpAccomplishmentTasksRequest
	GetToken() *string
}

type GetCorpAccomplishmentTasksRequest struct {
	// example:
	//
	// APP_PBKTxxx
	AppTypes *string `json:"AppTypes,omitempty" xml:"AppTypes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// corpIdxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateFromTimeGMT *int64 `json:"CreateFromTimeGMT,omitempty" xml:"CreateFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateToTimeGMT *int64 `json:"CreateToTimeGMT,omitempty" xml:"CreateToTimeGMT,omitempty"`
	// example:
	//
	// **
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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
	// ["xx","xxx"]
	ProcessCodes *string `json:"ProcessCodes,omitempty" xml:"ProcessCodes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetCorpAccomplishmentTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCorpAccomplishmentTasksRequest) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksRequest) GetAppTypes() *string {
	return s.AppTypes
}

func (s *GetCorpAccomplishmentTasksRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *GetCorpAccomplishmentTasksRequest) GetCreateFromTimeGMT() *int64 {
	return s.CreateFromTimeGMT
}

func (s *GetCorpAccomplishmentTasksRequest) GetCreateToTimeGMT() *int64 {
	return s.CreateToTimeGMT
}

func (s *GetCorpAccomplishmentTasksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetCorpAccomplishmentTasksRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetCorpAccomplishmentTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetCorpAccomplishmentTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCorpAccomplishmentTasksRequest) GetProcessCodes() *string {
	return s.ProcessCodes
}

func (s *GetCorpAccomplishmentTasksRequest) GetToken() *string {
	return s.Token
}

func (s *GetCorpAccomplishmentTasksRequest) SetAppTypes(v string) *GetCorpAccomplishmentTasksRequest {
	s.AppTypes = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetCorpId(v string) *GetCorpAccomplishmentTasksRequest {
	s.CorpId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetCreateFromTimeGMT(v int64) *GetCorpAccomplishmentTasksRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetCreateToTimeGMT(v int64) *GetCorpAccomplishmentTasksRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetKeyword(v string) *GetCorpAccomplishmentTasksRequest {
	s.Keyword = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetLanguage(v string) *GetCorpAccomplishmentTasksRequest {
	s.Language = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetPageNumber(v int32) *GetCorpAccomplishmentTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetPageSize(v int32) *GetCorpAccomplishmentTasksRequest {
	s.PageSize = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetProcessCodes(v string) *GetCorpAccomplishmentTasksRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetToken(v string) *GetCorpAccomplishmentTasksRequest {
	s.Token = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) Validate() error {
	return dara.Validate(s)
}
