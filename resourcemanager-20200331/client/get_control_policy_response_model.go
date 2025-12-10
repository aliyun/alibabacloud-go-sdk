// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetControlPolicyResponseBody) *GetControlPolicyResponse
	GetBody() *GetControlPolicyResponseBody
}

type GetControlPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetControlPolicyResponse) GetBody() *GetControlPolicyResponseBody {
	return s.Body
}

func (s *GetControlPolicyResponse) SetHeaders(v map[string]*string) *GetControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetControlPolicyResponse) SetStatusCode(v int32) *GetControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetControlPolicyResponse) SetBody(v *GetControlPolicyResponseBody) *GetControlPolicyResponse {
	s.Body = v
	return s
}

func (s *GetControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
