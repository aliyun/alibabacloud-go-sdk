// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDetectionTaskResultHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetImageDetectionTaskResultHeaders
	GetCommonHeaders() map[string]*string
	SetXLoadTest(v bool) *GetImageDetectionTaskResultHeaders
	GetXLoadTest() *bool
}

type GetImageDetectionTaskResultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XLoadTest     *bool              `json:"X-Load-Test,omitempty" xml:"X-Load-Test,omitempty"`
}

func (s GetImageDetectionTaskResultHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultHeaders) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetImageDetectionTaskResultHeaders) GetXLoadTest() *bool {
	return s.XLoadTest
}

func (s *GetImageDetectionTaskResultHeaders) SetCommonHeaders(v map[string]*string) *GetImageDetectionTaskResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetImageDetectionTaskResultHeaders) SetXLoadTest(v bool) *GetImageDetectionTaskResultHeaders {
	s.XLoadTest = &v
	return s
}

func (s *GetImageDetectionTaskResultHeaders) Validate() error {
	return dara.Validate(s)
}
