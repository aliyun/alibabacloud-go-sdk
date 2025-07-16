// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *SaveFormRemarkRequest
	GetAppType() *string
	SetAtUserId(v string) *SaveFormRemarkRequest
	GetAtUserId() *string
	SetContent(v string) *SaveFormRemarkRequest
	GetContent() *string
	SetFormInstanceId(v string) *SaveFormRemarkRequest
	GetFormInstanceId() *string
	SetLanguage(v string) *SaveFormRemarkRequest
	GetLanguage() *string
	SetReplyId(v int64) *SaveFormRemarkRequest
	GetReplyId() *int64
	SetSystemToken(v string) *SaveFormRemarkRequest
	GetSystemToken() *string
}

type SaveFormRemarkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 123456
	AtUserId *string `json:"AtUserId,omitempty" xml:"AtUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// instxxxx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// instxxxx
	FormInstanceId *string `json:"FormInstanceId,omitempty" xml:"FormInstanceId,omitempty"`
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// hexxxx
	ReplyId *int64 `json:"ReplyId,omitempty" xml:"ReplyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s SaveFormRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveFormRemarkRequest) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkRequest) GetAppType() *string {
	return s.AppType
}

func (s *SaveFormRemarkRequest) GetAtUserId() *string {
	return s.AtUserId
}

func (s *SaveFormRemarkRequest) GetContent() *string {
	return s.Content
}

func (s *SaveFormRemarkRequest) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *SaveFormRemarkRequest) GetLanguage() *string {
	return s.Language
}

func (s *SaveFormRemarkRequest) GetReplyId() *int64 {
	return s.ReplyId
}

func (s *SaveFormRemarkRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *SaveFormRemarkRequest) SetAppType(v string) *SaveFormRemarkRequest {
	s.AppType = &v
	return s
}

func (s *SaveFormRemarkRequest) SetAtUserId(v string) *SaveFormRemarkRequest {
	s.AtUserId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetContent(v string) *SaveFormRemarkRequest {
	s.Content = &v
	return s
}

func (s *SaveFormRemarkRequest) SetFormInstanceId(v string) *SaveFormRemarkRequest {
	s.FormInstanceId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetLanguage(v string) *SaveFormRemarkRequest {
	s.Language = &v
	return s
}

func (s *SaveFormRemarkRequest) SetReplyId(v int64) *SaveFormRemarkRequest {
	s.ReplyId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetSystemToken(v string) *SaveFormRemarkRequest {
	s.SystemToken = &v
	return s
}

func (s *SaveFormRemarkRequest) Validate() error {
	return dara.Validate(s)
}
