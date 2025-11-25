// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSceneDefensePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSceneDefensePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSceneDefensePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DisableSceneDefensePolicyResponseBody) *DisableSceneDefensePolicyResponse
	GetBody() *DisableSceneDefensePolicyResponseBody
}

type DisableSceneDefensePolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSceneDefensePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneDefensePolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableSceneDefensePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSceneDefensePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSceneDefensePolicyResponse) GetBody() *DisableSceneDefensePolicyResponseBody {
	return s.Body
}

func (s *DisableSceneDefensePolicyResponse) SetHeaders(v map[string]*string) *DisableSceneDefensePolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableSceneDefensePolicyResponse) SetStatusCode(v int32) *DisableSceneDefensePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSceneDefensePolicyResponse) SetBody(v *DisableSceneDefensePolicyResponseBody) *DisableSceneDefensePolicyResponse {
	s.Body = v
	return s
}

func (s *DisableSceneDefensePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
