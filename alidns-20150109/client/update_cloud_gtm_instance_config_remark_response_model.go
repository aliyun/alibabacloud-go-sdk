// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmInstanceConfigRemarkResponseBody) *UpdateCloudGtmInstanceConfigRemarkResponse
	GetBody() *UpdateCloudGtmInstanceConfigRemarkResponseBody
}

type UpdateCloudGtmInstanceConfigRemarkResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmInstanceConfigRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponse) GetBody() *UpdateCloudGtmInstanceConfigRemarkResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponse) SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponse) SetBody(v *UpdateCloudGtmInstanceConfigRemarkResponseBody) *UpdateCloudGtmInstanceConfigRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
