// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseVideoConferenceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CloseVideoConferenceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CloseVideoConferenceShrinkHeaders
	GetAccountContextShrink() *string
}

type CloseVideoConferenceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CloseVideoConferenceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CloseVideoConferenceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CloseVideoConferenceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CloseVideoConferenceShrinkHeaders) SetCommonHeaders(v map[string]*string) *CloseVideoConferenceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseVideoConferenceShrinkHeaders) SetAccountContextShrink(v string) *CloseVideoConferenceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CloseVideoConferenceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
