// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableRecordShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMultiDimTableRecordShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetMultiDimTableRecordShrinkHeaders
	GetAccountContextShrink() *string
}

type GetMultiDimTableRecordShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetMultiDimTableRecordShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMultiDimTableRecordShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetMultiDimTableRecordShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetMultiDimTableRecordShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultiDimTableRecordShrinkHeaders) SetAccountContextShrink(v string) *GetMultiDimTableRecordShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetMultiDimTableRecordShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
