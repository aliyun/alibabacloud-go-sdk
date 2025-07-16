// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertRowsBeforeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertRowsBeforeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InsertRowsBeforeShrinkHeaders
	GetAccountContextShrink() *string
}

type InsertRowsBeforeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InsertRowsBeforeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertRowsBeforeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertRowsBeforeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InsertRowsBeforeShrinkHeaders) SetCommonHeaders(v map[string]*string) *InsertRowsBeforeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertRowsBeforeShrinkHeaders) SetAccountContextShrink(v string) *InsertRowsBeforeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InsertRowsBeforeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
