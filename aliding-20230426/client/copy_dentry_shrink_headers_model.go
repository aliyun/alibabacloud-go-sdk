// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CopyDentryShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CopyDentryShrinkHeaders
	GetAccountContextShrink() *string
}

type CopyDentryShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CopyDentryShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CopyDentryShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CopyDentryShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CopyDentryShrinkHeaders) SetCommonHeaders(v map[string]*string) *CopyDentryShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyDentryShrinkHeaders) SetAccountContextShrink(v string) *CopyDentryShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CopyDentryShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
