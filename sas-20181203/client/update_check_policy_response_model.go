// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCheckPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCheckPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCheckPolicyResponseBody) *UpdateCheckPolicyResponse
	GetBody() *UpdateCheckPolicyResponseBody
}

type UpdateCheckPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCheckPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCheckPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateCheckPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCheckPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCheckPolicyResponse) GetBody() *UpdateCheckPolicyResponseBody {
	return s.Body
}

func (s *UpdateCheckPolicyResponse) SetHeaders(v map[string]*string) *UpdateCheckPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateCheckPolicyResponse) SetStatusCode(v int32) *UpdateCheckPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCheckPolicyResponse) SetBody(v *UpdateCheckPolicyResponseBody) *UpdateCheckPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateCheckPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
