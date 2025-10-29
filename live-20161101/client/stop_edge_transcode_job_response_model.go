// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEdgeTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopEdgeTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopEdgeTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *StopEdgeTranscodeJobResponseBody) *StopEdgeTranscodeJobResponse
	GetBody() *StopEdgeTranscodeJobResponseBody
}

type StopEdgeTranscodeJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopEdgeTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopEdgeTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopEdgeTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *StopEdgeTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopEdgeTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopEdgeTranscodeJobResponse) GetBody() *StopEdgeTranscodeJobResponseBody {
	return s.Body
}

func (s *StopEdgeTranscodeJobResponse) SetHeaders(v map[string]*string) *StopEdgeTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *StopEdgeTranscodeJobResponse) SetStatusCode(v int32) *StopEdgeTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopEdgeTranscodeJobResponse) SetBody(v *StopEdgeTranscodeJobResponseBody) *StopEdgeTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *StopEdgeTranscodeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
