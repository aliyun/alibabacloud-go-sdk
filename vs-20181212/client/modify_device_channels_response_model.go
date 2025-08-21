// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDeviceChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDeviceChannelsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDeviceChannelsResponseBody) *ModifyDeviceChannelsResponse
	GetBody() *ModifyDeviceChannelsResponseBody
}

type ModifyDeviceChannelsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDeviceChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDeviceChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceChannelsResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDeviceChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDeviceChannelsResponse) GetBody() *ModifyDeviceChannelsResponseBody {
	return s.Body
}

func (s *ModifyDeviceChannelsResponse) SetHeaders(v map[string]*string) *ModifyDeviceChannelsResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceChannelsResponse) SetStatusCode(v int32) *ModifyDeviceChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeviceChannelsResponse) SetBody(v *ModifyDeviceChannelsResponseBody) *ModifyDeviceChannelsResponse {
	s.Body = v
	return s
}

func (s *ModifyDeviceChannelsResponse) Validate() error {
	return dara.Validate(s)
}
