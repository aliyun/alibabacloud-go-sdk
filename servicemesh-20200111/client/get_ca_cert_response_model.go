// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCaCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCaCertResponse
	GetStatusCode() *int32
	SetBody(v *GetCaCertResponseBody) *GetCaCertResponse
	GetBody() *GetCaCertResponseBody
}

type GetCaCertResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCaCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCaCertResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCaCertResponse) GoString() string {
	return s.String()
}

func (s *GetCaCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCaCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCaCertResponse) GetBody() *GetCaCertResponseBody {
	return s.Body
}

func (s *GetCaCertResponse) SetHeaders(v map[string]*string) *GetCaCertResponse {
	s.Headers = v
	return s
}

func (s *GetCaCertResponse) SetStatusCode(v int32) *GetCaCertResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCaCertResponse) SetBody(v *GetCaCertResponseBody) *GetCaCertResponse {
	s.Body = v
	return s
}

func (s *GetCaCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
