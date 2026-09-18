// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanTasksByTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScanTasksByTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScanTasksByTargetResponse
	GetStatusCode() *int32
	SetBody(v *ListScanTasksByTargetResponseBody) *ListScanTasksByTargetResponse
	GetBody() *ListScanTasksByTargetResponseBody
}

type ListScanTasksByTargetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScanTasksByTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScanTasksByTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScanTasksByTargetResponse) GoString() string {
	return s.String()
}

func (s *ListScanTasksByTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScanTasksByTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScanTasksByTargetResponse) GetBody() *ListScanTasksByTargetResponseBody {
	return s.Body
}

func (s *ListScanTasksByTargetResponse) SetHeaders(v map[string]*string) *ListScanTasksByTargetResponse {
	s.Headers = v
	return s
}

func (s *ListScanTasksByTargetResponse) SetStatusCode(v int32) *ListScanTasksByTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScanTasksByTargetResponse) SetBody(v *ListScanTasksByTargetResponseBody) *ListScanTasksByTargetResponse {
	s.Body = v
	return s
}

func (s *ListScanTasksByTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
