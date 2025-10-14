// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmGlobalAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmGlobalAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmGlobalAlertResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmGlobalAlertResponseBody) *UpdateCloudGtmGlobalAlertResponse
	GetBody() *UpdateCloudGtmGlobalAlertResponseBody
}

type UpdateCloudGtmGlobalAlertResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmGlobalAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmGlobalAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmGlobalAlertResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmGlobalAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmGlobalAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmGlobalAlertResponse) GetBody() *UpdateCloudGtmGlobalAlertResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmGlobalAlertResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmGlobalAlertResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmGlobalAlertResponse) SetStatusCode(v int32) *UpdateCloudGtmGlobalAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertResponse) SetBody(v *UpdateCloudGtmGlobalAlertResponseBody) *UpdateCloudGtmGlobalAlertResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmGlobalAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
