// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssStorageAndAccByBucketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssStorageAndAccByBucketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssStorageAndAccByBucketsResponse
	GetStatusCode() *int32
	SetBody(v *GetOssStorageAndAccByBucketsResponseBody) *GetOssStorageAndAccByBucketsResponse
	GetBody() *GetOssStorageAndAccByBucketsResponseBody
}

type GetOssStorageAndAccByBucketsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssStorageAndAccByBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssStorageAndAccByBucketsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssStorageAndAccByBucketsResponse) GoString() string {
	return s.String()
}

func (s *GetOssStorageAndAccByBucketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssStorageAndAccByBucketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssStorageAndAccByBucketsResponse) GetBody() *GetOssStorageAndAccByBucketsResponseBody {
	return s.Body
}

func (s *GetOssStorageAndAccByBucketsResponse) SetHeaders(v map[string]*string) *GetOssStorageAndAccByBucketsResponse {
	s.Headers = v
	return s
}

func (s *GetOssStorageAndAccByBucketsResponse) SetStatusCode(v int32) *GetOssStorageAndAccByBucketsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssStorageAndAccByBucketsResponse) SetBody(v *GetOssStorageAndAccByBucketsResponseBody) *GetOssStorageAndAccByBucketsResponse {
	s.Body = v
	return s
}

func (s *GetOssStorageAndAccByBucketsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
