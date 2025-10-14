// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomScenePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomScenePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomScenePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomScenePolicyResponseBody) *CreateCustomScenePolicyResponse
	GetBody() *CreateCustomScenePolicyResponseBody
}

type CreateCustomScenePolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomScenePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomScenePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomScenePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomScenePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomScenePolicyResponse) GetBody() *CreateCustomScenePolicyResponseBody {
	return s.Body
}

func (s *CreateCustomScenePolicyResponse) SetHeaders(v map[string]*string) *CreateCustomScenePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomScenePolicyResponse) SetStatusCode(v int32) *CreateCustomScenePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomScenePolicyResponse) SetBody(v *CreateCustomScenePolicyResponseBody) *CreateCustomScenePolicyResponse {
	s.Body = v
	return s
}

func (s *CreateCustomScenePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
