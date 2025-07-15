// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonoRecordingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMonoRecordingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMonoRecordingsResponse
	GetStatusCode() *int32
	SetBody(v *ListMonoRecordingsResponseBody) *ListMonoRecordingsResponse
	GetBody() *ListMonoRecordingsResponseBody
}

type ListMonoRecordingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMonoRecordingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMonoRecordingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMonoRecordingsResponse) GoString() string {
	return s.String()
}

func (s *ListMonoRecordingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMonoRecordingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMonoRecordingsResponse) GetBody() *ListMonoRecordingsResponseBody {
	return s.Body
}

func (s *ListMonoRecordingsResponse) SetHeaders(v map[string]*string) *ListMonoRecordingsResponse {
	s.Headers = v
	return s
}

func (s *ListMonoRecordingsResponse) SetStatusCode(v int32) *ListMonoRecordingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMonoRecordingsResponse) SetBody(v *ListMonoRecordingsResponseBody) *ListMonoRecordingsResponse {
	s.Body = v
	return s
}

func (s *ListMonoRecordingsResponse) Validate() error {
	return dara.Validate(s)
}
