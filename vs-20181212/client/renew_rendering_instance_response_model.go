// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewRenderingInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewRenderingInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewRenderingInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RenewRenderingInstanceResponseBody) *RenewRenderingInstanceResponse
	GetBody() *RenewRenderingInstanceResponseBody
}

type RenewRenderingInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewRenderingInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewRenderingInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewRenderingInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewRenderingInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewRenderingInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewRenderingInstanceResponse) GetBody() *RenewRenderingInstanceResponseBody {
	return s.Body
}

func (s *RenewRenderingInstanceResponse) SetHeaders(v map[string]*string) *RenewRenderingInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewRenderingInstanceResponse) SetStatusCode(v int32) *RenewRenderingInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewRenderingInstanceResponse) SetBody(v *RenewRenderingInstanceResponseBody) *RenewRenderingInstanceResponse {
	s.Body = v
	return s
}

func (s *RenewRenderingInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
