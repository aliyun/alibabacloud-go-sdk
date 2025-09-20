// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachOSSBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachOSSBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachOSSBucketResponse
	GetStatusCode() *int32
	SetBody(v *DetachOSSBucketResponseBody) *DetachOSSBucketResponse
	GetBody() *DetachOSSBucketResponseBody
}

type DetachOSSBucketResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachOSSBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachOSSBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachOSSBucketResponse) GoString() string {
	return s.String()
}

func (s *DetachOSSBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachOSSBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachOSSBucketResponse) GetBody() *DetachOSSBucketResponseBody {
	return s.Body
}

func (s *DetachOSSBucketResponse) SetHeaders(v map[string]*string) *DetachOSSBucketResponse {
	s.Headers = v
	return s
}

func (s *DetachOSSBucketResponse) SetStatusCode(v int32) *DetachOSSBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachOSSBucketResponse) SetBody(v *DetachOSSBucketResponseBody) *DetachOSSBucketResponse {
	s.Body = v
	return s
}

func (s *DetachOSSBucketResponse) Validate() error {
	return dara.Validate(s)
}
