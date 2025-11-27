// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrePayInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrePayInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrePayInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrePayInstanceResponseBody) *CreatePrePayInstanceResponse
	GetBody() *CreatePrePayInstanceResponseBody
}

type CreatePrePayInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrePayInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrePayInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrePayInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrePayInstanceResponse) GetBody() *CreatePrePayInstanceResponseBody {
	return s.Body
}

func (s *CreatePrePayInstanceResponse) SetHeaders(v map[string]*string) *CreatePrePayInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePrePayInstanceResponse) SetStatusCode(v int32) *CreatePrePayInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrePayInstanceResponse) SetBody(v *CreatePrePayInstanceResponseBody) *CreatePrePayInstanceResponse {
	s.Body = v
	return s
}

func (s *CreatePrePayInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
