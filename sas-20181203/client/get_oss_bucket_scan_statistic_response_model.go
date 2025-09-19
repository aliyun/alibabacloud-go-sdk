// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssBucketScanStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssBucketScanStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssBucketScanStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetOssBucketScanStatisticResponseBody) *GetOssBucketScanStatisticResponse
	GetBody() *GetOssBucketScanStatisticResponseBody
}

type GetOssBucketScanStatisticResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssBucketScanStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssBucketScanStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssBucketScanStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetOssBucketScanStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssBucketScanStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssBucketScanStatisticResponse) GetBody() *GetOssBucketScanStatisticResponseBody {
	return s.Body
}

func (s *GetOssBucketScanStatisticResponse) SetHeaders(v map[string]*string) *GetOssBucketScanStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetOssBucketScanStatisticResponse) SetStatusCode(v int32) *GetOssBucketScanStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssBucketScanStatisticResponse) SetBody(v *GetOssBucketScanStatisticResponseBody) *GetOssBucketScanStatisticResponse {
	s.Body = v
	return s
}

func (s *GetOssBucketScanStatisticResponse) Validate() error {
	return dara.Validate(s)
}
