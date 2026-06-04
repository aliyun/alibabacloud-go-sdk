// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomAttributeResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomAttributeResponseBody) *CreateCustomAttributeResponse
	GetBody() *CreateCustomAttributeResponseBody
}

type CreateCustomAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAttributeResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomAttributeResponse) GetBody() *CreateCustomAttributeResponseBody {
	return s.Body
}

func (s *CreateCustomAttributeResponse) SetHeaders(v map[string]*string) *CreateCustomAttributeResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomAttributeResponse) SetStatusCode(v int32) *CreateCustomAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomAttributeResponse) SetBody(v *CreateCustomAttributeResponseBody) *CreateCustomAttributeResponse {
	s.Body = v
	return s
}

func (s *CreateCustomAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
