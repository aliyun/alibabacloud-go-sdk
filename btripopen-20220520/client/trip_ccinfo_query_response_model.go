// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripCCInfoQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TripCCInfoQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TripCCInfoQueryResponse
	GetStatusCode() *int32
	SetBody(v *TripCCInfoQueryResponseBody) *TripCCInfoQueryResponse
	GetBody() *TripCCInfoQueryResponseBody
}

type TripCCInfoQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TripCCInfoQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TripCCInfoQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TripCCInfoQueryResponse) GoString() string {
	return s.String()
}

func (s *TripCCInfoQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TripCCInfoQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TripCCInfoQueryResponse) GetBody() *TripCCInfoQueryResponseBody {
	return s.Body
}

func (s *TripCCInfoQueryResponse) SetHeaders(v map[string]*string) *TripCCInfoQueryResponse {
	s.Headers = v
	return s
}

func (s *TripCCInfoQueryResponse) SetStatusCode(v int32) *TripCCInfoQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TripCCInfoQueryResponse) SetBody(v *TripCCInfoQueryResponseBody) *TripCCInfoQueryResponse {
	s.Body = v
	return s
}

func (s *TripCCInfoQueryResponse) Validate() error {
	return dara.Validate(s)
}
