// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRayJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelRayJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelRayJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelRayJobResponseBody) *CancelRayJobResponse
	GetBody() *CancelRayJobResponseBody
}

type CancelRayJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRayJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRayJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelRayJobResponse) GoString() string {
	return s.String()
}

func (s *CancelRayJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelRayJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelRayJobResponse) GetBody() *CancelRayJobResponseBody {
	return s.Body
}

func (s *CancelRayJobResponse) SetHeaders(v map[string]*string) *CancelRayJobResponse {
	s.Headers = v
	return s
}

func (s *CancelRayJobResponse) SetStatusCode(v int32) *CancelRayJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRayJobResponse) SetBody(v *CancelRayJobResponseBody) *CancelRayJobResponse {
	s.Body = v
	return s
}

func (s *CancelRayJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
