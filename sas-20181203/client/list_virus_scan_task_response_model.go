// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirusScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirusScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListVirusScanTaskResponseBody) *ListVirusScanTaskResponse
	GetBody() *ListVirusScanTaskResponseBody
}

type ListVirusScanTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirusScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirusScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskResponse) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirusScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirusScanTaskResponse) GetBody() *ListVirusScanTaskResponseBody {
	return s.Body
}

func (s *ListVirusScanTaskResponse) SetHeaders(v map[string]*string) *ListVirusScanTaskResponse {
	s.Headers = v
	return s
}

func (s *ListVirusScanTaskResponse) SetStatusCode(v int32) *ListVirusScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirusScanTaskResponse) SetBody(v *ListVirusScanTaskResponseBody) *ListVirusScanTaskResponse {
	s.Body = v
	return s
}

func (s *ListVirusScanTaskResponse) Validate() error {
	return dara.Validate(s)
}
