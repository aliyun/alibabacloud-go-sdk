// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshOssBucketScanInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshOssBucketScanInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshOssBucketScanInfoResponse
	GetStatusCode() *int32
	SetBody(v *RefreshOssBucketScanInfoResponseBody) *RefreshOssBucketScanInfoResponse
	GetBody() *RefreshOssBucketScanInfoResponseBody
}

type RefreshOssBucketScanInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshOssBucketScanInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshOssBucketScanInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshOssBucketScanInfoResponse) GoString() string {
	return s.String()
}

func (s *RefreshOssBucketScanInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshOssBucketScanInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshOssBucketScanInfoResponse) GetBody() *RefreshOssBucketScanInfoResponseBody {
	return s.Body
}

func (s *RefreshOssBucketScanInfoResponse) SetHeaders(v map[string]*string) *RefreshOssBucketScanInfoResponse {
	s.Headers = v
	return s
}

func (s *RefreshOssBucketScanInfoResponse) SetStatusCode(v int32) *RefreshOssBucketScanInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshOssBucketScanInfoResponse) SetBody(v *RefreshOssBucketScanInfoResponseBody) *RefreshOssBucketScanInfoResponse {
	s.Body = v
	return s
}

func (s *RefreshOssBucketScanInfoResponse) Validate() error {
	return dara.Validate(s)
}
