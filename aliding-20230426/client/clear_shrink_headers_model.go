// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ClearShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ClearShrinkHeaders
	GetAccountContextShrink() *string
}

type ClearShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ClearShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ClearShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ClearShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ClearShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ClearShrinkHeaders) SetCommonHeaders(v map[string]*string) *ClearShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ClearShrinkHeaders) SetAccountContextShrink(v string) *ClearShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ClearShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
