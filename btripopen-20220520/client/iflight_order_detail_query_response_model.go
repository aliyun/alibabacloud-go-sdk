// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIFlightOrderDetailQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IFlightOrderDetailQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IFlightOrderDetailQueryResponse
	GetStatusCode() *int32
	SetBody(v *IFlightOrderDetailQueryResponseBody) *IFlightOrderDetailQueryResponse
	GetBody() *IFlightOrderDetailQueryResponseBody
}

type IFlightOrderDetailQueryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IFlightOrderDetailQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IFlightOrderDetailQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponse) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IFlightOrderDetailQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IFlightOrderDetailQueryResponse) GetBody() *IFlightOrderDetailQueryResponseBody {
	return s.Body
}

func (s *IFlightOrderDetailQueryResponse) SetHeaders(v map[string]*string) *IFlightOrderDetailQueryResponse {
	s.Headers = v
	return s
}

func (s *IFlightOrderDetailQueryResponse) SetStatusCode(v int32) *IFlightOrderDetailQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponse) SetBody(v *IFlightOrderDetailQueryResponseBody) *IFlightOrderDetailQueryResponse {
	s.Body = v
	return s
}

func (s *IFlightOrderDetailQueryResponse) Validate() error {
	return dara.Validate(s)
}
