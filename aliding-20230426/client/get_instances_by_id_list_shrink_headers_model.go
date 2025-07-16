// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesByIdListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInstancesByIdListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetInstancesByIdListShrinkHeaders
	GetAccountContextShrink() *string
}

type GetInstancesByIdListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetInstancesByIdListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInstancesByIdListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetInstancesByIdListShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetInstancesByIdListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstancesByIdListShrinkHeaders) SetAccountContextShrink(v string) *GetInstancesByIdListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetInstancesByIdListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
