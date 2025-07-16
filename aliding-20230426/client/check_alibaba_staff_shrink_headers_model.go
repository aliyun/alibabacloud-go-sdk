// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAlibabaStaffShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CheckAlibabaStaffShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CheckAlibabaStaffShrinkHeaders
	GetAccountContextShrink() *string
}

type CheckAlibabaStaffShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CheckAlibabaStaffShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CheckAlibabaStaffShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CheckAlibabaStaffShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CheckAlibabaStaffShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CheckAlibabaStaffShrinkHeaders) SetCommonHeaders(v map[string]*string) *CheckAlibabaStaffShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckAlibabaStaffShrinkHeaders) SetAccountContextShrink(v string) *CheckAlibabaStaffShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CheckAlibabaStaffShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
