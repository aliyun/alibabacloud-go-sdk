// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomScenePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomScenePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomScenePolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomScenePolicyResponseBody) *UpdateCustomScenePolicyResponse
	GetBody() *UpdateCustomScenePolicyResponseBody
}

type UpdateCustomScenePolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomScenePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomScenePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomScenePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomScenePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomScenePolicyResponse) GetBody() *UpdateCustomScenePolicyResponseBody {
	return s.Body
}

func (s *UpdateCustomScenePolicyResponse) SetHeaders(v map[string]*string) *UpdateCustomScenePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomScenePolicyResponse) SetStatusCode(v int32) *UpdateCustomScenePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomScenePolicyResponse) SetBody(v *UpdateCustomScenePolicyResponseBody) *UpdateCustomScenePolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomScenePolicyResponse) Validate() error {
	return dara.Validate(s)
}
