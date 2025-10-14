// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomScenePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCustomScenePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCustomScenePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DisableCustomScenePolicyResponseBody) *DisableCustomScenePolicyResponse
	GetBody() *DisableCustomScenePolicyResponseBody
}

type DisableCustomScenePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCustomScenePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomScenePolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableCustomScenePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCustomScenePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCustomScenePolicyResponse) GetBody() *DisableCustomScenePolicyResponseBody {
	return s.Body
}

func (s *DisableCustomScenePolicyResponse) SetHeaders(v map[string]*string) *DisableCustomScenePolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableCustomScenePolicyResponse) SetStatusCode(v int32) *DisableCustomScenePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCustomScenePolicyResponse) SetBody(v *DisableCustomScenePolicyResponseBody) *DisableCustomScenePolicyResponse {
	s.Body = v
	return s
}

func (s *DisableCustomScenePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
