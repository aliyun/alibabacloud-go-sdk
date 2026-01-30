// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssScanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOssScanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOssScanConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOssScanConfigResponseBody) *UpdateOssScanConfigResponse
	GetBody() *UpdateOssScanConfigResponseBody
}

type UpdateOssScanConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOssScanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOssScanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssScanConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssScanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOssScanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOssScanConfigResponse) GetBody() *UpdateOssScanConfigResponseBody {
	return s.Body
}

func (s *UpdateOssScanConfigResponse) SetHeaders(v map[string]*string) *UpdateOssScanConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssScanConfigResponse) SetStatusCode(v int32) *UpdateOssScanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOssScanConfigResponse) SetBody(v *UpdateOssScanConfigResponseBody) *UpdateOssScanConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateOssScanConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
