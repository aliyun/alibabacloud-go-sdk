// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmInstanceNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmInstanceNameResponseBody) *UpdateCloudGtmInstanceNameResponse
	GetBody() *UpdateCloudGtmInstanceNameResponseBody
}

type UpdateCloudGtmInstanceNameResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmInstanceNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmInstanceNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmInstanceNameResponse) GetBody() *UpdateCloudGtmInstanceNameResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmInstanceNameResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmInstanceNameResponse) SetStatusCode(v int32) *UpdateCloudGtmInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmInstanceNameResponse) SetBody(v *UpdateCloudGtmInstanceNameResponseBody) *UpdateCloudGtmInstanceNameResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmInstanceNameResponse) Validate() error {
	return dara.Validate(s)
}
