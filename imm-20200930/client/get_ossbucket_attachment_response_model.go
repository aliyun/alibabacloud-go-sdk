// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSBucketAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOSSBucketAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOSSBucketAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *GetOSSBucketAttachmentResponseBody) *GetOSSBucketAttachmentResponse
	GetBody() *GetOSSBucketAttachmentResponseBody
}

type GetOSSBucketAttachmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOSSBucketAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOSSBucketAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOSSBucketAttachmentResponse) GoString() string {
	return s.String()
}

func (s *GetOSSBucketAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOSSBucketAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOSSBucketAttachmentResponse) GetBody() *GetOSSBucketAttachmentResponseBody {
	return s.Body
}

func (s *GetOSSBucketAttachmentResponse) SetHeaders(v map[string]*string) *GetOSSBucketAttachmentResponse {
	s.Headers = v
	return s
}

func (s *GetOSSBucketAttachmentResponse) SetStatusCode(v int32) *GetOSSBucketAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSBucketAttachmentResponse) SetBody(v *GetOSSBucketAttachmentResponseBody) *GetOSSBucketAttachmentResponse {
	s.Body = v
	return s
}

func (s *GetOSSBucketAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
