// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePreCheckResponse
	GetStatusCode() *int32
	SetBody(v *CreatePreCheckResponseBody) *CreatePreCheckResponse
	GetBody() *CreatePreCheckResponseBody
}

type CreatePreCheckResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePreCheckResponse) GoString() string {
	return s.String()
}

func (s *CreatePreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePreCheckResponse) GetBody() *CreatePreCheckResponseBody {
	return s.Body
}

func (s *CreatePreCheckResponse) SetHeaders(v map[string]*string) *CreatePreCheckResponse {
	s.Headers = v
	return s
}

func (s *CreatePreCheckResponse) SetStatusCode(v int32) *CreatePreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePreCheckResponse) SetBody(v *CreatePreCheckResponseBody) *CreatePreCheckResponse {
	s.Body = v
	return s
}

func (s *CreatePreCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
