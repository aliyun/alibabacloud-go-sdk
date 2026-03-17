// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProbeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProbeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProbeTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListProbeTaskResponseBody) *ListProbeTaskResponse
	GetBody() *ListProbeTaskResponseBody
}

type ListProbeTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProbeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProbeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProbeTaskResponse) GoString() string {
	return s.String()
}

func (s *ListProbeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProbeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProbeTaskResponse) GetBody() *ListProbeTaskResponseBody {
	return s.Body
}

func (s *ListProbeTaskResponse) SetHeaders(v map[string]*string) *ListProbeTaskResponse {
	s.Headers = v
	return s
}

func (s *ListProbeTaskResponse) SetStatusCode(v int32) *ListProbeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProbeTaskResponse) SetBody(v *ListProbeTaskResponseBody) *ListProbeTaskResponse {
	s.Body = v
	return s
}

func (s *ListProbeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
