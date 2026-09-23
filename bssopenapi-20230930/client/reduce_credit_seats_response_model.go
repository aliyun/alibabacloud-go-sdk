// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReduceCreditSeatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReduceCreditSeatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReduceCreditSeatsResponse
	GetStatusCode() *int32
	SetBody(v *ReduceCreditSeatsResponseBody) *ReduceCreditSeatsResponse
	GetBody() *ReduceCreditSeatsResponseBody
}

type ReduceCreditSeatsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReduceCreditSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReduceCreditSeatsResponse) String() string {
	return dara.Prettify(s)
}

func (s ReduceCreditSeatsResponse) GoString() string {
	return s.String()
}

func (s *ReduceCreditSeatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReduceCreditSeatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReduceCreditSeatsResponse) GetBody() *ReduceCreditSeatsResponseBody {
	return s.Body
}

func (s *ReduceCreditSeatsResponse) SetHeaders(v map[string]*string) *ReduceCreditSeatsResponse {
	s.Headers = v
	return s
}

func (s *ReduceCreditSeatsResponse) SetStatusCode(v int32) *ReduceCreditSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReduceCreditSeatsResponse) SetBody(v *ReduceCreditSeatsResponseBody) *ReduceCreditSeatsResponse {
	s.Body = v
	return s
}

func (s *ReduceCreditSeatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
