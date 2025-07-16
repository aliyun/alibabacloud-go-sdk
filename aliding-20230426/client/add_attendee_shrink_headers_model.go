// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAttendeeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddAttendeeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddAttendeeShrinkHeaders
	GetAccountContextShrink() *string
}

type AddAttendeeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddAttendeeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddAttendeeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddAttendeeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddAttendeeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddAttendeeShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddAttendeeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAttendeeShrinkHeaders) SetAccountContextShrink(v string) *AddAttendeeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddAttendeeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
