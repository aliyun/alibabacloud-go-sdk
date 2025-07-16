// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSheetShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSheetShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateSheetShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateSheetShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateSheetShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSheetShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateSheetShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSheetShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateSheetShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateSheetShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSheetShrinkHeaders) SetAccountContextShrink(v string) *CreateSheetShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateSheetShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
