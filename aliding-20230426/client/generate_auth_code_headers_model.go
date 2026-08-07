// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAuthCodeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GenerateAuthCodeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GenerateAuthCodeHeadersAccountContext) *GenerateAuthCodeHeaders
	GetAccountContext() *GenerateAuthCodeHeadersAccountContext
}

type GenerateAuthCodeHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GenerateAuthCodeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GenerateAuthCodeHeaders) String() string {
	return dara.Prettify(s)
}

func (s GenerateAuthCodeHeaders) GoString() string {
	return s.String()
}

func (s *GenerateAuthCodeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GenerateAuthCodeHeaders) GetAccountContext() *GenerateAuthCodeHeadersAccountContext {
	return s.AccountContext
}

func (s *GenerateAuthCodeHeaders) SetCommonHeaders(v map[string]*string) *GenerateAuthCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GenerateAuthCodeHeaders) SetAccountContext(v *GenerateAuthCodeHeadersAccountContext) *GenerateAuthCodeHeaders {
	s.AccountContext = v
	return s
}

func (s *GenerateAuthCodeHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateAuthCodeHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GenerateAuthCodeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GenerateAuthCodeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GenerateAuthCodeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GenerateAuthCodeHeadersAccountContext) SetAccountId(v string) *GenerateAuthCodeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GenerateAuthCodeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
