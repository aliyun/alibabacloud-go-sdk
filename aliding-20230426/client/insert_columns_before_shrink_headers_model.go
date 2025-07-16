// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertColumnsBeforeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertColumnsBeforeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InsertColumnsBeforeShrinkHeaders
	GetAccountContextShrink() *string
}

type InsertColumnsBeforeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InsertColumnsBeforeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertColumnsBeforeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertColumnsBeforeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InsertColumnsBeforeShrinkHeaders) SetCommonHeaders(v map[string]*string) *InsertColumnsBeforeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertColumnsBeforeShrinkHeaders) SetAccountContextShrink(v string) *InsertColumnsBeforeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InsertColumnsBeforeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
