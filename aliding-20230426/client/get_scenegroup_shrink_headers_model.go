// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenegroupShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetScenegroupShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetScenegroupShrinkHeaders
	GetAccountContextShrink() *string
}

type GetScenegroupShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetScenegroupShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetScenegroupShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetScenegroupShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetScenegroupShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetScenegroupShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetScenegroupShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetScenegroupShrinkHeaders) SetAccountContextShrink(v string) *GetScenegroupShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetScenegroupShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
