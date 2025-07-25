// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdnsAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreatePdnsAppKeyRequest
	GetLang() *string
	SetRemark(v string) *CreatePdnsAppKeyRequest
	GetRemark() *string
}

type CreatePdnsAppKeyRequest struct {
	Lang   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreatePdnsAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePdnsAppKeyRequest) GoString() string {
	return s.String()
}

func (s *CreatePdnsAppKeyRequest) GetLang() *string {
	return s.Lang
}

func (s *CreatePdnsAppKeyRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreatePdnsAppKeyRequest) SetLang(v string) *CreatePdnsAppKeyRequest {
	s.Lang = &v
	return s
}

func (s *CreatePdnsAppKeyRequest) SetRemark(v string) *CreatePdnsAppKeyRequest {
	s.Remark = &v
	return s
}

func (s *CreatePdnsAppKeyRequest) Validate() error {
	return dara.Validate(s)
}
