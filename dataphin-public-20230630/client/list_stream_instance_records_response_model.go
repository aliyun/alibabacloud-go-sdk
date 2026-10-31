// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamInstanceRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStreamInstanceRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStreamInstanceRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListStreamInstanceRecordsResponseBody) *ListStreamInstanceRecordsResponse
	GetBody() *ListStreamInstanceRecordsResponseBody
}

type ListStreamInstanceRecordsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStreamInstanceRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStreamInstanceRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStreamInstanceRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListStreamInstanceRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStreamInstanceRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStreamInstanceRecordsResponse) GetBody() *ListStreamInstanceRecordsResponseBody {
	return s.Body
}

func (s *ListStreamInstanceRecordsResponse) SetHeaders(v map[string]*string) *ListStreamInstanceRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListStreamInstanceRecordsResponse) SetStatusCode(v int32) *ListStreamInstanceRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStreamInstanceRecordsResponse) SetBody(v *ListStreamInstanceRecordsResponseBody) *ListStreamInstanceRecordsResponse {
	s.Body = v
	return s
}

func (s *ListStreamInstanceRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
