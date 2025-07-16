// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAttendeeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RemoveAttendeeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *RemoveAttendeeShrinkHeaders
	GetAccountContextShrink() *string
}

type RemoveAttendeeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s RemoveAttendeeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s RemoveAttendeeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RemoveAttendeeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *RemoveAttendeeShrinkHeaders) SetCommonHeaders(v map[string]*string) *RemoveAttendeeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveAttendeeShrinkHeaders) SetAccountContextShrink(v string) *RemoveAttendeeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *RemoveAttendeeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
