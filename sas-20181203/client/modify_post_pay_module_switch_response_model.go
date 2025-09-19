// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPostPayModuleSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPostPayModuleSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPostPayModuleSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPostPayModuleSwitchResponseBody) *ModifyPostPayModuleSwitchResponse
	GetBody() *ModifyPostPayModuleSwitchResponseBody
}

type ModifyPostPayModuleSwitchResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPostPayModuleSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPostPayModuleSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPostPayModuleSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyPostPayModuleSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPostPayModuleSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPostPayModuleSwitchResponse) GetBody() *ModifyPostPayModuleSwitchResponseBody {
	return s.Body
}

func (s *ModifyPostPayModuleSwitchResponse) SetHeaders(v map[string]*string) *ModifyPostPayModuleSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyPostPayModuleSwitchResponse) SetStatusCode(v int32) *ModifyPostPayModuleSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPostPayModuleSwitchResponse) SetBody(v *ModifyPostPayModuleSwitchResponseBody) *ModifyPostPayModuleSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyPostPayModuleSwitchResponse) Validate() error {
	return dara.Validate(s)
}
