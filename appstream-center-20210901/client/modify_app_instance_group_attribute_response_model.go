// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppInstanceGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppInstanceGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppInstanceGroupAttributeResponseBody) *ModifyAppInstanceGroupAttributeResponse
	GetBody() *ModifyAppInstanceGroupAttributeResponseBody
}

type ModifyAppInstanceGroupAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppInstanceGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppInstanceGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppInstanceGroupAttributeResponse) GetBody() *ModifyAppInstanceGroupAttributeResponseBody {
	return s.Body
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetHeaders(v map[string]*string) *ModifyAppInstanceGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetStatusCode(v int32) *ModifyAppInstanceGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetBody(v *ModifyAppInstanceGroupAttributeResponseBody) *ModifyAppInstanceGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponse) Validate() error {
	return dara.Validate(s)
}
