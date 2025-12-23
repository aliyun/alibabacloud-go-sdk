// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountDeliveryChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultiAccountDeliveryChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultiAccountDeliveryChannelsResponse
	GetStatusCode() *int32
	SetBody(v *ListMultiAccountDeliveryChannelsResponseBody) *ListMultiAccountDeliveryChannelsResponse
	GetBody() *ListMultiAccountDeliveryChannelsResponseBody
}

type ListMultiAccountDeliveryChannelsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountDeliveryChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountDeliveryChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountDeliveryChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultiAccountDeliveryChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultiAccountDeliveryChannelsResponse) GetBody() *ListMultiAccountDeliveryChannelsResponseBody {
	return s.Body
}

func (s *ListMultiAccountDeliveryChannelsResponse) SetHeaders(v map[string]*string) *ListMultiAccountDeliveryChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponse) SetStatusCode(v int32) *ListMultiAccountDeliveryChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponse) SetBody(v *ListMultiAccountDeliveryChannelsResponseBody) *ListMultiAccountDeliveryChannelsResponse {
	s.Body = v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
