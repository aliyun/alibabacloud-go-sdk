// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaInfoJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaInfoJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaInfoJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaInfoJobsResponseBody) *ListMediaInfoJobsResponse
	GetBody() *ListMediaInfoJobsResponseBody
}

type ListMediaInfoJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaInfoJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaInfoJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaInfoJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaInfoJobsResponse) GetBody() *ListMediaInfoJobsResponseBody {
	return s.Body
}

func (s *ListMediaInfoJobsResponse) SetHeaders(v map[string]*string) *ListMediaInfoJobsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaInfoJobsResponse) SetStatusCode(v int32) *ListMediaInfoJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaInfoJobsResponse) SetBody(v *ListMediaInfoJobsResponseBody) *ListMediaInfoJobsResponse {
	s.Body = v
	return s
}

func (s *ListMediaInfoJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
