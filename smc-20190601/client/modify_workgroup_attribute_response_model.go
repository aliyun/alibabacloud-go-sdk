// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWorkgroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWorkgroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWorkgroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWorkgroupAttributeResponseBody) *ModifyWorkgroupAttributeResponse
	GetBody() *ModifyWorkgroupAttributeResponseBody
}

type ModifyWorkgroupAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWorkgroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWorkgroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWorkgroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyWorkgroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWorkgroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWorkgroupAttributeResponse) GetBody() *ModifyWorkgroupAttributeResponseBody {
	return s.Body
}

func (s *ModifyWorkgroupAttributeResponse) SetHeaders(v map[string]*string) *ModifyWorkgroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyWorkgroupAttributeResponse) SetStatusCode(v int32) *ModifyWorkgroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWorkgroupAttributeResponse) SetBody(v *ModifyWorkgroupAttributeResponseBody) *ModifyWorkgroupAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyWorkgroupAttributeResponse) Validate() error {
	return dara.Validate(s)
}
