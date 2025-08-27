// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TravelStandardRelateDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TravelStandardRelateDeleteResponse
	GetStatusCode() *int32
	SetBody(v *TravelStandardRelateDeleteResponseBody) *TravelStandardRelateDeleteResponse
	GetBody() *TravelStandardRelateDeleteResponseBody
}

type TravelStandardRelateDeleteResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TravelStandardRelateDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TravelStandardRelateDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateDeleteResponse) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TravelStandardRelateDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TravelStandardRelateDeleteResponse) GetBody() *TravelStandardRelateDeleteResponseBody {
	return s.Body
}

func (s *TravelStandardRelateDeleteResponse) SetHeaders(v map[string]*string) *TravelStandardRelateDeleteResponse {
	s.Headers = v
	return s
}

func (s *TravelStandardRelateDeleteResponse) SetStatusCode(v int32) *TravelStandardRelateDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *TravelStandardRelateDeleteResponse) SetBody(v *TravelStandardRelateDeleteResponseBody) *TravelStandardRelateDeleteResponse {
	s.Body = v
	return s
}

func (s *TravelStandardRelateDeleteResponse) Validate() error {
	return dara.Validate(s)
}
