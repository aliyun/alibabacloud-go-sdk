// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketResponse
	GetStatusCode() *int32
	SetBody(v *PutBucketResponseBody) *PutBucketResponse
	GetBody() *PutBucketResponseBody
}

type PutBucketResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketResponse) GoString() string {
	return s.String()
}

func (s *PutBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketResponse) GetBody() *PutBucketResponseBody {
	return s.Body
}

func (s *PutBucketResponse) SetHeaders(v map[string]*string) *PutBucketResponse {
	s.Headers = v
	return s
}

func (s *PutBucketResponse) SetStatusCode(v int32) *PutBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketResponse) SetBody(v *PutBucketResponseBody) *PutBucketResponse {
	s.Body = v
	return s
}

func (s *PutBucketResponse) Validate() error {
	return dara.Validate(s)
}
