// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketDeleteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketDeleteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketDeleteTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketDeleteTaskResponseBody) *GetBucketDeleteTaskResponse
	GetBody() *GetBucketDeleteTaskResponseBody
}

type GetBucketDeleteTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketDeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketDeleteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *GetBucketDeleteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketDeleteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketDeleteTaskResponse) GetBody() *GetBucketDeleteTaskResponseBody {
	return s.Body
}

func (s *GetBucketDeleteTaskResponse) SetHeaders(v map[string]*string) *GetBucketDeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *GetBucketDeleteTaskResponse) SetStatusCode(v int32) *GetBucketDeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketDeleteTaskResponse) SetBody(v *GetBucketDeleteTaskResponseBody) *GetBucketDeleteTaskResponse {
	s.Body = v
	return s
}

func (s *GetBucketDeleteTaskResponse) Validate() error {
	return dara.Validate(s)
}
