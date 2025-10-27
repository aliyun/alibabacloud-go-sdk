// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoginSwitchConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLoginSwitchConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLoginSwitchConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLoginSwitchConfigResponseBody) *ModifyLoginSwitchConfigResponse
	GetBody() *ModifyLoginSwitchConfigResponseBody
}

type ModifyLoginSwitchConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLoginSwitchConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLoginSwitchConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoginSwitchConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoginSwitchConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLoginSwitchConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLoginSwitchConfigResponse) GetBody() *ModifyLoginSwitchConfigResponseBody {
	return s.Body
}

func (s *ModifyLoginSwitchConfigResponse) SetHeaders(v map[string]*string) *ModifyLoginSwitchConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoginSwitchConfigResponse) SetStatusCode(v int32) *ModifyLoginSwitchConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoginSwitchConfigResponse) SetBody(v *ModifyLoginSwitchConfigResponseBody) *ModifyLoginSwitchConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyLoginSwitchConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
