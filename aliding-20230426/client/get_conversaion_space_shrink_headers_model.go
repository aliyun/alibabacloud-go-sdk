// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversaionSpaceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetConversaionSpaceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetConversaionSpaceShrinkHeaders
	GetAccountContextShrink() *string
}

type GetConversaionSpaceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetConversaionSpaceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetConversaionSpaceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetConversaionSpaceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetConversaionSpaceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetConversaionSpaceShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetConversaionSpaceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversaionSpaceShrinkHeaders) SetAccountContextShrink(v string) *GetConversaionSpaceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetConversaionSpaceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
