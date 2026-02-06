// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelJobsResponse
	GetStatusCode() *int32
	SetBody(v *CancelJobsResponseBody) *CancelJobsResponse
	GetBody() *CancelJobsResponseBody
}

type CancelJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelJobsResponse) GoString() string {
	return s.String()
}

func (s *CancelJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelJobsResponse) GetBody() *CancelJobsResponseBody {
	return s.Body
}

func (s *CancelJobsResponse) SetHeaders(v map[string]*string) *CancelJobsResponse {
	s.Headers = v
	return s
}

func (s *CancelJobsResponse) SetStatusCode(v int32) *CancelJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelJobsResponse) SetBody(v *CancelJobsResponseBody) *CancelJobsResponse {
	s.Body = v
	return s
}

func (s *CancelJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
