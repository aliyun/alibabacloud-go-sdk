// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceJobResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceJobResponseBody) *GetInstanceJobResponse
	GetBody() *GetInstanceJobResponseBody
}

type GetInstanceJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceJobResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceJobResponse) GetBody() *GetInstanceJobResponseBody {
	return s.Body
}

func (s *GetInstanceJobResponse) SetHeaders(v map[string]*string) *GetInstanceJobResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceJobResponse) SetStatusCode(v int32) *GetInstanceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceJobResponse) SetBody(v *GetInstanceJobResponseBody) *GetInstanceJobResponse {
	s.Body = v
	return s
}

func (s *GetInstanceJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
