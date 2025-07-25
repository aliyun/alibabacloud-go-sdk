// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePdnsAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKeyId(v string) *RemovePdnsAppKeyRequest
	GetAppKeyId() *string
	SetLang(v string) *RemovePdnsAppKeyRequest
	GetLang() *string
}

type RemovePdnsAppKeyRequest struct {
	AppKeyId *string `json:"AppKeyId,omitempty" xml:"AppKeyId,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s RemovePdnsAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePdnsAppKeyRequest) GoString() string {
	return s.String()
}

func (s *RemovePdnsAppKeyRequest) GetAppKeyId() *string {
	return s.AppKeyId
}

func (s *RemovePdnsAppKeyRequest) GetLang() *string {
	return s.Lang
}

func (s *RemovePdnsAppKeyRequest) SetAppKeyId(v string) *RemovePdnsAppKeyRequest {
	s.AppKeyId = &v
	return s
}

func (s *RemovePdnsAppKeyRequest) SetLang(v string) *RemovePdnsAppKeyRequest {
	s.Lang = &v
	return s
}

func (s *RemovePdnsAppKeyRequest) Validate() error {
	return dara.Validate(s)
}
