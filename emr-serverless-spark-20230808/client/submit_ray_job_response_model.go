// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRayJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitRayJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitRayJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitRayJobResponseBody) *SubmitRayJobResponse
	GetBody() *SubmitRayJobResponseBody
}

type SubmitRayJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitRayJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitRayJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitRayJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitRayJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitRayJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitRayJobResponse) GetBody() *SubmitRayJobResponseBody {
	return s.Body
}

func (s *SubmitRayJobResponse) SetHeaders(v map[string]*string) *SubmitRayJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitRayJobResponse) SetStatusCode(v int32) *SubmitRayJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitRayJobResponse) SetBody(v *SubmitRayJobResponseBody) *SubmitRayJobResponse {
	s.Body = v
	return s
}

func (s *SubmitRayJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
