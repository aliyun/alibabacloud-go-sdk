// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmInstanceGlobalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDnsGtmInstanceGlobalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDnsGtmInstanceGlobalConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDnsGtmInstanceGlobalConfigResponseBody) *UpdateDnsGtmInstanceGlobalConfigResponse
	GetBody() *UpdateDnsGtmInstanceGlobalConfigResponseBody
}

type UpdateDnsGtmInstanceGlobalConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDnsGtmInstanceGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDnsGtmInstanceGlobalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmInstanceGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponse) GetBody() *UpdateDnsGtmInstanceGlobalConfigResponseBody {
	return s.Body
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponse) SetHeaders(v map[string]*string) *UpdateDnsGtmInstanceGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponse) SetStatusCode(v int32) *UpdateDnsGtmInstanceGlobalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponse) SetBody(v *UpdateDnsGtmInstanceGlobalConfigResponseBody) *UpdateDnsGtmInstanceGlobalConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponse) Validate() error {
	return dara.Validate(s)
}
