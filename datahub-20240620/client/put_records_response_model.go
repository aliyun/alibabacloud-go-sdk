// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutRecordsResponse
	GetStatusCode() *int32
	SetBody(v *PutRecordsResponseBody) *PutRecordsResponse
	GetBody() *PutRecordsResponseBody
}

type PutRecordsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutRecordsResponse) GoString() string {
	return s.String()
}

func (s *PutRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutRecordsResponse) GetBody() *PutRecordsResponseBody {
	return s.Body
}

func (s *PutRecordsResponse) SetHeaders(v map[string]*string) *PutRecordsResponse {
	s.Headers = v
	return s
}

func (s *PutRecordsResponse) SetStatusCode(v int32) *PutRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutRecordsResponse) SetBody(v *PutRecordsResponseBody) *PutRecordsResponse {
	s.Body = v
	return s
}

func (s *PutRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
