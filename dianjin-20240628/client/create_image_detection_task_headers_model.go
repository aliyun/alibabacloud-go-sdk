// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageDetectionTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateImageDetectionTaskHeaders
	GetCommonHeaders() map[string]*string
	SetXLoadTest(v bool) *CreateImageDetectionTaskHeaders
	GetXLoadTest() *bool
}

type CreateImageDetectionTaskHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XLoadTest     *bool              `json:"X-Load-Test,omitempty" xml:"X-Load-Test,omitempty"`
}

func (s CreateImageDetectionTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDetectionTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateImageDetectionTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateImageDetectionTaskHeaders) GetXLoadTest() *bool {
	return s.XLoadTest
}

func (s *CreateImageDetectionTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateImageDetectionTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateImageDetectionTaskHeaders) SetXLoadTest(v bool) *CreateImageDetectionTaskHeaders {
	s.XLoadTest = &v
	return s
}

func (s *CreateImageDetectionTaskHeaders) Validate() error {
	return dara.Validate(s)
}
