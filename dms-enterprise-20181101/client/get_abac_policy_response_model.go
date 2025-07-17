// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAbacPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAbacPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAbacPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetAbacPolicyResponseBody) *GetAbacPolicyResponse
	GetBody() *GetAbacPolicyResponseBody
}

type GetAbacPolicyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAbacPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAbacPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAbacPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetAbacPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAbacPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAbacPolicyResponse) GetBody() *GetAbacPolicyResponseBody {
	return s.Body
}

func (s *GetAbacPolicyResponse) SetHeaders(v map[string]*string) *GetAbacPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetAbacPolicyResponse) SetStatusCode(v int32) *GetAbacPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAbacPolicyResponse) SetBody(v *GetAbacPolicyResponseBody) *GetAbacPolicyResponse {
	s.Body = v
	return s
}

func (s *GetAbacPolicyResponse) Validate() error {
	return dara.Validate(s)
}
