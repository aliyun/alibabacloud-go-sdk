// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInCheckMailTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *TransferInCheckMailTokenRequest
	GetLang() *string
	SetToken(v string) *TransferInCheckMailTokenRequest
	GetToken() *string
	SetUserClientIp(v string) *TransferInCheckMailTokenRequest
	GetUserClientIp() *string
}

type TransferInCheckMailTokenRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3bdbaa0e-faa2-4ad2-98f4-bcfeb0237054
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s TransferInCheckMailTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferInCheckMailTokenRequest) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenRequest) GetLang() *string {
	return s.Lang
}

func (s *TransferInCheckMailTokenRequest) GetToken() *string {
	return s.Token
}

func (s *TransferInCheckMailTokenRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *TransferInCheckMailTokenRequest) SetLang(v string) *TransferInCheckMailTokenRequest {
	s.Lang = &v
	return s
}

func (s *TransferInCheckMailTokenRequest) SetToken(v string) *TransferInCheckMailTokenRequest {
	s.Token = &v
	return s
}

func (s *TransferInCheckMailTokenRequest) SetUserClientIp(v string) *TransferInCheckMailTokenRequest {
	s.UserClientIp = &v
	return s
}

func (s *TransferInCheckMailTokenRequest) Validate() error {
	return dara.Validate(s)
}
