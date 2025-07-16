// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRowsVisibilityShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SetRowsVisibilityShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SetRowsVisibilityShrinkHeaders
	GetAccountContextShrink() *string
}

type SetRowsVisibilityShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SetRowsVisibilityShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SetRowsVisibilityShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SetRowsVisibilityShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SetRowsVisibilityShrinkHeaders) SetCommonHeaders(v map[string]*string) *SetRowsVisibilityShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetRowsVisibilityShrinkHeaders) SetAccountContextShrink(v string) *SetRowsVisibilityShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SetRowsVisibilityShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
