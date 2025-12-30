// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashscopeAsyncTaskFinishEventHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DashscopeAsyncTaskFinishEventHeaders
	GetCommonHeaders() map[string]*string
	SetXLoadTest(v bool) *DashscopeAsyncTaskFinishEventHeaders
	GetXLoadTest() *bool
}

type DashscopeAsyncTaskFinishEventHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XLoadTest     *bool              `json:"X-Load-Test,omitempty" xml:"X-Load-Test,omitempty"`
}

func (s DashscopeAsyncTaskFinishEventHeaders) String() string {
	return dara.Prettify(s)
}

func (s DashscopeAsyncTaskFinishEventHeaders) GoString() string {
	return s.String()
}

func (s *DashscopeAsyncTaskFinishEventHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DashscopeAsyncTaskFinishEventHeaders) GetXLoadTest() *bool {
	return s.XLoadTest
}

func (s *DashscopeAsyncTaskFinishEventHeaders) SetCommonHeaders(v map[string]*string) *DashscopeAsyncTaskFinishEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DashscopeAsyncTaskFinishEventHeaders) SetXLoadTest(v bool) *DashscopeAsyncTaskFinishEventHeaders {
	s.XLoadTest = &v
	return s
}

func (s *DashscopeAsyncTaskFinishEventHeaders) Validate() error {
	return dara.Validate(s)
}
