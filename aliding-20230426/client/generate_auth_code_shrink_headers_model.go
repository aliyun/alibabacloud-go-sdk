// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAuthCodeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GenerateAuthCodeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GenerateAuthCodeShrinkHeaders
	GetAccountContextShrink() *string
}

type GenerateAuthCodeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GenerateAuthCodeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GenerateAuthCodeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GenerateAuthCodeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GenerateAuthCodeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GenerateAuthCodeShrinkHeaders) SetCommonHeaders(v map[string]*string) *GenerateAuthCodeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GenerateAuthCodeShrinkHeaders) SetAccountContextShrink(v string) *GenerateAuthCodeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GenerateAuthCodeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
