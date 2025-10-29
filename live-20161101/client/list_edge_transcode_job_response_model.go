// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeTranscodeJobResponseBody) *ListEdgeTranscodeJobResponse
	GetBody() *ListEdgeTranscodeJobResponseBody
}

type ListEdgeTranscodeJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeTranscodeJobResponse) GetBody() *ListEdgeTranscodeJobResponseBody {
	return s.Body
}

func (s *ListEdgeTranscodeJobResponse) SetHeaders(v map[string]*string) *ListEdgeTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeTranscodeJobResponse) SetStatusCode(v int32) *ListEdgeTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeTranscodeJobResponse) SetBody(v *ListEdgeTranscodeJobResponseBody) *ListEdgeTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *ListEdgeTranscodeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
