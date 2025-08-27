// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateAddResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TravelStandardRelateAddResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TravelStandardRelateAddResponse
	GetStatusCode() *int32
	SetBody(v *TravelStandardRelateAddResponseBody) *TravelStandardRelateAddResponse
	GetBody() *TravelStandardRelateAddResponseBody
}

type TravelStandardRelateAddResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TravelStandardRelateAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TravelStandardRelateAddResponse) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateAddResponse) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateAddResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TravelStandardRelateAddResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TravelStandardRelateAddResponse) GetBody() *TravelStandardRelateAddResponseBody {
	return s.Body
}

func (s *TravelStandardRelateAddResponse) SetHeaders(v map[string]*string) *TravelStandardRelateAddResponse {
	s.Headers = v
	return s
}

func (s *TravelStandardRelateAddResponse) SetStatusCode(v int32) *TravelStandardRelateAddResponse {
	s.StatusCode = &v
	return s
}

func (s *TravelStandardRelateAddResponse) SetBody(v *TravelStandardRelateAddResponseBody) *TravelStandardRelateAddResponse {
	s.Body = v
	return s
}

func (s *TravelStandardRelateAddResponse) Validate() error {
	return dara.Validate(s)
}
