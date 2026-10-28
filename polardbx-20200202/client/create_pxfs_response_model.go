// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePxfsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePxfsResponse
	GetStatusCode() *int32
	SetBody(v *CreatePxfsResponseBody) *CreatePxfsResponse
	GetBody() *CreatePxfsResponseBody
}

type CreatePxfsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePxfsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePxfsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsResponse) GoString() string {
	return s.String()
}

func (s *CreatePxfsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePxfsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePxfsResponse) GetBody() *CreatePxfsResponseBody {
	return s.Body
}

func (s *CreatePxfsResponse) SetHeaders(v map[string]*string) *CreatePxfsResponse {
	s.Headers = v
	return s
}

func (s *CreatePxfsResponse) SetStatusCode(v int32) *CreatePxfsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePxfsResponse) SetBody(v *CreatePxfsResponseBody) *CreatePxfsResponse {
	s.Body = v
	return s
}

func (s *CreatePxfsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
