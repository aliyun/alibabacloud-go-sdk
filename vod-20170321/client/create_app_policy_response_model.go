// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppPolicyResponseBody) *CreateAppPolicyResponse
	GetBody() *CreateAppPolicyResponseBody
}

type CreateAppPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAppPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppPolicyResponse) GetBody() *CreateAppPolicyResponseBody {
	return s.Body
}

func (s *CreateAppPolicyResponse) SetHeaders(v map[string]*string) *CreateAppPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAppPolicyResponse) SetStatusCode(v int32) *CreateAppPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppPolicyResponse) SetBody(v *CreateAppPolicyResponseBody) *CreateAppPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateAppPolicyResponse) Validate() error {
	return dara.Validate(s)
}
