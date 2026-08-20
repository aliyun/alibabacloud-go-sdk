// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSessionNetworkHeaderValueReplacement interface {
	dara.Model
	String() string
	GoString() string
	SetPlaceholder(v string) *SessionNetworkHeaderValueReplacement
	GetPlaceholder() *string
	SetValue(v string) *SessionNetworkHeaderValueReplacement
	GetValue() *string
}

type SessionNetworkHeaderValueReplacement struct {
	// The fake value. A placeholder used by code in the sandbox. The gateway performs an exact substring match on this string within the header value.
	//
	// example:
	//
	// sbx-notion-key-0001
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// The real value. The actual value after the placeholder is replaced.
	//
	// example:
	//
	// ntn_real_secret_xxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SessionNetworkHeaderValueReplacement) String() string {
	return dara.Prettify(s)
}

func (s SessionNetworkHeaderValueReplacement) GoString() string {
	return s.String()
}

func (s *SessionNetworkHeaderValueReplacement) GetPlaceholder() *string {
	return s.Placeholder
}

func (s *SessionNetworkHeaderValueReplacement) GetValue() *string {
	return s.Value
}

func (s *SessionNetworkHeaderValueReplacement) SetPlaceholder(v string) *SessionNetworkHeaderValueReplacement {
	s.Placeholder = &v
	return s
}

func (s *SessionNetworkHeaderValueReplacement) SetValue(v string) *SessionNetworkHeaderValueReplacement {
	s.Value = &v
	return s
}

func (s *SessionNetworkHeaderValueReplacement) Validate() error {
	return dara.Validate(s)
}
