// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRenderingInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRenderingInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRenderingInstanceResponseBody) *ModifyRenderingInstanceResponse
	GetBody() *ModifyRenderingInstanceResponseBody
}

type ModifyRenderingInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRenderingInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRenderingInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyRenderingInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRenderingInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRenderingInstanceResponse) GetBody() *ModifyRenderingInstanceResponseBody {
	return s.Body
}

func (s *ModifyRenderingInstanceResponse) SetHeaders(v map[string]*string) *ModifyRenderingInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyRenderingInstanceResponse) SetStatusCode(v int32) *ModifyRenderingInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRenderingInstanceResponse) SetBody(v *ModifyRenderingInstanceResponseBody) *ModifyRenderingInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyRenderingInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
