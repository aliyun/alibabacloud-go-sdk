// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryByNodeIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CopyDentryByNodeIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CopyDentryByNodeIdShrinkHeaders
	GetAccountContextShrink() *string
}

type CopyDentryByNodeIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CopyDentryByNodeIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryByNodeIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CopyDentryByNodeIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CopyDentryByNodeIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CopyDentryByNodeIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *CopyDentryByNodeIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyDentryByNodeIdShrinkHeaders) SetAccountContextShrink(v string) *CopyDentryByNodeIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CopyDentryByNodeIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
