// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssScanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOssScanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOssScanConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateOssScanConfigResponseBody) *CreateOssScanConfigResponse
	GetBody() *CreateOssScanConfigResponseBody
}

type CreateOssScanConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOssScanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOssScanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOssScanConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateOssScanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOssScanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOssScanConfigResponse) GetBody() *CreateOssScanConfigResponseBody {
	return s.Body
}

func (s *CreateOssScanConfigResponse) SetHeaders(v map[string]*string) *CreateOssScanConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateOssScanConfigResponse) SetStatusCode(v int32) *CreateOssScanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOssScanConfigResponse) SetBody(v *CreateOssScanConfigResponseBody) *CreateOssScanConfigResponse {
	s.Body = v
	return s
}

func (s *CreateOssScanConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
