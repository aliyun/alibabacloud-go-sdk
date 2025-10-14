// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketAclResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketAclResponseBody) *GetBucketAclResponse
	GetBody() *GetBucketAclResponseBody
}

type GetBucketAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketAclResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketAclResponse) GoString() string {
	return s.String()
}

func (s *GetBucketAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketAclResponse) GetBody() *GetBucketAclResponseBody {
	return s.Body
}

func (s *GetBucketAclResponse) SetHeaders(v map[string]*string) *GetBucketAclResponse {
	s.Headers = v
	return s
}

func (s *GetBucketAclResponse) SetStatusCode(v int32) *GetBucketAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketAclResponse) SetBody(v *GetBucketAclResponseBody) *GetBucketAclResponse {
	s.Body = v
	return s
}

func (s *GetBucketAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
