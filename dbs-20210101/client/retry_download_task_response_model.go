// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDownloadTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryDownloadTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryDownloadTaskResponse
	GetStatusCode() *int32
	SetBody(v *RetryDownloadTaskResponseBody) *RetryDownloadTaskResponse
	GetBody() *RetryDownloadTaskResponseBody
}

type RetryDownloadTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryDownloadTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryDownloadTaskResponse) GoString() string {
	return s.String()
}

func (s *RetryDownloadTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryDownloadTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryDownloadTaskResponse) GetBody() *RetryDownloadTaskResponseBody {
	return s.Body
}

func (s *RetryDownloadTaskResponse) SetHeaders(v map[string]*string) *RetryDownloadTaskResponse {
	s.Headers = v
	return s
}

func (s *RetryDownloadTaskResponse) SetStatusCode(v int32) *RetryDownloadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryDownloadTaskResponse) SetBody(v *RetryDownloadTaskResponseBody) *RetryDownloadTaskResponse {
	s.Body = v
	return s
}

func (s *RetryDownloadTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
