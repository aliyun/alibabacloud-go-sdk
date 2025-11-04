// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodPackagingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVodPackagingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVodPackagingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVodPackagingConfigurationResponseBody) *DeleteVodPackagingConfigurationResponse
	GetBody() *DeleteVodPackagingConfigurationResponseBody
}

type DeleteVodPackagingConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVodPackagingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVodPackagingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodPackagingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodPackagingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVodPackagingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVodPackagingConfigurationResponse) GetBody() *DeleteVodPackagingConfigurationResponseBody {
	return s.Body
}

func (s *DeleteVodPackagingConfigurationResponse) SetHeaders(v map[string]*string) *DeleteVodPackagingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodPackagingConfigurationResponse) SetStatusCode(v int32) *DeleteVodPackagingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVodPackagingConfigurationResponse) SetBody(v *DeleteVodPackagingConfigurationResponseBody) *DeleteVodPackagingConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteVodPackagingConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
