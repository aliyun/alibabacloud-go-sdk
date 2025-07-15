// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenterPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCenterPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCenterPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateCenterPolicyResponseBody) *CreateCenterPolicyResponse
	GetBody() *CreateCenterPolicyResponseBody
}

type CreateCenterPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCenterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCenterPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCenterPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCenterPolicyResponse) GetBody() *CreateCenterPolicyResponseBody {
	return s.Body
}

func (s *CreateCenterPolicyResponse) SetHeaders(v map[string]*string) *CreateCenterPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateCenterPolicyResponse) SetStatusCode(v int32) *CreateCenterPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCenterPolicyResponse) SetBody(v *CreateCenterPolicyResponseBody) *CreateCenterPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateCenterPolicyResponse) Validate() error {
	return dara.Validate(s)
}
