// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLangResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserLangResponseBody
	GetRequestId() *string
	SetSasUserLang(v *GetUserLangResponseBodySasUserLang) *GetUserLangResponseBody
	GetSasUserLang() *GetUserLangResponseBodySasUserLang
}

type GetUserLangResponseBody struct {
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SasUserLang *GetUserLangResponseBodySasUserLang `json:"SasUserLang,omitempty" xml:"SasUserLang,omitempty" type:"Struct"`
}

func (s GetUserLangResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserLangResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserLangResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserLangResponseBody) GetSasUserLang() *GetUserLangResponseBodySasUserLang {
	return s.SasUserLang
}

func (s *GetUserLangResponseBody) SetRequestId(v string) *GetUserLangResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserLangResponseBody) SetSasUserLang(v *GetUserLangResponseBodySasUserLang) *GetUserLangResponseBody {
	s.SasUserLang = v
	return s
}

func (s *GetUserLangResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserLangResponseBodySasUserLang struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetUserLangResponseBodySasUserLang) String() string {
	return dara.Prettify(s)
}

func (s GetUserLangResponseBodySasUserLang) GoString() string {
	return s.String()
}

func (s *GetUserLangResponseBodySasUserLang) GetLang() *string {
	return s.Lang
}

func (s *GetUserLangResponseBodySasUserLang) SetLang(v string) *GetUserLangResponseBodySasUserLang {
	s.Lang = &v
	return s
}

func (s *GetUserLangResponseBodySasUserLang) Validate() error {
	return dara.Validate(s)
}
