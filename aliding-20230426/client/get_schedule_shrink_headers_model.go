// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetScheduleShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetScheduleShrinkHeaders
	GetAccountContextShrink() *string
}

type GetScheduleShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetScheduleShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetScheduleShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetScheduleShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetScheduleShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetScheduleShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetScheduleShrinkHeaders) SetAccountContextShrink(v string) *GetScheduleShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetScheduleShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
