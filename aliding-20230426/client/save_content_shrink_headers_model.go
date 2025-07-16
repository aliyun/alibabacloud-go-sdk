// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContentShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SaveContentShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SaveContentShrinkHeaders
	GetAccountContextShrink() *string
}

type SaveContentShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SaveContentShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SaveContentShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SaveContentShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SaveContentShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SaveContentShrinkHeaders) SetCommonHeaders(v map[string]*string) *SaveContentShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveContentShrinkHeaders) SetAccountContextShrink(v string) *SaveContentShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SaveContentShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
