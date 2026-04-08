// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusinessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBusinessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBusinessResponse
	GetStatusCode() *int32
	SetBody(v *GetBusinessResponseBody) *GetBusinessResponse
	GetBody() *GetBusinessResponseBody
}

type GetBusinessResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBusinessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBusinessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessResponse) GoString() string {
	return s.String()
}

func (s *GetBusinessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBusinessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBusinessResponse) GetBody() *GetBusinessResponseBody {
	return s.Body
}

func (s *GetBusinessResponse) SetHeaders(v map[string]*string) *GetBusinessResponse {
	s.Headers = v
	return s
}

func (s *GetBusinessResponse) SetStatusCode(v int32) *GetBusinessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBusinessResponse) SetBody(v *GetBusinessResponseBody) *GetBusinessResponse {
	s.Body = v
	return s
}

func (s *GetBusinessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
