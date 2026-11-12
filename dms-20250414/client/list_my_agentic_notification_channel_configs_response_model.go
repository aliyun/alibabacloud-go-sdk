// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyAgenticNotificationChannelConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMyAgenticNotificationChannelConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMyAgenticNotificationChannelConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListMyAgenticNotificationChannelConfigsResponseBody) *ListMyAgenticNotificationChannelConfigsResponse
	GetBody() *ListMyAgenticNotificationChannelConfigsResponseBody
}

type ListMyAgenticNotificationChannelConfigsResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyAgenticNotificationChannelConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyAgenticNotificationChannelConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMyAgenticNotificationChannelConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListMyAgenticNotificationChannelConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMyAgenticNotificationChannelConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMyAgenticNotificationChannelConfigsResponse) GetBody() *ListMyAgenticNotificationChannelConfigsResponseBody {
	return s.Body
}

func (s *ListMyAgenticNotificationChannelConfigsResponse) SetHeaders(v map[string]*string) *ListMyAgenticNotificationChannelConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponse) SetStatusCode(v int32) *ListMyAgenticNotificationChannelConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponse) SetBody(v *ListMyAgenticNotificationChannelConfigsResponseBody) *ListMyAgenticNotificationChannelConfigsResponse {
	s.Body = v
	return s
}

func (s *ListMyAgenticNotificationChannelConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
