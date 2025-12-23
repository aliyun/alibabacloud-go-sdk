// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeliveryChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeliveryChannelsResponse
	GetStatusCode() *int32
	SetBody(v *ListDeliveryChannelsResponseBody) *ListDeliveryChannelsResponse
	GetBody() *ListDeliveryChannelsResponseBody
}

type ListDeliveryChannelsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeliveryChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeliveryChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListDeliveryChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeliveryChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeliveryChannelsResponse) GetBody() *ListDeliveryChannelsResponseBody {
	return s.Body
}

func (s *ListDeliveryChannelsResponse) SetHeaders(v map[string]*string) *ListDeliveryChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListDeliveryChannelsResponse) SetStatusCode(v int32) *ListDeliveryChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeliveryChannelsResponse) SetBody(v *ListDeliveryChannelsResponseBody) *ListDeliveryChannelsResponse {
	s.Body = v
	return s
}

func (s *ListDeliveryChannelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
