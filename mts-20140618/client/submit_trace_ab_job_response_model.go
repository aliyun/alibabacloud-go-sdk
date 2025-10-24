// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceAbJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTraceAbJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTraceAbJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTraceAbJobResponseBody) *SubmitTraceAbJobResponse
	GetBody() *SubmitTraceAbJobResponseBody
}

type SubmitTraceAbJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTraceAbJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTraceAbJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceAbJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTraceAbJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTraceAbJobResponse) GetBody() *SubmitTraceAbJobResponseBody {
	return s.Body
}

func (s *SubmitTraceAbJobResponse) SetHeaders(v map[string]*string) *SubmitTraceAbJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitTraceAbJobResponse) SetStatusCode(v int32) *SubmitTraceAbJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTraceAbJobResponse) SetBody(v *SubmitTraceAbJobResponseBody) *SubmitTraceAbJobResponse {
	s.Body = v
	return s
}

func (s *SubmitTraceAbJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
