// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigAlertResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmInstanceConfigAlertResponseBody) *UpdateCloudGtmInstanceConfigAlertResponse
	GetBody() *UpdateCloudGtmInstanceConfigAlertResponseBody
}

type UpdateCloudGtmInstanceConfigAlertResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmInstanceConfigAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigAlertResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmInstanceConfigAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmInstanceConfigAlertResponse) GetBody() *UpdateCloudGtmInstanceConfigAlertResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmInstanceConfigAlertResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigAlertResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertResponse) SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertResponse) SetBody(v *UpdateCloudGtmInstanceConfigAlertResponseBody) *UpdateCloudGtmInstanceConfigAlertResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
