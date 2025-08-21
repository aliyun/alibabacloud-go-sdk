// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegionConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegionConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetRegionConfigurationResponseBody) *GetRegionConfigurationResponse
	GetBody() *GetRegionConfigurationResponseBody
}

type GetRegionConfigurationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegionConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegionConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegionConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegionConfigurationResponse) GetBody() *GetRegionConfigurationResponseBody {
	return s.Body
}

func (s *GetRegionConfigurationResponse) SetHeaders(v map[string]*string) *GetRegionConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetRegionConfigurationResponse) SetStatusCode(v int32) *GetRegionConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegionConfigurationResponse) SetBody(v *GetRegionConfigurationResponseBody) *GetRegionConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetRegionConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
