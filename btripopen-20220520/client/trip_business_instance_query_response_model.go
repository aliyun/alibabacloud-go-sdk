// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripBusinessInstanceQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TripBusinessInstanceQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TripBusinessInstanceQueryResponse
	GetStatusCode() *int32
	SetBody(v *TripBusinessInstanceQueryResponseBody) *TripBusinessInstanceQueryResponse
	GetBody() *TripBusinessInstanceQueryResponseBody
}

type TripBusinessInstanceQueryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TripBusinessInstanceQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TripBusinessInstanceQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TripBusinessInstanceQueryResponse) GoString() string {
	return s.String()
}

func (s *TripBusinessInstanceQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TripBusinessInstanceQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TripBusinessInstanceQueryResponse) GetBody() *TripBusinessInstanceQueryResponseBody {
	return s.Body
}

func (s *TripBusinessInstanceQueryResponse) SetHeaders(v map[string]*string) *TripBusinessInstanceQueryResponse {
	s.Headers = v
	return s
}

func (s *TripBusinessInstanceQueryResponse) SetStatusCode(v int32) *TripBusinessInstanceQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TripBusinessInstanceQueryResponse) SetBody(v *TripBusinessInstanceQueryResponseBody) *TripBusinessInstanceQueryResponse {
	s.Body = v
	return s
}

func (s *TripBusinessInstanceQueryResponse) Validate() error {
	return dara.Validate(s)
}
