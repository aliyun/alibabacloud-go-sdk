// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTaskStatusesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirusScanTaskStatusesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirusScanTaskStatusesResponse
	GetStatusCode() *int32
	SetBody(v *ListVirusScanTaskStatusesResponseBody) *ListVirusScanTaskStatusesResponse
	GetBody() *ListVirusScanTaskStatusesResponseBody
}

type ListVirusScanTaskStatusesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirusScanTaskStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirusScanTaskStatusesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskStatusesResponse) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskStatusesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirusScanTaskStatusesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirusScanTaskStatusesResponse) GetBody() *ListVirusScanTaskStatusesResponseBody {
	return s.Body
}

func (s *ListVirusScanTaskStatusesResponse) SetHeaders(v map[string]*string) *ListVirusScanTaskStatusesResponse {
	s.Headers = v
	return s
}

func (s *ListVirusScanTaskStatusesResponse) SetStatusCode(v int32) *ListVirusScanTaskStatusesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirusScanTaskStatusesResponse) SetBody(v *ListVirusScanTaskStatusesResponseBody) *ListVirusScanTaskStatusesResponse {
	s.Body = v
	return s
}

func (s *ListVirusScanTaskStatusesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
