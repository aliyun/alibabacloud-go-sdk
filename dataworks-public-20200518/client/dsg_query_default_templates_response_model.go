// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryDefaultTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgQueryDefaultTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgQueryDefaultTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DsgQueryDefaultTemplatesResponseBody) *DsgQueryDefaultTemplatesResponse
	GetBody() *DsgQueryDefaultTemplatesResponseBody
}

type DsgQueryDefaultTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgQueryDefaultTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgQueryDefaultTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDefaultTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DsgQueryDefaultTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgQueryDefaultTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgQueryDefaultTemplatesResponse) GetBody() *DsgQueryDefaultTemplatesResponseBody {
	return s.Body
}

func (s *DsgQueryDefaultTemplatesResponse) SetHeaders(v map[string]*string) *DsgQueryDefaultTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DsgQueryDefaultTemplatesResponse) SetStatusCode(v int32) *DsgQueryDefaultTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgQueryDefaultTemplatesResponse) SetBody(v *DsgQueryDefaultTemplatesResponseBody) *DsgQueryDefaultTemplatesResponse {
	s.Body = v
	return s
}

func (s *DsgQueryDefaultTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
