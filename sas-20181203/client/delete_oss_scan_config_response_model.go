// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOssScanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOssScanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOssScanConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOssScanConfigResponseBody) *DeleteOssScanConfigResponse
	GetBody() *DeleteOssScanConfigResponseBody
}

type DeleteOssScanConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOssScanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOssScanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOssScanConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteOssScanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOssScanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOssScanConfigResponse) GetBody() *DeleteOssScanConfigResponseBody {
	return s.Body
}

func (s *DeleteOssScanConfigResponse) SetHeaders(v map[string]*string) *DeleteOssScanConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteOssScanConfigResponse) SetStatusCode(v int32) *DeleteOssScanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOssScanConfigResponse) SetBody(v *DeleteOssScanConfigResponseBody) *DeleteOssScanConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteOssScanConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
