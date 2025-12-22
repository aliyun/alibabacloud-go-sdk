// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCreationTaskResultHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetVideoCreationTaskResultHeaders
	GetCommonHeaders() map[string]*string
	SetXLoadTest(v bool) *GetVideoCreationTaskResultHeaders
	GetXLoadTest() *bool
}

type GetVideoCreationTaskResultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XLoadTest     *bool              `json:"X-Load-Test,omitempty" xml:"X-Load-Test,omitempty"`
}

func (s GetVideoCreationTaskResultHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultHeaders) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetVideoCreationTaskResultHeaders) GetXLoadTest() *bool {
	return s.XLoadTest
}

func (s *GetVideoCreationTaskResultHeaders) SetCommonHeaders(v map[string]*string) *GetVideoCreationTaskResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetVideoCreationTaskResultHeaders) SetXLoadTest(v bool) *GetVideoCreationTaskResultHeaders {
	s.XLoadTest = &v
	return s
}

func (s *GetVideoCreationTaskResultHeaders) Validate() error {
	return dara.Validate(s)
}
