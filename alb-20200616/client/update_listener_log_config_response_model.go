// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateListenerLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateListenerLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateListenerLogConfigResponseBody) *UpdateListenerLogConfigResponse
	GetBody() *UpdateListenerLogConfigResponseBody
}

type UpdateListenerLogConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateListenerLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateListenerLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerLogConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateListenerLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateListenerLogConfigResponse) GetBody() *UpdateListenerLogConfigResponseBody {
	return s.Body
}

func (s *UpdateListenerLogConfigResponse) SetHeaders(v map[string]*string) *UpdateListenerLogConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateListenerLogConfigResponse) SetStatusCode(v int32) *UpdateListenerLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateListenerLogConfigResponse) SetBody(v *UpdateListenerLogConfigResponseBody) *UpdateListenerLogConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateListenerLogConfigResponse) Validate() error {
	return dara.Validate(s)
}
