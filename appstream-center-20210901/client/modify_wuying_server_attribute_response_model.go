// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWuyingServerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWuyingServerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWuyingServerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWuyingServerAttributeResponseBody) *ModifyWuyingServerAttributeResponse
	GetBody() *ModifyWuyingServerAttributeResponseBody
}

type ModifyWuyingServerAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWuyingServerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWuyingServerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWuyingServerAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyWuyingServerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWuyingServerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWuyingServerAttributeResponse) GetBody() *ModifyWuyingServerAttributeResponseBody {
	return s.Body
}

func (s *ModifyWuyingServerAttributeResponse) SetHeaders(v map[string]*string) *ModifyWuyingServerAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyWuyingServerAttributeResponse) SetStatusCode(v int32) *ModifyWuyingServerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWuyingServerAttributeResponse) SetBody(v *ModifyWuyingServerAttributeResponseBody) *ModifyWuyingServerAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyWuyingServerAttributeResponse) Validate() error {
	return dara.Validate(s)
}
