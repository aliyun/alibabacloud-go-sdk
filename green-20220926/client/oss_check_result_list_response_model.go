// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOssCheckResultListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OssCheckResultListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OssCheckResultListResponse
	GetStatusCode() *int32
	SetBody(v *OssCheckResultListResponseBody) *OssCheckResultListResponse
	GetBody() *OssCheckResultListResponseBody
}

type OssCheckResultListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OssCheckResultListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OssCheckResultListResponse) String() string {
	return dara.Prettify(s)
}

func (s OssCheckResultListResponse) GoString() string {
	return s.String()
}

func (s *OssCheckResultListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OssCheckResultListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OssCheckResultListResponse) GetBody() *OssCheckResultListResponseBody {
	return s.Body
}

func (s *OssCheckResultListResponse) SetHeaders(v map[string]*string) *OssCheckResultListResponse {
	s.Headers = v
	return s
}

func (s *OssCheckResultListResponse) SetStatusCode(v int32) *OssCheckResultListResponse {
	s.StatusCode = &v
	return s
}

func (s *OssCheckResultListResponse) SetBody(v *OssCheckResultListResponseBody) *OssCheckResultListResponse {
	s.Body = v
	return s
}

func (s *OssCheckResultListResponse) Validate() error {
	return dara.Validate(s)
}
