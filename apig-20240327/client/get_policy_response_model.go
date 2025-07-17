// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetPolicyResponseBody) *GetPolicyResponse
	GetBody() *GetPolicyResponseBody
}

type GetPolicyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolicyResponse) GetBody() *GetPolicyResponseBody {
	return s.Body
}

func (s *GetPolicyResponse) SetHeaders(v map[string]*string) *GetPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyResponse) SetStatusCode(v int32) *GetPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyResponse) SetBody(v *GetPolicyResponseBody) *GetPolicyResponse {
	s.Body = v
	return s
}

func (s *GetPolicyResponse) Validate() error {
	return dara.Validate(s)
}
