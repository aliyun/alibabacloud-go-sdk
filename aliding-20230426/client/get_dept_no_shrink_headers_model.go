// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeptNoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeptNoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDeptNoShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDeptNoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDeptNoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeptNoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDeptNoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeptNoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDeptNoShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDeptNoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeptNoShrinkHeaders) SetAccountContextShrink(v string) *GetDeptNoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDeptNoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
