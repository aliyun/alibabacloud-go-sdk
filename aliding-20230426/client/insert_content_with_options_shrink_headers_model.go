// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertContentWithOptionsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertContentWithOptionsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InsertContentWithOptionsShrinkHeaders
	GetAccountContextShrink() *string
}

type InsertContentWithOptionsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InsertContentWithOptionsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertContentWithOptionsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InsertContentWithOptionsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertContentWithOptionsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InsertContentWithOptionsShrinkHeaders) SetCommonHeaders(v map[string]*string) *InsertContentWithOptionsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertContentWithOptionsShrinkHeaders) SetAccountContextShrink(v string) *InsertContentWithOptionsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InsertContentWithOptionsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
