// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TravelStandardRelateQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TravelStandardRelateQueryResponse
	GetStatusCode() *int32
	SetBody(v *TravelStandardRelateQueryResponseBody) *TravelStandardRelateQueryResponse
	GetBody() *TravelStandardRelateQueryResponseBody
}

type TravelStandardRelateQueryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TravelStandardRelateQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TravelStandardRelateQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateQueryResponse) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TravelStandardRelateQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TravelStandardRelateQueryResponse) GetBody() *TravelStandardRelateQueryResponseBody {
	return s.Body
}

func (s *TravelStandardRelateQueryResponse) SetHeaders(v map[string]*string) *TravelStandardRelateQueryResponse {
	s.Headers = v
	return s
}

func (s *TravelStandardRelateQueryResponse) SetStatusCode(v int32) *TravelStandardRelateQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TravelStandardRelateQueryResponse) SetBody(v *TravelStandardRelateQueryResponseBody) *TravelStandardRelateQueryResponse {
	s.Body = v
	return s
}

func (s *TravelStandardRelateQueryResponse) Validate() error {
	return dara.Validate(s)
}
