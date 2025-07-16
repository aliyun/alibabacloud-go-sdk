// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchEventShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PatchEventShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *PatchEventShrinkHeaders
	GetAccountContextShrink() *string
}

type PatchEventShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s PatchEventShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s PatchEventShrinkHeaders) GoString() string {
	return s.String()
}

func (s *PatchEventShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PatchEventShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *PatchEventShrinkHeaders) SetCommonHeaders(v map[string]*string) *PatchEventShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchEventShrinkHeaders) SetAccountContextShrink(v string) *PatchEventShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *PatchEventShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
