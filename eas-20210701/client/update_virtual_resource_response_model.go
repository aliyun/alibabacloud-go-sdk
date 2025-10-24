// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirtualResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVirtualResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVirtualResourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVirtualResourceResponseBody) *UpdateVirtualResourceResponse
	GetBody() *UpdateVirtualResourceResponseBody
}

type UpdateVirtualResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVirtualResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVirtualResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateVirtualResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVirtualResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVirtualResourceResponse) GetBody() *UpdateVirtualResourceResponseBody {
	return s.Body
}

func (s *UpdateVirtualResourceResponse) SetHeaders(v map[string]*string) *UpdateVirtualResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateVirtualResourceResponse) SetStatusCode(v int32) *UpdateVirtualResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVirtualResourceResponse) SetBody(v *UpdateVirtualResourceResponseBody) *UpdateVirtualResourceResponse {
	s.Body = v
	return s
}

func (s *UpdateVirtualResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
