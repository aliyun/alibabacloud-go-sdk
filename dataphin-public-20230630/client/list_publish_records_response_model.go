// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublishRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublishRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListPublishRecordsResponseBody) *ListPublishRecordsResponse
	GetBody() *ListPublishRecordsResponseBody
}

type ListPublishRecordsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublishRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListPublishRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublishRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublishRecordsResponse) GetBody() *ListPublishRecordsResponseBody {
	return s.Body
}

func (s *ListPublishRecordsResponse) SetHeaders(v map[string]*string) *ListPublishRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListPublishRecordsResponse) SetStatusCode(v int32) *ListPublishRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishRecordsResponse) SetBody(v *ListPublishRecordsResponseBody) *ListPublishRecordsResponse {
	s.Body = v
	return s
}

func (s *ListPublishRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
