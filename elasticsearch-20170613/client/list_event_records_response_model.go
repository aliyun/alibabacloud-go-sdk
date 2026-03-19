// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListEventRecordsResponseBody) *ListEventRecordsResponse
	GetBody() *ListEventRecordsResponseBody
}

type ListEventRecordsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListEventRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventRecordsResponse) GetBody() *ListEventRecordsResponseBody {
	return s.Body
}

func (s *ListEventRecordsResponse) SetHeaders(v map[string]*string) *ListEventRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListEventRecordsResponse) SetStatusCode(v int32) *ListEventRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventRecordsResponse) SetBody(v *ListEventRecordsResponseBody) *ListEventRecordsResponse {
	s.Body = v
	return s
}

func (s *ListEventRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
