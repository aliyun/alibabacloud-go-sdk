// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardListQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TravelStandardListQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TravelStandardListQueryResponse
	GetStatusCode() *int32
	SetBody(v *TravelStandardListQueryResponseBody) *TravelStandardListQueryResponse
	GetBody() *TravelStandardListQueryResponseBody
}

type TravelStandardListQueryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TravelStandardListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TravelStandardListQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardListQueryResponse) GoString() string {
	return s.String()
}

func (s *TravelStandardListQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TravelStandardListQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TravelStandardListQueryResponse) GetBody() *TravelStandardListQueryResponseBody {
	return s.Body
}

func (s *TravelStandardListQueryResponse) SetHeaders(v map[string]*string) *TravelStandardListQueryResponse {
	s.Headers = v
	return s
}

func (s *TravelStandardListQueryResponse) SetStatusCode(v int32) *TravelStandardListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TravelStandardListQueryResponse) SetBody(v *TravelStandardListQueryResponseBody) *TravelStandardListQueryResponse {
	s.Body = v
	return s
}

func (s *TravelStandardListQueryResponse) Validate() error {
	return dara.Validate(s)
}
