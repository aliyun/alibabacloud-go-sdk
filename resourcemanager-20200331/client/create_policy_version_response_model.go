// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolicyVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolicyVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolicyVersionResponseBody) *CreatePolicyVersionResponse
	GetBody() *CreatePolicyVersionResponseBody
}

type CreatePolicyVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolicyVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolicyVersionResponse) GetBody() *CreatePolicyVersionResponseBody {
	return s.Body
}

func (s *CreatePolicyVersionResponse) SetHeaders(v map[string]*string) *CreatePolicyVersionResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyVersionResponse) SetStatusCode(v int32) *CreatePolicyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyVersionResponse) SetBody(v *CreatePolicyVersionResponseBody) *CreatePolicyVersionResponse {
	s.Body = v
	return s
}

func (s *CreatePolicyVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
