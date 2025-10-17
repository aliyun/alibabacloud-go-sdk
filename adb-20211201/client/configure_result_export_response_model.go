// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureResultExportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureResultExportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureResultExportResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureResultExportResponseBody) *ConfigureResultExportResponse
	GetBody() *ConfigureResultExportResponseBody
}

type ConfigureResultExportResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureResultExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureResultExportResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureResultExportResponse) GoString() string {
	return s.String()
}

func (s *ConfigureResultExportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureResultExportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureResultExportResponse) GetBody() *ConfigureResultExportResponseBody {
	return s.Body
}

func (s *ConfigureResultExportResponse) SetHeaders(v map[string]*string) *ConfigureResultExportResponse {
	s.Headers = v
	return s
}

func (s *ConfigureResultExportResponse) SetStatusCode(v int32) *ConfigureResultExportResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureResultExportResponse) SetBody(v *ConfigureResultExportResponseBody) *ConfigureResultExportResponse {
	s.Body = v
	return s
}

func (s *ConfigureResultExportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
