// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProhibitedPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProhibitedPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProhibitedPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateProhibitedPolicyResponseBody) *CreateProhibitedPolicyResponse
	GetBody() *CreateProhibitedPolicyResponseBody
}

type CreateProhibitedPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProhibitedPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProhibitedPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateProhibitedPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProhibitedPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProhibitedPolicyResponse) GetBody() *CreateProhibitedPolicyResponseBody {
	return s.Body
}

func (s *CreateProhibitedPolicyResponse) SetHeaders(v map[string]*string) *CreateProhibitedPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateProhibitedPolicyResponse) SetStatusCode(v int32) *CreateProhibitedPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProhibitedPolicyResponse) SetBody(v *CreateProhibitedPolicyResponseBody) *CreateProhibitedPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateProhibitedPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
