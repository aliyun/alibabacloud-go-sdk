// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRayJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRayJobResponse
	GetStatusCode() *int32
	SetBody(v *GetRayJobResponseBody) *GetRayJobResponse
	GetBody() *GetRayJobResponseBody
}

type GetRayJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRayJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRayJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRayJobResponse) GoString() string {
	return s.String()
}

func (s *GetRayJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRayJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRayJobResponse) GetBody() *GetRayJobResponseBody {
	return s.Body
}

func (s *GetRayJobResponse) SetHeaders(v map[string]*string) *GetRayJobResponse {
	s.Headers = v
	return s
}

func (s *GetRayJobResponse) SetStatusCode(v int32) *GetRayJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRayJobResponse) SetBody(v *GetRayJobResponseBody) *GetRayJobResponse {
	s.Body = v
	return s
}

func (s *GetRayJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
