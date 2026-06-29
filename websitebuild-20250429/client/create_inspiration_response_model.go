// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspirationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInspirationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInspirationResponse
	GetStatusCode() *int32
	SetBody(v *CreateInspirationResponseBody) *CreateInspirationResponse
	GetBody() *CreateInspirationResponseBody
}

type CreateInspirationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInspirationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInspirationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInspirationResponse) GoString() string {
	return s.String()
}

func (s *CreateInspirationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInspirationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInspirationResponse) GetBody() *CreateInspirationResponseBody {
	return s.Body
}

func (s *CreateInspirationResponse) SetHeaders(v map[string]*string) *CreateInspirationResponse {
	s.Headers = v
	return s
}

func (s *CreateInspirationResponse) SetStatusCode(v int32) *CreateInspirationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInspirationResponse) SetBody(v *CreateInspirationResponseBody) *CreateInspirationResponse {
	s.Body = v
	return s
}

func (s *CreateInspirationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
