// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateAppKeyRequest
	GetLang() *string
	SetRegId(v string) *CreateAppKeyRequest
	GetRegId() *string
}

type CreateAppKeyRequest struct {
	// Set the language type for requests and responses, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s CreateAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAppKeyRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateAppKeyRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateAppKeyRequest) SetLang(v string) *CreateAppKeyRequest {
	s.Lang = &v
	return s
}

func (s *CreateAppKeyRequest) SetRegId(v string) *CreateAppKeyRequest {
	s.RegId = &v
	return s
}

func (s *CreateAppKeyRequest) Validate() error {
	return dara.Validate(s)
}
