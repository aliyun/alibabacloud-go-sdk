// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetOpenUrlRequest
	GetAppType() *string
	SetFileUrl(v string) *GetOpenUrlRequest
	GetFileUrl() *string
	SetLanguage(v string) *GetOpenUrlRequest
	GetLanguage() *string
	SetSystemToken(v string) *GetOpenUrlRequest
	GetSystemToken() *string
	SetTimeout(v int64) *GetOpenUrlRequest
	GetTimeout() *int64
}

type GetOpenUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0xxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliwork.com/fileHandle?appType=APP_VN7I6xxx&fileName=fileName.xlsx&instId=&type=download
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	// example:
	//
	// 60000L
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetOpenUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpenUrlRequest) GoString() string {
	return s.String()
}

func (s *GetOpenUrlRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetOpenUrlRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetOpenUrlRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetOpenUrlRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetOpenUrlRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetOpenUrlRequest) SetAppType(v string) *GetOpenUrlRequest {
	s.AppType = &v
	return s
}

func (s *GetOpenUrlRequest) SetFileUrl(v string) *GetOpenUrlRequest {
	s.FileUrl = &v
	return s
}

func (s *GetOpenUrlRequest) SetLanguage(v string) *GetOpenUrlRequest {
	s.Language = &v
	return s
}

func (s *GetOpenUrlRequest) SetSystemToken(v string) *GetOpenUrlRequest {
	s.SystemToken = &v
	return s
}

func (s *GetOpenUrlRequest) SetTimeout(v int64) *GetOpenUrlRequest {
	s.Timeout = &v
	return s
}

func (s *GetOpenUrlRequest) Validate() error {
	return dara.Validate(s)
}
