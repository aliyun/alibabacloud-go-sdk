// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySwimmingLaneByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySwimmingLaneByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySwimmingLaneByIdResponse
	GetStatusCode() *int32
	SetBody(v *QuerySwimmingLaneByIdResponseBody) *QuerySwimmingLaneByIdResponse
	GetBody() *QuerySwimmingLaneByIdResponseBody
}

type QuerySwimmingLaneByIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySwimmingLaneByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySwimmingLaneByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySwimmingLaneByIdResponse) GoString() string {
	return s.String()
}

func (s *QuerySwimmingLaneByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySwimmingLaneByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySwimmingLaneByIdResponse) GetBody() *QuerySwimmingLaneByIdResponseBody {
	return s.Body
}

func (s *QuerySwimmingLaneByIdResponse) SetHeaders(v map[string]*string) *QuerySwimmingLaneByIdResponse {
	s.Headers = v
	return s
}

func (s *QuerySwimmingLaneByIdResponse) SetStatusCode(v int32) *QuerySwimmingLaneByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponse) SetBody(v *QuerySwimmingLaneByIdResponseBody) *QuerySwimmingLaneByIdResponse {
	s.Body = v
	return s
}

func (s *QuerySwimmingLaneByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
