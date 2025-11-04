// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodPackagingConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVodPackagingConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVodPackagingConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListVodPackagingConfigurationsResponseBody) *ListVodPackagingConfigurationsResponse
	GetBody() *ListVodPackagingConfigurationsResponseBody
}

type ListVodPackagingConfigurationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVodPackagingConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVodPackagingConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListVodPackagingConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVodPackagingConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVodPackagingConfigurationsResponse) GetBody() *ListVodPackagingConfigurationsResponseBody {
	return s.Body
}

func (s *ListVodPackagingConfigurationsResponse) SetHeaders(v map[string]*string) *ListVodPackagingConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListVodPackagingConfigurationsResponse) SetStatusCode(v int32) *ListVodPackagingConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVodPackagingConfigurationsResponse) SetBody(v *ListVodPackagingConfigurationsResponseBody) *ListVodPackagingConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListVodPackagingConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
