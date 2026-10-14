// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAifgJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAifgJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAifgJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListAifgJobsResponseBody) *ListAifgJobsResponse
	GetBody() *ListAifgJobsResponseBody
}

type ListAifgJobsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAifgJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAifgJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAifgJobsResponse) GoString() string {
	return s.String()
}

func (s *ListAifgJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAifgJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAifgJobsResponse) GetBody() *ListAifgJobsResponseBody {
	return s.Body
}

func (s *ListAifgJobsResponse) SetHeaders(v map[string]*string) *ListAifgJobsResponse {
	s.Headers = v
	return s
}

func (s *ListAifgJobsResponse) SetStatusCode(v int32) *ListAifgJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAifgJobsResponse) SetBody(v *ListAifgJobsResponseBody) *ListAifgJobsResponse {
	s.Body = v
	return s
}

func (s *ListAifgJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
