// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableRecordsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableRecordsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateMultiDimTableRecordsShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateMultiDimTableRecordsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateMultiDimTableRecordsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMultiDimTableRecordsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateMultiDimTableRecordsShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableRecordsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMultiDimTableRecordsShrinkHeaders) SetAccountContextShrink(v string) *UpdateMultiDimTableRecordsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateMultiDimTableRecordsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
