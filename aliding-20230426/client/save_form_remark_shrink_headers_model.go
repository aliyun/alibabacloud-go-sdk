// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormRemarkShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SaveFormRemarkShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SaveFormRemarkShrinkHeaders
	GetAccountContextShrink() *string
}

type SaveFormRemarkShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SaveFormRemarkShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SaveFormRemarkShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SaveFormRemarkShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SaveFormRemarkShrinkHeaders) SetCommonHeaders(v map[string]*string) *SaveFormRemarkShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormRemarkShrinkHeaders) SetAccountContextShrink(v string) *SaveFormRemarkShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SaveFormRemarkShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
