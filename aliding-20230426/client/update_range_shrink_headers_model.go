// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRangeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateRangeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateRangeShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateRangeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateRangeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateRangeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRangeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateRangeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateRangeShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateRangeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRangeShrinkHeaders) SetAccountContextShrink(v string) *UpdateRangeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateRangeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
