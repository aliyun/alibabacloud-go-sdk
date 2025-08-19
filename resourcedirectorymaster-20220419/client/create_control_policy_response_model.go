// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateControlPolicyResponseBody) *CreateControlPolicyResponse
	GetBody() *CreateControlPolicyResponseBody
}

type CreateControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateControlPolicyResponse) GetBody() *CreateControlPolicyResponseBody {
	return s.Body
}

func (s *CreateControlPolicyResponse) SetHeaders(v map[string]*string) *CreateControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateControlPolicyResponse) SetStatusCode(v int32) *CreateControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateControlPolicyResponse) SetBody(v *CreateControlPolicyResponseBody) *CreateControlPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateControlPolicyResponse) Validate() error {
	return dara.Validate(s)
}
