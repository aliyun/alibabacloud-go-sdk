// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorHotelEventPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CooperatorHotelEventPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CooperatorHotelEventPushResponse
	GetStatusCode() *int32
	SetBody(v *CooperatorHotelEventPushResponseBody) *CooperatorHotelEventPushResponse
	GetBody() *CooperatorHotelEventPushResponseBody
}

type CooperatorHotelEventPushResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CooperatorHotelEventPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CooperatorHotelEventPushResponse) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelEventPushResponse) GoString() string {
	return s.String()
}

func (s *CooperatorHotelEventPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CooperatorHotelEventPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CooperatorHotelEventPushResponse) GetBody() *CooperatorHotelEventPushResponseBody {
	return s.Body
}

func (s *CooperatorHotelEventPushResponse) SetHeaders(v map[string]*string) *CooperatorHotelEventPushResponse {
	s.Headers = v
	return s
}

func (s *CooperatorHotelEventPushResponse) SetStatusCode(v int32) *CooperatorHotelEventPushResponse {
	s.StatusCode = &v
	return s
}

func (s *CooperatorHotelEventPushResponse) SetBody(v *CooperatorHotelEventPushResponseBody) *CooperatorHotelEventPushResponse {
	s.Body = v
	return s
}

func (s *CooperatorHotelEventPushResponse) Validate() error {
	return dara.Validate(s)
}
