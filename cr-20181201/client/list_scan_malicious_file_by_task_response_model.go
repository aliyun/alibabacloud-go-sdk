// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanMaliciousFileByTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScanMaliciousFileByTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScanMaliciousFileByTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListScanMaliciousFileByTaskResponseBody) *ListScanMaliciousFileByTaskResponse
	GetBody() *ListScanMaliciousFileByTaskResponseBody
}

type ListScanMaliciousFileByTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScanMaliciousFileByTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScanMaliciousFileByTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScanMaliciousFileByTaskResponse) GoString() string {
	return s.String()
}

func (s *ListScanMaliciousFileByTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScanMaliciousFileByTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScanMaliciousFileByTaskResponse) GetBody() *ListScanMaliciousFileByTaskResponseBody {
	return s.Body
}

func (s *ListScanMaliciousFileByTaskResponse) SetHeaders(v map[string]*string) *ListScanMaliciousFileByTaskResponse {
	s.Headers = v
	return s
}

func (s *ListScanMaliciousFileByTaskResponse) SetStatusCode(v int32) *ListScanMaliciousFileByTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponse) SetBody(v *ListScanMaliciousFileByTaskResponseBody) *ListScanMaliciousFileByTaskResponse {
	s.Body = v
	return s
}

func (s *ListScanMaliciousFileByTaskResponse) Validate() error {
	return dara.Validate(s)
}
