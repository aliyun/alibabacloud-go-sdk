// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiDimTableRecordsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListMultiDimTableRecordsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListMultiDimTableRecordsShrinkHeaders
	GetAccountContextShrink() *string
}

type ListMultiDimTableRecordsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListMultiDimTableRecordsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListMultiDimTableRecordsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListMultiDimTableRecordsShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListMultiDimTableRecordsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMultiDimTableRecordsShrinkHeaders) SetAccountContextShrink(v string) *ListMultiDimTableRecordsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListMultiDimTableRecordsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
