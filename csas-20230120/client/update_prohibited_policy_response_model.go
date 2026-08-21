// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProhibitedPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProhibitedPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProhibitedPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProhibitedPolicyResponseBody) *UpdateProhibitedPolicyResponse
	GetBody() *UpdateProhibitedPolicyResponseBody
}

type UpdateProhibitedPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProhibitedPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProhibitedPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProhibitedPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProhibitedPolicyResponse) GetBody() *UpdateProhibitedPolicyResponseBody {
	return s.Body
}

func (s *UpdateProhibitedPolicyResponse) SetHeaders(v map[string]*string) *UpdateProhibitedPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateProhibitedPolicyResponse) SetStatusCode(v int32) *UpdateProhibitedPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProhibitedPolicyResponse) SetBody(v *UpdateProhibitedPolicyResponseBody) *UpdateProhibitedPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateProhibitedPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
