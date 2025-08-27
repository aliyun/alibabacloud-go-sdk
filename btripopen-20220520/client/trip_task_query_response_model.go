// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripTaskQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TripTaskQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TripTaskQueryResponse
	GetStatusCode() *int32
	SetBody(v *TripTaskQueryResponseBody) *TripTaskQueryResponse
	GetBody() *TripTaskQueryResponseBody
}

type TripTaskQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TripTaskQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TripTaskQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TripTaskQueryResponse) GoString() string {
	return s.String()
}

func (s *TripTaskQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TripTaskQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TripTaskQueryResponse) GetBody() *TripTaskQueryResponseBody {
	return s.Body
}

func (s *TripTaskQueryResponse) SetHeaders(v map[string]*string) *TripTaskQueryResponse {
	s.Headers = v
	return s
}

func (s *TripTaskQueryResponse) SetStatusCode(v int32) *TripTaskQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TripTaskQueryResponse) SetBody(v *TripTaskQueryResponseBody) *TripTaskQueryResponse {
	s.Body = v
	return s
}

func (s *TripTaskQueryResponse) Validate() error {
	return dara.Validate(s)
}
