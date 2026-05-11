// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGreetingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGreetingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGreetingConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGreetingConfigResponseBody) *ModifyGreetingConfigResponse
	GetBody() *ModifyGreetingConfigResponseBody
}

type ModifyGreetingConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGreetingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGreetingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGreetingConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyGreetingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGreetingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGreetingConfigResponse) GetBody() *ModifyGreetingConfigResponseBody {
	return s.Body
}

func (s *ModifyGreetingConfigResponse) SetHeaders(v map[string]*string) *ModifyGreetingConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyGreetingConfigResponse) SetStatusCode(v int32) *ModifyGreetingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGreetingConfigResponse) SetBody(v *ModifyGreetingConfigResponseBody) *ModifyGreetingConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyGreetingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
