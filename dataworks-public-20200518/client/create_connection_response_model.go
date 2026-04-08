// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateConnectionResponseBody) *CreateConnectionResponse
	GetBody() *CreateConnectionResponseBody
}

type CreateConnectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConnectionResponse) GetBody() *CreateConnectionResponseBody {
	return s.Body
}

func (s *CreateConnectionResponse) SetHeaders(v map[string]*string) *CreateConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateConnectionResponse) SetStatusCode(v int32) *CreateConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConnectionResponse) SetBody(v *CreateConnectionResponseBody) *CreateConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
