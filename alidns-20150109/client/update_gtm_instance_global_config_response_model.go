// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmInstanceGlobalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGtmInstanceGlobalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGtmInstanceGlobalConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGtmInstanceGlobalConfigResponseBody) *UpdateGtmInstanceGlobalConfigResponse
	GetBody() *UpdateGtmInstanceGlobalConfigResponseBody
}

type UpdateGtmInstanceGlobalConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGtmInstanceGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGtmInstanceGlobalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmInstanceGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmInstanceGlobalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGtmInstanceGlobalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGtmInstanceGlobalConfigResponse) GetBody() *UpdateGtmInstanceGlobalConfigResponseBody {
	return s.Body
}

func (s *UpdateGtmInstanceGlobalConfigResponse) SetHeaders(v map[string]*string) *UpdateGtmInstanceGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigResponse) SetStatusCode(v int32) *UpdateGtmInstanceGlobalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigResponse) SetBody(v *UpdateGtmInstanceGlobalConfigResponseBody) *UpdateGtmInstanceGlobalConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigResponse) Validate() error {
	return dara.Validate(s)
}
