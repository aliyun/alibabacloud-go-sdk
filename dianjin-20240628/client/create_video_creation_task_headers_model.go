// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoCreationTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateVideoCreationTaskHeaders
	GetCommonHeaders() map[string]*string
	SetXLoadTest(v bool) *CreateVideoCreationTaskHeaders
	GetXLoadTest() *bool
}

type CreateVideoCreationTaskHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XLoadTest     *bool              `json:"X-Load-Test,omitempty" xml:"X-Load-Test,omitempty"`
}

func (s CreateVideoCreationTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoCreationTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateVideoCreationTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateVideoCreationTaskHeaders) GetXLoadTest() *bool {
	return s.XLoadTest
}

func (s *CreateVideoCreationTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateVideoCreationTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateVideoCreationTaskHeaders) SetXLoadTest(v bool) *CreateVideoCreationTaskHeaders {
	s.XLoadTest = &v
	return s
}

func (s *CreateVideoCreationTaskHeaders) Validate() error {
	return dara.Validate(s)
}
