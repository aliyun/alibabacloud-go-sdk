// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceConnectionResponseBody) *CreateServiceConnectionResponse
	GetBody() *CreateServiceConnectionResponseBody
}

type CreateServiceConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceConnectionResponse) GetBody() *CreateServiceConnectionResponseBody {
	return s.Body
}

func (s *CreateServiceConnectionResponse) SetHeaders(v map[string]*string) *CreateServiceConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceConnectionResponse) SetStatusCode(v int32) *CreateServiceConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceConnectionResponse) SetBody(v *CreateServiceConnectionResponseBody) *CreateServiceConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateServiceConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
