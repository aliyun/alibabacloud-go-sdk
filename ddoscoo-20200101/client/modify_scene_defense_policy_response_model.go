// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySceneDefensePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySceneDefensePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySceneDefensePolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifySceneDefensePolicyResponseBody) *ModifySceneDefensePolicyResponse
	GetBody() *ModifySceneDefensePolicyResponseBody
}

type ModifySceneDefensePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySceneDefensePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySceneDefensePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifySceneDefensePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySceneDefensePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySceneDefensePolicyResponse) GetBody() *ModifySceneDefensePolicyResponseBody {
	return s.Body
}

func (s *ModifySceneDefensePolicyResponse) SetHeaders(v map[string]*string) *ModifySceneDefensePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifySceneDefensePolicyResponse) SetStatusCode(v int32) *ModifySceneDefensePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySceneDefensePolicyResponse) SetBody(v *ModifySceneDefensePolicyResponseBody) *ModifySceneDefensePolicyResponse {
	s.Body = v
	return s
}

func (s *ModifySceneDefensePolicyResponse) Validate() error {
	return dara.Validate(s)
}
