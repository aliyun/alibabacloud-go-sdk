// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssScanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssScanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssScanConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetOssScanConfigResponseBody) *GetOssScanConfigResponse
	GetBody() *GetOssScanConfigResponseBody
}

type GetOssScanConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssScanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssScanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssScanConfigResponse) GoString() string {
	return s.String()
}

func (s *GetOssScanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssScanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssScanConfigResponse) GetBody() *GetOssScanConfigResponseBody {
	return s.Body
}

func (s *GetOssScanConfigResponse) SetHeaders(v map[string]*string) *GetOssScanConfigResponse {
	s.Headers = v
	return s
}

func (s *GetOssScanConfigResponse) SetStatusCode(v int32) *GetOssScanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssScanConfigResponse) SetBody(v *GetOssScanConfigResponseBody) *GetOssScanConfigResponse {
	s.Body = v
	return s
}

func (s *GetOssScanConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
