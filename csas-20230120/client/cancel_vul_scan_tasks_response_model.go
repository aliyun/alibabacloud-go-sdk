// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelVulScanTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelVulScanTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelVulScanTasksResponse
	GetStatusCode() *int32
	SetBody(v *CancelVulScanTasksResponseBody) *CancelVulScanTasksResponse
	GetBody() *CancelVulScanTasksResponseBody
}

type CancelVulScanTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelVulScanTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelVulScanTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelVulScanTasksResponse) GoString() string {
	return s.String()
}

func (s *CancelVulScanTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelVulScanTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelVulScanTasksResponse) GetBody() *CancelVulScanTasksResponseBody {
	return s.Body
}

func (s *CancelVulScanTasksResponse) SetHeaders(v map[string]*string) *CancelVulScanTasksResponse {
	s.Headers = v
	return s
}

func (s *CancelVulScanTasksResponse) SetStatusCode(v int32) *CancelVulScanTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelVulScanTasksResponse) SetBody(v *CancelVulScanTasksResponseBody) *CancelVulScanTasksResponse {
	s.Body = v
	return s
}

func (s *CancelVulScanTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
