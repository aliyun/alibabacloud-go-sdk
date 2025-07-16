// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppTypes(v string) *GetCorpTasksRequest
	GetAppTypes() *string
	SetCorpId(v string) *GetCorpTasksRequest
	GetCorpId() *string
	SetCreateFromTimeGMT(v int64) *GetCorpTasksRequest
	GetCreateFromTimeGMT() *int64
	SetCreateToTimeGMT(v int64) *GetCorpTasksRequest
	GetCreateToTimeGMT() *int64
	SetKeyword(v string) *GetCorpTasksRequest
	GetKeyword() *string
	SetLanguage(v string) *GetCorpTasksRequest
	GetLanguage() *string
	SetPageNumber(v int32) *GetCorpTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetCorpTasksRequest
	GetPageSize() *int32
	SetProcessCodes(v string) *GetCorpTasksRequest
	GetProcessCodes() *string
	SetToken(v string) *GetCorpTasksRequest
	GetToken() *string
}

type GetCorpTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKxxx
	AppTypes *string `json:"AppTypes,omitempty" xml:"AppTypes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// corpId
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
	// keyword
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

func (s GetCorpTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCorpTasksRequest) GoString() string {
	return s.String()
}

func (s *GetCorpTasksRequest) GetAppTypes() *string {
	return s.AppTypes
}

func (s *GetCorpTasksRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *GetCorpTasksRequest) GetCreateFromTimeGMT() *int64 {
	return s.CreateFromTimeGMT
}

func (s *GetCorpTasksRequest) GetCreateToTimeGMT() *int64 {
	return s.CreateToTimeGMT
}

func (s *GetCorpTasksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetCorpTasksRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetCorpTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetCorpTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCorpTasksRequest) GetProcessCodes() *string {
	return s.ProcessCodes
}

func (s *GetCorpTasksRequest) GetToken() *string {
	return s.Token
}

func (s *GetCorpTasksRequest) SetAppTypes(v string) *GetCorpTasksRequest {
	s.AppTypes = &v
	return s
}

func (s *GetCorpTasksRequest) SetCorpId(v string) *GetCorpTasksRequest {
	s.CorpId = &v
	return s
}

func (s *GetCorpTasksRequest) SetCreateFromTimeGMT(v int64) *GetCorpTasksRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetCorpTasksRequest) SetCreateToTimeGMT(v int64) *GetCorpTasksRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetCorpTasksRequest) SetKeyword(v string) *GetCorpTasksRequest {
	s.Keyword = &v
	return s
}

func (s *GetCorpTasksRequest) SetLanguage(v string) *GetCorpTasksRequest {
	s.Language = &v
	return s
}

func (s *GetCorpTasksRequest) SetPageNumber(v int32) *GetCorpTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *GetCorpTasksRequest) SetPageSize(v int32) *GetCorpTasksRequest {
	s.PageSize = &v
	return s
}

func (s *GetCorpTasksRequest) SetProcessCodes(v string) *GetCorpTasksRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetCorpTasksRequest) SetToken(v string) *GetCorpTasksRequest {
	s.Token = &v
	return s
}

func (s *GetCorpTasksRequest) Validate() error {
	return dara.Validate(s)
}
