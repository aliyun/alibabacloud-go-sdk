// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUsageHeaders
	GetCommonHeaders() map[string]*string
	SetXLoadTest(v bool) *GetUsageHeaders
	GetXLoadTest() *bool
}

type GetUsageHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// false
	XLoadTest *bool `json:"X-Load-Test,omitempty" xml:"X-Load-Test,omitempty"`
}

func (s GetUsageHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUsageHeaders) GoString() string {
	return s.String()
}

func (s *GetUsageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUsageHeaders) GetXLoadTest() *bool {
	return s.XLoadTest
}

func (s *GetUsageHeaders) SetCommonHeaders(v map[string]*string) *GetUsageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUsageHeaders) SetXLoadTest(v bool) *GetUsageHeaders {
	s.XLoadTest = &v
	return s
}

func (s *GetUsageHeaders) Validate() error {
	return dara.Validate(s)
}
