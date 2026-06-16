// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyAppKeyRequest
	GetLang() *string
	SetAppKey(v string) *ModifyAppKeyRequest
	GetAppKey() *string
	SetMemo(v string) *ModifyAppKeyRequest
	GetMemo() *string
	SetRegId(v string) *ModifyAppKeyRequest
	GetRegId() *string
}

type ModifyAppKeyRequest struct {
	// Specifies the language type of the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The AppKey information.
	//
	// example:
	//
	// ***
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// The memo information of the application.
	//
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s ModifyAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppKeyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppKeyRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyAppKeyRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *ModifyAppKeyRequest) GetMemo() *string {
	return s.Memo
}

func (s *ModifyAppKeyRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModifyAppKeyRequest) SetLang(v string) *ModifyAppKeyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyAppKeyRequest) SetAppKey(v string) *ModifyAppKeyRequest {
	s.AppKey = &v
	return s
}

func (s *ModifyAppKeyRequest) SetMemo(v string) *ModifyAppKeyRequest {
	s.Memo = &v
	return s
}

func (s *ModifyAppKeyRequest) SetRegId(v string) *ModifyAppKeyRequest {
	s.RegId = &v
	return s
}

func (s *ModifyAppKeyRequest) Validate() error {
	return dara.Validate(s)
}
