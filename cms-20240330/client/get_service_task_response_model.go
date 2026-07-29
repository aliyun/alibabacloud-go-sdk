// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceTaskResponseBody) *GetServiceTaskResponse
	GetBody() *GetServiceTaskResponseBody
}

type GetServiceTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTaskResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceTaskResponse) GetBody() *GetServiceTaskResponseBody {
	return s.Body
}

func (s *GetServiceTaskResponse) SetHeaders(v map[string]*string) *GetServiceTaskResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTaskResponse) SetStatusCode(v int32) *GetServiceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTaskResponse) SetBody(v *GetServiceTaskResponseBody) *GetServiceTaskResponse {
	s.Body = v
	return s
}

func (s *GetServiceTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
