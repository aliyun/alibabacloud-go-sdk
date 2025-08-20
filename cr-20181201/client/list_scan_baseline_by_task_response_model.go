// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanBaselineByTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScanBaselineByTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScanBaselineByTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListScanBaselineByTaskResponseBody) *ListScanBaselineByTaskResponse
	GetBody() *ListScanBaselineByTaskResponseBody
}

type ListScanBaselineByTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScanBaselineByTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScanBaselineByTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScanBaselineByTaskResponse) GoString() string {
	return s.String()
}

func (s *ListScanBaselineByTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScanBaselineByTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScanBaselineByTaskResponse) GetBody() *ListScanBaselineByTaskResponseBody {
	return s.Body
}

func (s *ListScanBaselineByTaskResponse) SetHeaders(v map[string]*string) *ListScanBaselineByTaskResponse {
	s.Headers = v
	return s
}

func (s *ListScanBaselineByTaskResponse) SetStatusCode(v int32) *ListScanBaselineByTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScanBaselineByTaskResponse) SetBody(v *ListScanBaselineByTaskResponseBody) *ListScanBaselineByTaskResponse {
	s.Body = v
	return s
}

func (s *ListScanBaselineByTaskResponse) Validate() error {
	return dara.Validate(s)
}
