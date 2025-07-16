// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDocContentShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDocContentShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDocContentShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDocContentShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDocContentShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDocContentShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDocContentShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDocContentShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDocContentShrinkHeaders) SetAccountContextShrink(v string) *GetDocContentShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDocContentShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
