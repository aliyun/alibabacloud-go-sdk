// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssScanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOssScanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOssScanConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListOssScanConfigResponseBody) *ListOssScanConfigResponse
	GetBody() *ListOssScanConfigResponseBody
}

type ListOssScanConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOssScanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOssScanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOssScanConfigResponse) GoString() string {
	return s.String()
}

func (s *ListOssScanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOssScanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOssScanConfigResponse) GetBody() *ListOssScanConfigResponseBody {
	return s.Body
}

func (s *ListOssScanConfigResponse) SetHeaders(v map[string]*string) *ListOssScanConfigResponse {
	s.Headers = v
	return s
}

func (s *ListOssScanConfigResponse) SetStatusCode(v int32) *ListOssScanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOssScanConfigResponse) SetBody(v *ListOssScanConfigResponseBody) *ListOssScanConfigResponse {
	s.Body = v
	return s
}

func (s *ListOssScanConfigResponse) Validate() error {
	return dara.Validate(s)
}
