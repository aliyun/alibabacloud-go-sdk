// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIdListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInstanceIdListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetInstanceIdListShrinkHeaders
	GetAccountContextShrink() *string
}

type GetInstanceIdListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetInstanceIdListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIdListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInstanceIdListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetInstanceIdListShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceIdListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceIdListShrinkHeaders) SetAccountContextShrink(v string) *GetInstanceIdListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetInstanceIdListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
