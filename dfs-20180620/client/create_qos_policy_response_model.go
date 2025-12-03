// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQosPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQosPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateQosPolicyResponseBody) *CreateQosPolicyResponse
	GetBody() *CreateQosPolicyResponseBody
}

type CreateQosPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQosPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQosPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQosPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateQosPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQosPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQosPolicyResponse) GetBody() *CreateQosPolicyResponseBody {
	return s.Body
}

func (s *CreateQosPolicyResponse) SetHeaders(v map[string]*string) *CreateQosPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateQosPolicyResponse) SetStatusCode(v int32) *CreateQosPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQosPolicyResponse) SetBody(v *CreateQosPolicyResponseBody) *CreateQosPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateQosPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
