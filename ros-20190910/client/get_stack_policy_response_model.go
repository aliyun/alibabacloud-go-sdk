// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetStackPolicyResponseBody) *GetStackPolicyResponse
	GetBody() *GetStackPolicyResponseBody
}

type GetStackPolicyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetStackPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackPolicyResponse) GetBody() *GetStackPolicyResponseBody {
	return s.Body
}

func (s *GetStackPolicyResponse) SetHeaders(v map[string]*string) *GetStackPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetStackPolicyResponse) SetStatusCode(v int32) *GetStackPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackPolicyResponse) SetBody(v *GetStackPolicyResponseBody) *GetStackPolicyResponse {
	s.Body = v
	return s
}

func (s *GetStackPolicyResponse) Validate() error {
	return dara.Validate(s)
}
