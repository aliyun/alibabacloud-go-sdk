// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBucketDeleteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitBucketDeleteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitBucketDeleteTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitBucketDeleteTaskResponseBody) *SubmitBucketDeleteTaskResponse
	GetBody() *SubmitBucketDeleteTaskResponseBody
}

type SubmitBucketDeleteTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitBucketDeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBucketDeleteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitBucketDeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitBucketDeleteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitBucketDeleteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitBucketDeleteTaskResponse) GetBody() *SubmitBucketDeleteTaskResponseBody {
	return s.Body
}

func (s *SubmitBucketDeleteTaskResponse) SetHeaders(v map[string]*string) *SubmitBucketDeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitBucketDeleteTaskResponse) SetStatusCode(v int32) *SubmitBucketDeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBucketDeleteTaskResponse) SetBody(v *SubmitBucketDeleteTaskResponseBody) *SubmitBucketDeleteTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitBucketDeleteTaskResponse) Validate() error {
	return dara.Validate(s)
}
