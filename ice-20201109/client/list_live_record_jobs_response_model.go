// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveRecordJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveRecordJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveRecordJobsResponseBody) *ListLiveRecordJobsResponse
	GetBody() *ListLiveRecordJobsResponseBody
}

type ListLiveRecordJobsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveRecordJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveRecordJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordJobsResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRecordJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveRecordJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveRecordJobsResponse) GetBody() *ListLiveRecordJobsResponseBody {
	return s.Body
}

func (s *ListLiveRecordJobsResponse) SetHeaders(v map[string]*string) *ListLiveRecordJobsResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRecordJobsResponse) SetStatusCode(v int32) *ListLiveRecordJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveRecordJobsResponse) SetBody(v *ListLiveRecordJobsResponseBody) *ListLiveRecordJobsResponse {
	s.Body = v
	return s
}

func (s *ListLiveRecordJobsResponse) Validate() error {
	return dara.Validate(s)
}
