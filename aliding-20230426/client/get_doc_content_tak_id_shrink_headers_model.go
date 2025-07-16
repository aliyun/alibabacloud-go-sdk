// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentTakIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDocContentTakIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDocContentTakIdShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDocContentTakIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDocContentTakIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentTakIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDocContentTakIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDocContentTakIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDocContentTakIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDocContentTakIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDocContentTakIdShrinkHeaders) SetAccountContextShrink(v string) *GetDocContentTakIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDocContentTakIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
