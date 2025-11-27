// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachOSSBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachOSSBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachOSSBucketResponse
	GetStatusCode() *int32
	SetBody(v *AttachOSSBucketResponseBody) *AttachOSSBucketResponse
	GetBody() *AttachOSSBucketResponseBody
}

type AttachOSSBucketResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachOSSBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachOSSBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachOSSBucketResponse) GoString() string {
	return s.String()
}

func (s *AttachOSSBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachOSSBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachOSSBucketResponse) GetBody() *AttachOSSBucketResponseBody {
	return s.Body
}

func (s *AttachOSSBucketResponse) SetHeaders(v map[string]*string) *AttachOSSBucketResponse {
	s.Headers = v
	return s
}

func (s *AttachOSSBucketResponse) SetStatusCode(v int32) *AttachOSSBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachOSSBucketResponse) SetBody(v *AttachOSSBucketResponseBody) *AttachOSSBucketResponse {
	s.Body = v
	return s
}

func (s *AttachOSSBucketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
