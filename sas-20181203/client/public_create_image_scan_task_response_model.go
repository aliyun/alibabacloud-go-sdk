// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicCreateImageScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublicCreateImageScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublicCreateImageScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *PublicCreateImageScanTaskResponseBody) *PublicCreateImageScanTaskResponse
	GetBody() *PublicCreateImageScanTaskResponseBody
}

type PublicCreateImageScanTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublicCreateImageScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublicCreateImageScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PublicCreateImageScanTaskResponse) GoString() string {
	return s.String()
}

func (s *PublicCreateImageScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublicCreateImageScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublicCreateImageScanTaskResponse) GetBody() *PublicCreateImageScanTaskResponseBody {
	return s.Body
}

func (s *PublicCreateImageScanTaskResponse) SetHeaders(v map[string]*string) *PublicCreateImageScanTaskResponse {
	s.Headers = v
	return s
}

func (s *PublicCreateImageScanTaskResponse) SetStatusCode(v int32) *PublicCreateImageScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PublicCreateImageScanTaskResponse) SetBody(v *PublicCreateImageScanTaskResponseBody) *PublicCreateImageScanTaskResponse {
	s.Body = v
	return s
}

func (s *PublicCreateImageScanTaskResponse) Validate() error {
	return dara.Validate(s)
}
