// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteJobRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteJobRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteJobRecordsResponseBody) *DeleteJobRecordsResponse
	GetBody() *DeleteJobRecordsResponseBody
}

type DeleteJobRecordsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJobRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobRecordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteJobRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteJobRecordsResponse) GetBody() *DeleteJobRecordsResponseBody {
	return s.Body
}

func (s *DeleteJobRecordsResponse) SetHeaders(v map[string]*string) *DeleteJobRecordsResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobRecordsResponse) SetStatusCode(v int32) *DeleteJobRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobRecordsResponse) SetBody(v *DeleteJobRecordsResponseBody) *DeleteJobRecordsResponse {
	s.Body = v
	return s
}

func (s *DeleteJobRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
