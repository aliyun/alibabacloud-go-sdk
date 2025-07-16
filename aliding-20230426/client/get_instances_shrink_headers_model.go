// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInstancesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetInstancesShrinkHeaders
	GetAccountContextShrink() *string
}

type GetInstancesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetInstancesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetInstancesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInstancesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetInstancesShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetInstancesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstancesShrinkHeaders) SetAccountContextShrink(v string) *GetInstancesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetInstancesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
