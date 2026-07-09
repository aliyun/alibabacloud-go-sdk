// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceRecordsResponseBody) *ListServiceRecordsResponse
	GetBody() *ListServiceRecordsResponseBody
}

type ListServiceRecordsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceRecordsResponse) GetBody() *ListServiceRecordsResponseBody {
	return s.Body
}

func (s *ListServiceRecordsResponse) SetHeaders(v map[string]*string) *ListServiceRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceRecordsResponse) SetStatusCode(v int32) *ListServiceRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceRecordsResponse) SetBody(v *ListServiceRecordsResponseBody) *ListServiceRecordsResponse {
	s.Body = v
	return s
}

func (s *ListServiceRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
