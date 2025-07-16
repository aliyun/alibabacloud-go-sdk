// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRangeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetRangeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetRangeShrinkHeaders
	GetAccountContextShrink() *string
}

type GetRangeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetRangeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetRangeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetRangeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetRangeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetRangeShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetRangeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRangeShrinkHeaders) SetAccountContextShrink(v string) *GetRangeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetRangeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
