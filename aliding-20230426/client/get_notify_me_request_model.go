// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotifyMeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppTypes(v string) *GetNotifyMeRequest
	GetAppTypes() *string
	SetCorpId(v string) *GetNotifyMeRequest
	GetCorpId() *string
	SetCreateFromTimeGMT(v int64) *GetNotifyMeRequest
	GetCreateFromTimeGMT() *int64
	SetCreateToTimeGMT(v int64) *GetNotifyMeRequest
	GetCreateToTimeGMT() *int64
	SetInstanceCreateFromTimeGMT(v int64) *GetNotifyMeRequest
	GetInstanceCreateFromTimeGMT() *int64
	SetInstanceCreateToTimeGMT(v int64) *GetNotifyMeRequest
	GetInstanceCreateToTimeGMT() *int64
	SetKeyword(v string) *GetNotifyMeRequest
	GetKeyword() *string
	SetLanguage(v string) *GetNotifyMeRequest
	GetLanguage() *string
	SetPageNumber(v int32) *GetNotifyMeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetNotifyMeRequest
	GetPageSize() *int32
	SetProcessCodes(v string) *GetNotifyMeRequest
	GetProcessCodes() *string
	SetToken(v string) *GetNotifyMeRequest
	GetToken() *string
}

type GetNotifyMeRequest struct {
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
	// 2021-05-01
	InstanceCreateFromTimeGMT *int64 `json:"InstanceCreateFromTimeGMT,omitempty" xml:"InstanceCreateFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-05-01
	InstanceCreateToTimeGMT *int64 `json:"InstanceCreateToTimeGMT,omitempty" xml:"InstanceCreateToTimeGMT,omitempty"`
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

func (s GetNotifyMeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyMeRequest) GoString() string {
	return s.String()
}

func (s *GetNotifyMeRequest) GetAppTypes() *string {
	return s.AppTypes
}

func (s *GetNotifyMeRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *GetNotifyMeRequest) GetCreateFromTimeGMT() *int64 {
	return s.CreateFromTimeGMT
}

func (s *GetNotifyMeRequest) GetCreateToTimeGMT() *int64 {
	return s.CreateToTimeGMT
}

func (s *GetNotifyMeRequest) GetInstanceCreateFromTimeGMT() *int64 {
	return s.InstanceCreateFromTimeGMT
}

func (s *GetNotifyMeRequest) GetInstanceCreateToTimeGMT() *int64 {
	return s.InstanceCreateToTimeGMT
}

func (s *GetNotifyMeRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetNotifyMeRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetNotifyMeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetNotifyMeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetNotifyMeRequest) GetProcessCodes() *string {
	return s.ProcessCodes
}

func (s *GetNotifyMeRequest) GetToken() *string {
	return s.Token
}

func (s *GetNotifyMeRequest) SetAppTypes(v string) *GetNotifyMeRequest {
	s.AppTypes = &v
	return s
}

func (s *GetNotifyMeRequest) SetCorpId(v string) *GetNotifyMeRequest {
	s.CorpId = &v
	return s
}

func (s *GetNotifyMeRequest) SetCreateFromTimeGMT(v int64) *GetNotifyMeRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetNotifyMeRequest) SetCreateToTimeGMT(v int64) *GetNotifyMeRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetNotifyMeRequest) SetInstanceCreateFromTimeGMT(v int64) *GetNotifyMeRequest {
	s.InstanceCreateFromTimeGMT = &v
	return s
}

func (s *GetNotifyMeRequest) SetInstanceCreateToTimeGMT(v int64) *GetNotifyMeRequest {
	s.InstanceCreateToTimeGMT = &v
	return s
}

func (s *GetNotifyMeRequest) SetKeyword(v string) *GetNotifyMeRequest {
	s.Keyword = &v
	return s
}

func (s *GetNotifyMeRequest) SetLanguage(v string) *GetNotifyMeRequest {
	s.Language = &v
	return s
}

func (s *GetNotifyMeRequest) SetPageNumber(v int32) *GetNotifyMeRequest {
	s.PageNumber = &v
	return s
}

func (s *GetNotifyMeRequest) SetPageSize(v int32) *GetNotifyMeRequest {
	s.PageSize = &v
	return s
}

func (s *GetNotifyMeRequest) SetProcessCodes(v string) *GetNotifyMeRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetNotifyMeRequest) SetToken(v string) *GetNotifyMeRequest {
	s.Token = &v
	return s
}

func (s *GetNotifyMeRequest) Validate() error {
	return dara.Validate(s)
}
