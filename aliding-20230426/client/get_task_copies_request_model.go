// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskCopiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetTaskCopiesRequest
	GetAppType() *string
	SetCreateFromTimeGMT(v int64) *GetTaskCopiesRequest
	GetCreateFromTimeGMT() *int64
	SetCreateToTimeGMT(v int64) *GetTaskCopiesRequest
	GetCreateToTimeGMT() *int64
	SetKeyword(v string) *GetTaskCopiesRequest
	GetKeyword() *string
	SetLanguage(v string) *GetTaskCopiesRequest
	GetLanguage() *string
	SetPageNumber(v int32) *GetTaskCopiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetTaskCopiesRequest
	GetPageSize() *int32
	SetProcessCodes(v string) *GetTaskCopiesRequest
	GetProcessCodes() *string
	SetSystemToken(v string) *GetTaskCopiesRequest
	GetSystemToken() *string
}

type GetTaskCopiesRequest struct {
	// example:
	//
	// APP_PBKxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
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
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s GetTaskCopiesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskCopiesRequest) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetTaskCopiesRequest) GetCreateFromTimeGMT() *int64 {
	return s.CreateFromTimeGMT
}

func (s *GetTaskCopiesRequest) GetCreateToTimeGMT() *int64 {
	return s.CreateToTimeGMT
}

func (s *GetTaskCopiesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetTaskCopiesRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetTaskCopiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetTaskCopiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTaskCopiesRequest) GetProcessCodes() *string {
	return s.ProcessCodes
}

func (s *GetTaskCopiesRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetTaskCopiesRequest) SetAppType(v string) *GetTaskCopiesRequest {
	s.AppType = &v
	return s
}

func (s *GetTaskCopiesRequest) SetCreateFromTimeGMT(v int64) *GetTaskCopiesRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetTaskCopiesRequest) SetCreateToTimeGMT(v int64) *GetTaskCopiesRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetTaskCopiesRequest) SetKeyword(v string) *GetTaskCopiesRequest {
	s.Keyword = &v
	return s
}

func (s *GetTaskCopiesRequest) SetLanguage(v string) *GetTaskCopiesRequest {
	s.Language = &v
	return s
}

func (s *GetTaskCopiesRequest) SetPageNumber(v int32) *GetTaskCopiesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTaskCopiesRequest) SetPageSize(v int32) *GetTaskCopiesRequest {
	s.PageSize = &v
	return s
}

func (s *GetTaskCopiesRequest) SetProcessCodes(v string) *GetTaskCopiesRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetTaskCopiesRequest) SetSystemToken(v string) *GetTaskCopiesRequest {
	s.SystemToken = &v
	return s
}

func (s *GetTaskCopiesRequest) Validate() error {
	return dara.Validate(s)
}
