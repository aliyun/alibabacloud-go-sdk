// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccountZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccountZonalResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccountZonalResponseBody) *CreateAccountZonalResponse
	GetBody() *CreateAccountZonalResponseBody
}

type CreateAccountZonalResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccountZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountZonalResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccountZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccountZonalResponse) GetBody() *CreateAccountZonalResponseBody {
	return s.Body
}

func (s *CreateAccountZonalResponse) SetHeaders(v map[string]*string) *CreateAccountZonalResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountZonalResponse) SetStatusCode(v int32) *CreateAccountZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountZonalResponse) SetBody(v *CreateAccountZonalResponseBody) *CreateAccountZonalResponse {
	s.Body = v
	return s
}

func (s *CreateAccountZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
