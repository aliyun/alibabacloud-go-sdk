// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCheckPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCheckPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateCheckPolicyResponseBody) *CreateCheckPolicyResponse
	GetBody() *CreateCheckPolicyResponseBody
}

type CreateCheckPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCheckPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCheckPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateCheckPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCheckPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCheckPolicyResponse) GetBody() *CreateCheckPolicyResponseBody {
	return s.Body
}

func (s *CreateCheckPolicyResponse) SetHeaders(v map[string]*string) *CreateCheckPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateCheckPolicyResponse) SetStatusCode(v int32) *CreateCheckPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCheckPolicyResponse) SetBody(v *CreateCheckPolicyResponseBody) *CreateCheckPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateCheckPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
