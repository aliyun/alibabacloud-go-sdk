// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIngressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIngressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIngressResponse
	GetStatusCode() *int32
	SetBody(v *CreateIngressResponseBody) *CreateIngressResponse
	GetBody() *CreateIngressResponseBody
}

type CreateIngressResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIngressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIngressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIngressResponse) GoString() string {
	return s.String()
}

func (s *CreateIngressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIngressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIngressResponse) GetBody() *CreateIngressResponseBody {
	return s.Body
}

func (s *CreateIngressResponse) SetHeaders(v map[string]*string) *CreateIngressResponse {
	s.Headers = v
	return s
}

func (s *CreateIngressResponse) SetStatusCode(v int32) *CreateIngressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIngressResponse) SetBody(v *CreateIngressResponseBody) *CreateIngressResponse {
	s.Body = v
	return s
}

func (s *CreateIngressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
