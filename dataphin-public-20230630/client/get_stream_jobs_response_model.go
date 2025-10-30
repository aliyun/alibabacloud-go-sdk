// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStreamJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStreamJobsResponse
	GetStatusCode() *int32
	SetBody(v *GetStreamJobsResponseBody) *GetStreamJobsResponse
	GetBody() *GetStreamJobsResponseBody
}

type GetStreamJobsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStreamJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStreamJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStreamJobsResponse) GoString() string {
	return s.String()
}

func (s *GetStreamJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStreamJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStreamJobsResponse) GetBody() *GetStreamJobsResponseBody {
	return s.Body
}

func (s *GetStreamJobsResponse) SetHeaders(v map[string]*string) *GetStreamJobsResponse {
	s.Headers = v
	return s
}

func (s *GetStreamJobsResponse) SetStatusCode(v int32) *GetStreamJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStreamJobsResponse) SetBody(v *GetStreamJobsResponseBody) *GetStreamJobsResponse {
	s.Body = v
	return s
}

func (s *GetStreamJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
