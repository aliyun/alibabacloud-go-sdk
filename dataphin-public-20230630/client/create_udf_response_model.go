// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUdfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUdfResponse
	GetStatusCode() *int32
	SetBody(v *CreateUdfResponseBody) *CreateUdfResponse
	GetBody() *CreateUdfResponseBody
}

type CreateUdfResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUdfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUdfResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfResponse) GoString() string {
	return s.String()
}

func (s *CreateUdfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUdfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUdfResponse) GetBody() *CreateUdfResponseBody {
	return s.Body
}

func (s *CreateUdfResponse) SetHeaders(v map[string]*string) *CreateUdfResponse {
	s.Headers = v
	return s
}

func (s *CreateUdfResponse) SetStatusCode(v int32) *CreateUdfResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUdfResponse) SetBody(v *CreateUdfResponseBody) *CreateUdfResponse {
	s.Body = v
	return s
}

func (s *CreateUdfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
