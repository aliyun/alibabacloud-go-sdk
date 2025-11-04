// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVodPackagingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVodPackagingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateVodPackagingConfigurationResponseBody) *CreateVodPackagingConfigurationResponse
	GetBody() *CreateVodPackagingConfigurationResponseBody
}

type CreateVodPackagingConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVodPackagingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVodPackagingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVodPackagingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVodPackagingConfigurationResponse) GetBody() *CreateVodPackagingConfigurationResponseBody {
	return s.Body
}

func (s *CreateVodPackagingConfigurationResponse) SetHeaders(v map[string]*string) *CreateVodPackagingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateVodPackagingConfigurationResponse) SetStatusCode(v int32) *CreateVodPackagingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVodPackagingConfigurationResponse) SetBody(v *CreateVodPackagingConfigurationResponseBody) *CreateVodPackagingConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateVodPackagingConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
