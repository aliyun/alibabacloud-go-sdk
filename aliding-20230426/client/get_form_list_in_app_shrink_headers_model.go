// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormListInAppShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFormListInAppShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetFormListInAppShrinkHeaders
	GetAccountContextShrink() *string
}

type GetFormListInAppShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetFormListInAppShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFormListInAppShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetFormListInAppShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFormListInAppShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetFormListInAppShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetFormListInAppShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormListInAppShrinkHeaders) SetAccountContextShrink(v string) *GetFormListInAppShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetFormListInAppShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
