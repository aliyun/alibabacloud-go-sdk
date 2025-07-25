// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigEnableStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigEnableStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigEnableStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmInstanceConfigEnableStatusResponseBody) *UpdateCloudGtmInstanceConfigEnableStatusResponse
	GetBody() *UpdateCloudGtmInstanceConfigEnableStatusResponseBody
}

type UpdateCloudGtmInstanceConfigEnableStatusResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmInstanceConfigEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigEnableStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponse) GetBody() *UpdateCloudGtmInstanceConfigEnableStatusResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponse) SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponse) SetBody(v *UpdateCloudGtmInstanceConfigEnableStatusResponseBody) *UpdateCloudGtmInstanceConfigEnableStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponse) Validate() error {
	return dara.Validate(s)
}
