// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomIncidentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendCustomIncidentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendCustomIncidentsResponse
	GetStatusCode() *int32
	SetBody(v *SendCustomIncidentsResponseBody) *SendCustomIncidentsResponse
	GetBody() *SendCustomIncidentsResponseBody
}

type SendCustomIncidentsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCustomIncidentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendCustomIncidentsResponse) String() string {
	return dara.Prettify(s)
}

func (s SendCustomIncidentsResponse) GoString() string {
	return s.String()
}

func (s *SendCustomIncidentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendCustomIncidentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendCustomIncidentsResponse) GetBody() *SendCustomIncidentsResponseBody {
	return s.Body
}

func (s *SendCustomIncidentsResponse) SetHeaders(v map[string]*string) *SendCustomIncidentsResponse {
	s.Headers = v
	return s
}

func (s *SendCustomIncidentsResponse) SetStatusCode(v int32) *SendCustomIncidentsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCustomIncidentsResponse) SetBody(v *SendCustomIncidentsResponseBody) *SendCustomIncidentsResponse {
	s.Body = v
	return s
}

func (s *SendCustomIncidentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
