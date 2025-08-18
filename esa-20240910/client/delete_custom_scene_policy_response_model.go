// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomScenePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomScenePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomScenePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomScenePolicyResponseBody) *DeleteCustomScenePolicyResponse
	GetBody() *DeleteCustomScenePolicyResponseBody
}

type DeleteCustomScenePolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomScenePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomScenePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomScenePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomScenePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomScenePolicyResponse) GetBody() *DeleteCustomScenePolicyResponseBody {
	return s.Body
}

func (s *DeleteCustomScenePolicyResponse) SetHeaders(v map[string]*string) *DeleteCustomScenePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomScenePolicyResponse) SetStatusCode(v int32) *DeleteCustomScenePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomScenePolicyResponse) SetBody(v *DeleteCustomScenePolicyResponseBody) *DeleteCustomScenePolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomScenePolicyResponse) Validate() error {
	return dara.Validate(s)
}
