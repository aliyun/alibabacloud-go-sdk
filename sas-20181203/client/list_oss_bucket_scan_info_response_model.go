// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssBucketScanInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOssBucketScanInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOssBucketScanInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListOssBucketScanInfoResponseBody) *ListOssBucketScanInfoResponse
	GetBody() *ListOssBucketScanInfoResponseBody
}

type ListOssBucketScanInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOssBucketScanInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOssBucketScanInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOssBucketScanInfoResponse) GoString() string {
	return s.String()
}

func (s *ListOssBucketScanInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOssBucketScanInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOssBucketScanInfoResponse) GetBody() *ListOssBucketScanInfoResponseBody {
	return s.Body
}

func (s *ListOssBucketScanInfoResponse) SetHeaders(v map[string]*string) *ListOssBucketScanInfoResponse {
	s.Headers = v
	return s
}

func (s *ListOssBucketScanInfoResponse) SetStatusCode(v int32) *ListOssBucketScanInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOssBucketScanInfoResponse) SetBody(v *ListOssBucketScanInfoResponseBody) *ListOssBucketScanInfoResponse {
	s.Body = v
	return s
}

func (s *ListOssBucketScanInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
