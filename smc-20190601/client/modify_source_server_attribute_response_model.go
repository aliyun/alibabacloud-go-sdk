// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySourceServerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySourceServerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySourceServerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifySourceServerAttributeResponseBody) *ModifySourceServerAttributeResponse
	GetBody() *ModifySourceServerAttributeResponseBody
}

type ModifySourceServerAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySourceServerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySourceServerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySourceServerAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifySourceServerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySourceServerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySourceServerAttributeResponse) GetBody() *ModifySourceServerAttributeResponseBody {
	return s.Body
}

func (s *ModifySourceServerAttributeResponse) SetHeaders(v map[string]*string) *ModifySourceServerAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifySourceServerAttributeResponse) SetStatusCode(v int32) *ModifySourceServerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySourceServerAttributeResponse) SetBody(v *ModifySourceServerAttributeResponseBody) *ModifySourceServerAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifySourceServerAttributeResponse) Validate() error {
	return dara.Validate(s)
}
