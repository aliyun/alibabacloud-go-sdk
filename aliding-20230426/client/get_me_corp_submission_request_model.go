// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeCorpSubmissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppTypes(v string) *GetMeCorpSubmissionRequest
	GetAppTypes() *string
	SetCorpId(v string) *GetMeCorpSubmissionRequest
	GetCorpId() *string
	SetCreateFromTimeGMT(v int64) *GetMeCorpSubmissionRequest
	GetCreateFromTimeGMT() *int64
	SetCreateToTimeGMT(v int64) *GetMeCorpSubmissionRequest
	GetCreateToTimeGMT() *int64
	SetKeyword(v string) *GetMeCorpSubmissionRequest
	GetKeyword() *string
	SetLanguage(v string) *GetMeCorpSubmissionRequest
	GetLanguage() *string
	SetPageNumber(v int32) *GetMeCorpSubmissionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetMeCorpSubmissionRequest
	GetPageSize() *int32
	SetProcessCodes(v string) *GetMeCorpSubmissionRequest
	GetProcessCodes() *string
	SetToken(v string) *GetMeCorpSubmissionRequest
	GetToken() *string
}

type GetMeCorpSubmissionRequest struct {
	AppTypes          *string `json:"AppTypes,omitempty" xml:"AppTypes,omitempty"`
	CorpId            *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	CreateFromTimeGMT *int64  `json:"CreateFromTimeGMT,omitempty" xml:"CreateFromTimeGMT,omitempty"`
	CreateToTimeGMT   *int64  `json:"CreateToTimeGMT,omitempty" xml:"CreateToTimeGMT,omitempty"`
	Keyword           *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Language          *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProcessCodes      *string `json:"ProcessCodes,omitempty" xml:"ProcessCodes,omitempty"`
	Token             *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetMeCorpSubmissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMeCorpSubmissionRequest) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionRequest) GetAppTypes() *string {
	return s.AppTypes
}

func (s *GetMeCorpSubmissionRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *GetMeCorpSubmissionRequest) GetCreateFromTimeGMT() *int64 {
	return s.CreateFromTimeGMT
}

func (s *GetMeCorpSubmissionRequest) GetCreateToTimeGMT() *int64 {
	return s.CreateToTimeGMT
}

func (s *GetMeCorpSubmissionRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetMeCorpSubmissionRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetMeCorpSubmissionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMeCorpSubmissionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMeCorpSubmissionRequest) GetProcessCodes() *string {
	return s.ProcessCodes
}

func (s *GetMeCorpSubmissionRequest) GetToken() *string {
	return s.Token
}

func (s *GetMeCorpSubmissionRequest) SetAppTypes(v string) *GetMeCorpSubmissionRequest {
	s.AppTypes = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetCorpId(v string) *GetMeCorpSubmissionRequest {
	s.CorpId = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetCreateFromTimeGMT(v int64) *GetMeCorpSubmissionRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetCreateToTimeGMT(v int64) *GetMeCorpSubmissionRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetKeyword(v string) *GetMeCorpSubmissionRequest {
	s.Keyword = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetLanguage(v string) *GetMeCorpSubmissionRequest {
	s.Language = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetPageNumber(v int32) *GetMeCorpSubmissionRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetPageSize(v int32) *GetMeCorpSubmissionRequest {
	s.PageSize = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetProcessCodes(v string) *GetMeCorpSubmissionRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetToken(v string) *GetMeCorpSubmissionRequest {
	s.Token = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) Validate() error {
	return dara.Validate(s)
}
