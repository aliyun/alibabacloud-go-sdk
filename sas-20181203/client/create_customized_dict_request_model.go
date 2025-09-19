// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedDictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateCustomizedDictRequest
	GetLang() *string
	SetOverride(v bool) *CreateCustomizedDictRequest
	GetOverride() *bool
	SetSourceIp(v string) *CreateCustomizedDictRequest
	GetSourceIp() *string
}

type CreateCustomizedDictRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Override *bool   `json:"Override,omitempty" xml:"Override,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 106.11.43.***
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateCustomizedDictRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedDictRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedDictRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateCustomizedDictRequest) GetOverride() *bool {
	return s.Override
}

func (s *CreateCustomizedDictRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateCustomizedDictRequest) SetLang(v string) *CreateCustomizedDictRequest {
	s.Lang = &v
	return s
}

func (s *CreateCustomizedDictRequest) SetOverride(v bool) *CreateCustomizedDictRequest {
	s.Override = &v
	return s
}

func (s *CreateCustomizedDictRequest) SetSourceIp(v string) *CreateCustomizedDictRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateCustomizedDictRequest) Validate() error {
	return dara.Validate(s)
}
