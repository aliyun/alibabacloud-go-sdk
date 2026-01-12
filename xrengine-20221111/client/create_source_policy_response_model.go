// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourcePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSourcePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSourcePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateSourcePolicyResponseBody) *CreateSourcePolicyResponse
	GetBody() *CreateSourcePolicyResponseBody
}

type CreateSourcePolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSourcePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSourcePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSourcePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateSourcePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSourcePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSourcePolicyResponse) GetBody() *CreateSourcePolicyResponseBody {
	return s.Body
}

func (s *CreateSourcePolicyResponse) SetHeaders(v map[string]*string) *CreateSourcePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateSourcePolicyResponse) SetStatusCode(v int32) *CreateSourcePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSourcePolicyResponse) SetBody(v *CreateSourcePolicyResponseBody) *CreateSourcePolicyResponse {
	s.Body = v
	return s
}

func (s *CreateSourcePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
