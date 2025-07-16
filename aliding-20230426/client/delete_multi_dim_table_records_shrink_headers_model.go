// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableRecordsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMultiDimTableRecordsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteMultiDimTableRecordsShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteMultiDimTableRecordsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteMultiDimTableRecordsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableRecordsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableRecordsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMultiDimTableRecordsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteMultiDimTableRecordsShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteMultiDimTableRecordsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMultiDimTableRecordsShrinkHeaders) SetAccountContextShrink(v string) *DeleteMultiDimTableRecordsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteMultiDimTableRecordsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
