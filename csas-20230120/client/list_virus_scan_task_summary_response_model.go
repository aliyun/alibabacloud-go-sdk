// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTaskSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirusScanTaskSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirusScanTaskSummaryResponse
	GetStatusCode() *int32
	SetBody(v *ListVirusScanTaskSummaryResponseBody) *ListVirusScanTaskSummaryResponse
	GetBody() *ListVirusScanTaskSummaryResponseBody
}

type ListVirusScanTaskSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirusScanTaskSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirusScanTaskSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskSummaryResponse) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirusScanTaskSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirusScanTaskSummaryResponse) GetBody() *ListVirusScanTaskSummaryResponseBody {
	return s.Body
}

func (s *ListVirusScanTaskSummaryResponse) SetHeaders(v map[string]*string) *ListVirusScanTaskSummaryResponse {
	s.Headers = v
	return s
}

func (s *ListVirusScanTaskSummaryResponse) SetStatusCode(v int32) *ListVirusScanTaskSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirusScanTaskSummaryResponse) SetBody(v *ListVirusScanTaskSummaryResponseBody) *ListVirusScanTaskSummaryResponse {
	s.Body = v
	return s
}

func (s *ListVirusScanTaskSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
