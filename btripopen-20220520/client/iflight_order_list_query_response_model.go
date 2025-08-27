// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIFlightOrderListQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IFlightOrderListQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IFlightOrderListQueryResponse
	GetStatusCode() *int32
	SetBody(v *IFlightOrderListQueryResponseBody) *IFlightOrderListQueryResponse
	GetBody() *IFlightOrderListQueryResponseBody
}

type IFlightOrderListQueryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IFlightOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IFlightOrderListQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IFlightOrderListQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IFlightOrderListQueryResponse) GetBody() *IFlightOrderListQueryResponseBody {
	return s.Body
}

func (s *IFlightOrderListQueryResponse) SetHeaders(v map[string]*string) *IFlightOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *IFlightOrderListQueryResponse) SetStatusCode(v int32) *IFlightOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *IFlightOrderListQueryResponse) SetBody(v *IFlightOrderListQueryResponseBody) *IFlightOrderListQueryResponse {
	s.Body = v
	return s
}

func (s *IFlightOrderListQueryResponse) Validate() error {
	return dara.Validate(s)
}
