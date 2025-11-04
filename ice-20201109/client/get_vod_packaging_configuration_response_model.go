// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPackagingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVodPackagingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVodPackagingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetVodPackagingConfigurationResponseBody) *GetVodPackagingConfigurationResponse
	GetBody() *GetVodPackagingConfigurationResponseBody
}

type GetVodPackagingConfigurationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVodPackagingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVodPackagingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetVodPackagingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVodPackagingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVodPackagingConfigurationResponse) GetBody() *GetVodPackagingConfigurationResponseBody {
	return s.Body
}

func (s *GetVodPackagingConfigurationResponse) SetHeaders(v map[string]*string) *GetVodPackagingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetVodPackagingConfigurationResponse) SetStatusCode(v int32) *GetVodPackagingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVodPackagingConfigurationResponse) SetBody(v *GetVodPackagingConfigurationResponseBody) *GetVodPackagingConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetVodPackagingConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
