// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePxfsUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePxfsUserResponse
	GetStatusCode() *int32
	SetBody(v *CreatePxfsUserResponseBody) *CreatePxfsUserResponse
	GetBody() *CreatePxfsUserResponseBody
}

type CreatePxfsUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePxfsUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePxfsUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsUserResponse) GoString() string {
	return s.String()
}

func (s *CreatePxfsUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePxfsUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePxfsUserResponse) GetBody() *CreatePxfsUserResponseBody {
	return s.Body
}

func (s *CreatePxfsUserResponse) SetHeaders(v map[string]*string) *CreatePxfsUserResponse {
	s.Headers = v
	return s
}

func (s *CreatePxfsUserResponse) SetStatusCode(v int32) *CreatePxfsUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePxfsUserResponse) SetBody(v *CreatePxfsUserResponseBody) *CreatePxfsUserResponse {
	s.Body = v
	return s
}

func (s *CreatePxfsUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
