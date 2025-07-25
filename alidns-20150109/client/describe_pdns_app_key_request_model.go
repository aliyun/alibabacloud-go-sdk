// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKeyId(v string) *DescribePdnsAppKeyRequest
	GetAppKeyId() *string
	SetAuthCode(v string) *DescribePdnsAppKeyRequest
	GetAuthCode() *string
	SetLang(v string) *DescribePdnsAppKeyRequest
	GetLang() *string
}

type DescribePdnsAppKeyRequest struct {
	AppKeyId *string `json:"AppKeyId,omitempty" xml:"AppKeyId,omitempty"`
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePdnsAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAppKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsAppKeyRequest) GetAppKeyId() *string {
	return s.AppKeyId
}

func (s *DescribePdnsAppKeyRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePdnsAppKeyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsAppKeyRequest) SetAppKeyId(v string) *DescribePdnsAppKeyRequest {
	s.AppKeyId = &v
	return s
}

func (s *DescribePdnsAppKeyRequest) SetAuthCode(v string) *DescribePdnsAppKeyRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePdnsAppKeyRequest) SetLang(v string) *DescribePdnsAppKeyRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsAppKeyRequest) Validate() error {
	return dara.Validate(s)
}
