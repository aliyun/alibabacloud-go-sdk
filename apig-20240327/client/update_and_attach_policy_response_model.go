// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndAttachPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAndAttachPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAndAttachPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAndAttachPolicyResponseBody) *UpdateAndAttachPolicyResponse
	GetBody() *UpdateAndAttachPolicyResponseBody
}

type UpdateAndAttachPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAndAttachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAndAttachPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndAttachPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAndAttachPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAndAttachPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAndAttachPolicyResponse) GetBody() *UpdateAndAttachPolicyResponseBody {
	return s.Body
}

func (s *UpdateAndAttachPolicyResponse) SetHeaders(v map[string]*string) *UpdateAndAttachPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAndAttachPolicyResponse) SetStatusCode(v int32) *UpdateAndAttachPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAndAttachPolicyResponse) SetBody(v *UpdateAndAttachPolicyResponseBody) *UpdateAndAttachPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateAndAttachPolicyResponse) Validate() error {
	return dara.Validate(s)
}
