// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSceneDefensePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSceneDefensePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSceneDefensePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateSceneDefensePolicyResponseBody) *CreateSceneDefensePolicyResponse
	GetBody() *CreateSceneDefensePolicyResponseBody
}

type CreateSceneDefensePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSceneDefensePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneDefensePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateSceneDefensePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSceneDefensePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSceneDefensePolicyResponse) GetBody() *CreateSceneDefensePolicyResponseBody {
	return s.Body
}

func (s *CreateSceneDefensePolicyResponse) SetHeaders(v map[string]*string) *CreateSceneDefensePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateSceneDefensePolicyResponse) SetStatusCode(v int32) *CreateSceneDefensePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSceneDefensePolicyResponse) SetBody(v *CreateSceneDefensePolicyResponseBody) *CreateSceneDefensePolicyResponse {
	s.Body = v
	return s
}

func (s *CreateSceneDefensePolicyResponse) Validate() error {
	return dara.Validate(s)
}
