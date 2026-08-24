// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelVirusScanTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelVirusScanTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelVirusScanTasksResponse
	GetStatusCode() *int32
	SetBody(v *CancelVirusScanTasksResponseBody) *CancelVirusScanTasksResponse
	GetBody() *CancelVirusScanTasksResponseBody
}

type CancelVirusScanTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelVirusScanTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelVirusScanTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelVirusScanTasksResponse) GoString() string {
	return s.String()
}

func (s *CancelVirusScanTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelVirusScanTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelVirusScanTasksResponse) GetBody() *CancelVirusScanTasksResponseBody {
	return s.Body
}

func (s *CancelVirusScanTasksResponse) SetHeaders(v map[string]*string) *CancelVirusScanTasksResponse {
	s.Headers = v
	return s
}

func (s *CancelVirusScanTasksResponse) SetStatusCode(v int32) *CancelVirusScanTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelVirusScanTasksResponse) SetBody(v *CancelVirusScanTasksResponseBody) *CancelVirusScanTasksResponse {
	s.Body = v
	return s
}

func (s *CancelVirusScanTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
