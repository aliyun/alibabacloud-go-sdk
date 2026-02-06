// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssignJobsAsyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssignJobsAsyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssignJobsAsyncResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAssignJobsAsyncResultResponseBody) *GetAssignJobsAsyncResultResponse
	GetBody() *GetAssignJobsAsyncResultResponseBody
}

type GetAssignJobsAsyncResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssignJobsAsyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssignJobsAsyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssignJobsAsyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetAssignJobsAsyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssignJobsAsyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssignJobsAsyncResultResponse) GetBody() *GetAssignJobsAsyncResultResponseBody {
	return s.Body
}

func (s *GetAssignJobsAsyncResultResponse) SetHeaders(v map[string]*string) *GetAssignJobsAsyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetAssignJobsAsyncResultResponse) SetStatusCode(v int32) *GetAssignJobsAsyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssignJobsAsyncResultResponse) SetBody(v *GetAssignJobsAsyncResultResponseBody) *GetAssignJobsAsyncResultResponse {
	s.Body = v
	return s
}

func (s *GetAssignJobsAsyncResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
