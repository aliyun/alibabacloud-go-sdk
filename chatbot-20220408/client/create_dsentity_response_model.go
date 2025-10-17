// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDSEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDSEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDSEntityResponse
	GetStatusCode() *int32
	SetBody(v *CreateDSEntityResponseBody) *CreateDSEntityResponse
	GetBody() *CreateDSEntityResponseBody
}

type CreateDSEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDSEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDSEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateDSEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDSEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDSEntityResponse) GetBody() *CreateDSEntityResponseBody {
	return s.Body
}

func (s *CreateDSEntityResponse) SetHeaders(v map[string]*string) *CreateDSEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateDSEntityResponse) SetStatusCode(v int32) *CreateDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDSEntityResponse) SetBody(v *CreateDSEntityResponseBody) *CreateDSEntityResponse {
	s.Body = v
	return s
}

func (s *CreateDSEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
