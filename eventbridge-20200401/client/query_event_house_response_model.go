// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventHouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEventHouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEventHouseResponse
	GetStatusCode() *int32
	SetBody(v *QueryEventHouseResponseBody) *QueryEventHouseResponse
	GetBody() *QueryEventHouseResponseBody
}

type QueryEventHouseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventHouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventHouseResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEventHouseResponse) GoString() string {
	return s.String()
}

func (s *QueryEventHouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEventHouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEventHouseResponse) GetBody() *QueryEventHouseResponseBody {
	return s.Body
}

func (s *QueryEventHouseResponse) SetHeaders(v map[string]*string) *QueryEventHouseResponse {
	s.Headers = v
	return s
}

func (s *QueryEventHouseResponse) SetStatusCode(v int32) *QueryEventHouseResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventHouseResponse) SetBody(v *QueryEventHouseResponseBody) *QueryEventHouseResponse {
	s.Body = v
	return s
}

func (s *QueryEventHouseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
