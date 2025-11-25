// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAIProtectSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebAIProtectSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebAIProtectSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebAIProtectSwitchResponseBody) *ModifyWebAIProtectSwitchResponse
	GetBody() *ModifyWebAIProtectSwitchResponseBody
}

type ModifyWebAIProtectSwitchResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebAIProtectSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebAIProtectSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAIProtectSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebAIProtectSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebAIProtectSwitchResponse) GetBody() *ModifyWebAIProtectSwitchResponseBody {
	return s.Body
}

func (s *ModifyWebAIProtectSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebAIProtectSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAIProtectSwitchResponse) SetStatusCode(v int32) *ModifyWebAIProtectSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebAIProtectSwitchResponse) SetBody(v *ModifyWebAIProtectSwitchResponseBody) *ModifyWebAIProtectSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyWebAIProtectSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
