// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelSuggestV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelSuggestV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelSuggestV2Response
	GetStatusCode() *int32
	SetBody(v *HotelSuggestV2ResponseBody) *HotelSuggestV2Response
	GetBody() *HotelSuggestV2ResponseBody
}

type HotelSuggestV2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelSuggestV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelSuggestV2Response) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2Response) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelSuggestV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelSuggestV2Response) GetBody() *HotelSuggestV2ResponseBody {
	return s.Body
}

func (s *HotelSuggestV2Response) SetHeaders(v map[string]*string) *HotelSuggestV2Response {
	s.Headers = v
	return s
}

func (s *HotelSuggestV2Response) SetStatusCode(v int32) *HotelSuggestV2Response {
	s.StatusCode = &v
	return s
}

func (s *HotelSuggestV2Response) SetBody(v *HotelSuggestV2ResponseBody) *HotelSuggestV2Response {
	s.Body = v
	return s
}

func (s *HotelSuggestV2Response) Validate() error {
	return dara.Validate(s)
}
