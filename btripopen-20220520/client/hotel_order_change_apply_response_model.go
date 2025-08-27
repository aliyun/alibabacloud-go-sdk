// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderChangeApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderChangeApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderChangeApplyResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderChangeApplyResponseBody) *HotelOrderChangeApplyResponse
	GetBody() *HotelOrderChangeApplyResponseBody
}

type HotelOrderChangeApplyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderChangeApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderChangeApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeApplyResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderChangeApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderChangeApplyResponse) GetBody() *HotelOrderChangeApplyResponseBody {
	return s.Body
}

func (s *HotelOrderChangeApplyResponse) SetHeaders(v map[string]*string) *HotelOrderChangeApplyResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderChangeApplyResponse) SetStatusCode(v int32) *HotelOrderChangeApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderChangeApplyResponse) SetBody(v *HotelOrderChangeApplyResponseBody) *HotelOrderChangeApplyResponse {
	s.Body = v
	return s
}

func (s *HotelOrderChangeApplyResponse) Validate() error {
	return dara.Validate(s)
}
