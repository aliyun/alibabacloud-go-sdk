// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAIProtectModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebAIProtectModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebAIProtectModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebAIProtectModeResponseBody) *ModifyWebAIProtectModeResponse
	GetBody() *ModifyWebAIProtectModeResponseBody
}

type ModifyWebAIProtectModeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebAIProtectModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebAIProtectModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAIProtectModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebAIProtectModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebAIProtectModeResponse) GetBody() *ModifyWebAIProtectModeResponseBody {
	return s.Body
}

func (s *ModifyWebAIProtectModeResponse) SetHeaders(v map[string]*string) *ModifyWebAIProtectModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAIProtectModeResponse) SetStatusCode(v int32) *ModifyWebAIProtectModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebAIProtectModeResponse) SetBody(v *ModifyWebAIProtectModeResponseBody) *ModifyWebAIProtectModeResponse {
	s.Body = v
	return s
}

func (s *ModifyWebAIProtectModeResponse) Validate() error {
	return dara.Validate(s)
}
