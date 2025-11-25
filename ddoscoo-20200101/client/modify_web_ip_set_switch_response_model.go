// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebIpSetSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebIpSetSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebIpSetSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebIpSetSwitchResponseBody) *ModifyWebIpSetSwitchResponse
	GetBody() *ModifyWebIpSetSwitchResponseBody
}

type ModifyWebIpSetSwitchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebIpSetSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebIpSetSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebIpSetSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebIpSetSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebIpSetSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebIpSetSwitchResponse) GetBody() *ModifyWebIpSetSwitchResponseBody {
	return s.Body
}

func (s *ModifyWebIpSetSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebIpSetSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebIpSetSwitchResponse) SetStatusCode(v int32) *ModifyWebIpSetSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebIpSetSwitchResponse) SetBody(v *ModifyWebIpSetSwitchResponseBody) *ModifyWebIpSetSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyWebIpSetSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
