// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentriesInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryDentriesInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryDentriesInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryDentriesInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryDentriesInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryDentriesInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryDentriesInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryDentriesInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDentriesInfoShrinkHeaders) SetAccountContextShrink(v string) *QueryDentriesInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryDentriesInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
