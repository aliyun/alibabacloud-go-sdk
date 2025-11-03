// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoTagScanResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepoTagScanResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepoTagScanResultResponse
	GetStatusCode() *int32
	SetBody(v *ListRepoTagScanResultResponseBody) *ListRepoTagScanResultResponse
	GetBody() *ListRepoTagScanResultResponseBody
}

type ListRepoTagScanResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoTagScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoTagScanResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTagScanResultResponse) GoString() string {
	return s.String()
}

func (s *ListRepoTagScanResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepoTagScanResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepoTagScanResultResponse) GetBody() *ListRepoTagScanResultResponseBody {
	return s.Body
}

func (s *ListRepoTagScanResultResponse) SetHeaders(v map[string]*string) *ListRepoTagScanResultResponse {
	s.Headers = v
	return s
}

func (s *ListRepoTagScanResultResponse) SetStatusCode(v int32) *ListRepoTagScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoTagScanResultResponse) SetBody(v *ListRepoTagScanResultResponseBody) *ListRepoTagScanResultResponse {
	s.Body = v
	return s
}

func (s *ListRepoTagScanResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
