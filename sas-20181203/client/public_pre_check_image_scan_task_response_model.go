// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicPreCheckImageScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublicPreCheckImageScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublicPreCheckImageScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *PublicPreCheckImageScanTaskResponseBody) *PublicPreCheckImageScanTaskResponse
	GetBody() *PublicPreCheckImageScanTaskResponseBody
}

type PublicPreCheckImageScanTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublicPreCheckImageScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublicPreCheckImageScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PublicPreCheckImageScanTaskResponse) GoString() string {
	return s.String()
}

func (s *PublicPreCheckImageScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublicPreCheckImageScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublicPreCheckImageScanTaskResponse) GetBody() *PublicPreCheckImageScanTaskResponseBody {
	return s.Body
}

func (s *PublicPreCheckImageScanTaskResponse) SetHeaders(v map[string]*string) *PublicPreCheckImageScanTaskResponse {
	s.Headers = v
	return s
}

func (s *PublicPreCheckImageScanTaskResponse) SetStatusCode(v int32) *PublicPreCheckImageScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PublicPreCheckImageScanTaskResponse) SetBody(v *PublicPreCheckImageScanTaskResponseBody) *PublicPreCheckImageScanTaskResponse {
	s.Body = v
	return s
}

func (s *PublicPreCheckImageScanTaskResponse) Validate() error {
	return dara.Validate(s)
}
