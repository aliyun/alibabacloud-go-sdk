// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScenegroupShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateScenegroupShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateScenegroupShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateScenegroupShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateScenegroupShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateScenegroupShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateScenegroupShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateScenegroupShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateScenegroupShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateScenegroupShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScenegroupShrinkHeaders) SetAccountContextShrink(v string) *CreateScenegroupShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateScenegroupShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
