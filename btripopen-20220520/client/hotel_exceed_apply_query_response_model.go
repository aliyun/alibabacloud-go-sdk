// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelExceedApplyQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelExceedApplyQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelExceedApplyQueryResponse
	GetStatusCode() *int32
	SetBody(v *HotelExceedApplyQueryResponseBody) *HotelExceedApplyQueryResponse
	GetBody() *HotelExceedApplyQueryResponseBody
}

type HotelExceedApplyQueryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelExceedApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelExceedApplyQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelExceedApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelExceedApplyQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelExceedApplyQueryResponse) GetBody() *HotelExceedApplyQueryResponseBody {
	return s.Body
}

func (s *HotelExceedApplyQueryResponse) SetHeaders(v map[string]*string) *HotelExceedApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *HotelExceedApplyQueryResponse) SetStatusCode(v int32) *HotelExceedApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelExceedApplyQueryResponse) SetBody(v *HotelExceedApplyQueryResponseBody) *HotelExceedApplyQueryResponse {
	s.Body = v
	return s
}

func (s *HotelExceedApplyQueryResponse) Validate() error {
	return dara.Validate(s)
}
