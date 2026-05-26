// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicySetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolicySetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolicySetResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolicySetResponseBody) *CreatePolicySetResponse
	GetBody() *CreatePolicySetResponseBody
}

type CreatePolicySetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicySetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicySetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicySetResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicySetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolicySetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolicySetResponse) GetBody() *CreatePolicySetResponseBody {
	return s.Body
}

func (s *CreatePolicySetResponse) SetHeaders(v map[string]*string) *CreatePolicySetResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicySetResponse) SetStatusCode(v int32) *CreatePolicySetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicySetResponse) SetBody(v *CreatePolicySetResponseBody) *CreatePolicySetResponse {
	s.Body = v
	return s
}

func (s *CreatePolicySetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
