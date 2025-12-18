// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigDeliveryChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigDeliveryChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigDeliveryChannelsResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigDeliveryChannelsResponseBody) *ListConfigDeliveryChannelsResponse
	GetBody() *ListConfigDeliveryChannelsResponseBody
}

type ListConfigDeliveryChannelsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigDeliveryChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigDeliveryChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigDeliveryChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigDeliveryChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigDeliveryChannelsResponse) GetBody() *ListConfigDeliveryChannelsResponseBody {
	return s.Body
}

func (s *ListConfigDeliveryChannelsResponse) SetHeaders(v map[string]*string) *ListConfigDeliveryChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigDeliveryChannelsResponse) SetStatusCode(v int32) *ListConfigDeliveryChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponse) SetBody(v *ListConfigDeliveryChannelsResponseBody) *ListConfigDeliveryChannelsResponse {
	s.Body = v
	return s
}

func (s *ListConfigDeliveryChannelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
