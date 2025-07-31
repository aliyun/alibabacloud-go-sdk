// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOssCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOssCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *ListOssCheckResultResponseBody) *ListOssCheckResultResponse
	GetBody() *ListOssCheckResultResponseBody
}

type ListOssCheckResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOssCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOssCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOssCheckResultResponse) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOssCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOssCheckResultResponse) GetBody() *ListOssCheckResultResponseBody {
	return s.Body
}

func (s *ListOssCheckResultResponse) SetHeaders(v map[string]*string) *ListOssCheckResultResponse {
	s.Headers = v
	return s
}

func (s *ListOssCheckResultResponse) SetStatusCode(v int32) *ListOssCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOssCheckResultResponse) SetBody(v *ListOssCheckResultResponseBody) *ListOssCheckResultResponse {
	s.Body = v
	return s
}

func (s *ListOssCheckResultResponse) Validate() error {
	return dara.Validate(s)
}
