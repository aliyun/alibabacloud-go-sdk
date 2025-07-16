// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertMultiDimTableRecordShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertMultiDimTableRecordShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InsertMultiDimTableRecordShrinkHeaders
	GetAccountContextShrink() *string
}

type InsertMultiDimTableRecordShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InsertMultiDimTableRecordShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertMultiDimTableRecordShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InsertMultiDimTableRecordShrinkHeaders) SetCommonHeaders(v map[string]*string) *InsertMultiDimTableRecordShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertMultiDimTableRecordShrinkHeaders) SetAccountContextShrink(v string) *InsertMultiDimTableRecordShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InsertMultiDimTableRecordShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
