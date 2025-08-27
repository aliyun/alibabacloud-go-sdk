// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TravelStandardQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TravelStandardQueryResponse
	GetStatusCode() *int32
	SetBody(v *TravelStandardQueryResponseBody) *TravelStandardQueryResponse
	GetBody() *TravelStandardQueryResponseBody
}

type TravelStandardQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TravelStandardQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TravelStandardQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardQueryResponse) GoString() string {
	return s.String()
}

func (s *TravelStandardQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TravelStandardQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TravelStandardQueryResponse) GetBody() *TravelStandardQueryResponseBody {
	return s.Body
}

func (s *TravelStandardQueryResponse) SetHeaders(v map[string]*string) *TravelStandardQueryResponse {
	s.Headers = v
	return s
}

func (s *TravelStandardQueryResponse) SetStatusCode(v int32) *TravelStandardQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TravelStandardQueryResponse) SetBody(v *TravelStandardQueryResponseBody) *TravelStandardQueryResponse {
	s.Body = v
	return s
}

func (s *TravelStandardQueryResponse) Validate() error {
	return dara.Validate(s)
}
