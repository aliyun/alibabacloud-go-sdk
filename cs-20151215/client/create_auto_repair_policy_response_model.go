// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoRepairPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAutoRepairPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAutoRepairPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAutoRepairPolicyResponseBody) *CreateAutoRepairPolicyResponse
	GetBody() *CreateAutoRepairPolicyResponseBody
}

type CreateAutoRepairPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoRepairPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoRepairPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAutoRepairPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAutoRepairPolicyResponse) GetBody() *CreateAutoRepairPolicyResponseBody {
	return s.Body
}

func (s *CreateAutoRepairPolicyResponse) SetHeaders(v map[string]*string) *CreateAutoRepairPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoRepairPolicyResponse) SetStatusCode(v int32) *CreateAutoRepairPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoRepairPolicyResponse) SetBody(v *CreateAutoRepairPolicyResponseBody) *CreateAutoRepairPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateAutoRepairPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
