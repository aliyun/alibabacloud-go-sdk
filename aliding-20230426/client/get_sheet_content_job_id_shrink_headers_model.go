// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetContentJobIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSheetContentJobIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetSheetContentJobIdShrinkHeaders
	GetAccountContextShrink() *string
}

type GetSheetContentJobIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetSheetContentJobIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSheetContentJobIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetSheetContentJobIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSheetContentJobIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetSheetContentJobIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetSheetContentJobIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSheetContentJobIdShrinkHeaders) SetAccountContextShrink(v string) *GetSheetContentJobIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetSheetContentJobIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
