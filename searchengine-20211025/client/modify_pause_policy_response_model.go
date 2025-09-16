// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPausePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPausePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPausePolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPausePolicyResponseBody) *ModifyPausePolicyResponse
	GetBody() *ModifyPausePolicyResponseBody
}

type ModifyPausePolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPausePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPausePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPausePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyPausePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPausePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPausePolicyResponse) GetBody() *ModifyPausePolicyResponseBody {
	return s.Body
}

func (s *ModifyPausePolicyResponse) SetHeaders(v map[string]*string) *ModifyPausePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyPausePolicyResponse) SetStatusCode(v int32) *ModifyPausePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPausePolicyResponse) SetBody(v *ModifyPausePolicyResponseBody) *ModifyPausePolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyPausePolicyResponse) Validate() error {
	return dara.Validate(s)
}
