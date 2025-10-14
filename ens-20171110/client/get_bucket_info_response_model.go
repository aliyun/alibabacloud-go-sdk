// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketInfoResponseBody) *GetBucketInfoResponse
	GetBody() *GetBucketInfoResponseBody
}

type GetBucketInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketInfoResponse) GetBody() *GetBucketInfoResponseBody {
	return s.Body
}

func (s *GetBucketInfoResponse) SetHeaders(v map[string]*string) *GetBucketInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBucketInfoResponse) SetStatusCode(v int32) *GetBucketInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketInfoResponse) SetBody(v *GetBucketInfoResponseBody) *GetBucketInfoResponse {
	s.Body = v
	return s
}

func (s *GetBucketInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
