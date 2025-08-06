// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateBucketDeleteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateBucketDeleteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateBucketDeleteTaskResponse
	GetStatusCode() *int32
	SetBody(v *TerminateBucketDeleteTaskResponseBody) *TerminateBucketDeleteTaskResponse
	GetBody() *TerminateBucketDeleteTaskResponseBody
}

type TerminateBucketDeleteTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateBucketDeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateBucketDeleteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateBucketDeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *TerminateBucketDeleteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateBucketDeleteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateBucketDeleteTaskResponse) GetBody() *TerminateBucketDeleteTaskResponseBody {
	return s.Body
}

func (s *TerminateBucketDeleteTaskResponse) SetHeaders(v map[string]*string) *TerminateBucketDeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *TerminateBucketDeleteTaskResponse) SetStatusCode(v int32) *TerminateBucketDeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateBucketDeleteTaskResponse) SetBody(v *TerminateBucketDeleteTaskResponseBody) *TerminateBucketDeleteTaskResponse {
	s.Body = v
	return s
}

func (s *TerminateBucketDeleteTaskResponse) Validate() error {
	return dara.Validate(s)
}
