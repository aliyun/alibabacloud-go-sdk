// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRenderingInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRenderingInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRenderingInstanceAttributeResponseBody) *ModifyRenderingInstanceAttributeResponse
	GetBody() *ModifyRenderingInstanceAttributeResponseBody
}

type ModifyRenderingInstanceAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRenderingInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRenderingInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyRenderingInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRenderingInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRenderingInstanceAttributeResponse) GetBody() *ModifyRenderingInstanceAttributeResponseBody {
	return s.Body
}

func (s *ModifyRenderingInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyRenderingInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyRenderingInstanceAttributeResponse) SetStatusCode(v int32) *ModifyRenderingInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRenderingInstanceAttributeResponse) SetBody(v *ModifyRenderingInstanceAttributeResponseBody) *ModifyRenderingInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyRenderingInstanceAttributeResponse) Validate() error {
	return dara.Validate(s)
}
