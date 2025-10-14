// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSetRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSetRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSetRecordsResponseBody) *ListDataSetRecordsResponse
	GetBody() *ListDataSetRecordsResponseBody
}

type ListDataSetRecordsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSetRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSetRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListDataSetRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSetRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSetRecordsResponse) GetBody() *ListDataSetRecordsResponseBody {
	return s.Body
}

func (s *ListDataSetRecordsResponse) SetHeaders(v map[string]*string) *ListDataSetRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListDataSetRecordsResponse) SetStatusCode(v int32) *ListDataSetRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSetRecordsResponse) SetBody(v *ListDataSetRecordsResponseBody) *ListDataSetRecordsResponse {
	s.Body = v
	return s
}

func (s *ListDataSetRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
