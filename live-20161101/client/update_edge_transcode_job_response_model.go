// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEdgeTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEdgeTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEdgeTranscodeJobResponseBody) *UpdateEdgeTranscodeJobResponse
	GetBody() *UpdateEdgeTranscodeJobResponseBody
}

type UpdateEdgeTranscodeJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEdgeTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEdgeTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateEdgeTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEdgeTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEdgeTranscodeJobResponse) GetBody() *UpdateEdgeTranscodeJobResponseBody {
	return s.Body
}

func (s *UpdateEdgeTranscodeJobResponse) SetHeaders(v map[string]*string) *UpdateEdgeTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateEdgeTranscodeJobResponse) SetStatusCode(v int32) *UpdateEdgeTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEdgeTranscodeJobResponse) SetBody(v *UpdateEdgeTranscodeJobResponseBody) *UpdateEdgeTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *UpdateEdgeTranscodeJobResponse) Validate() error {
	return dara.Validate(s)
}
