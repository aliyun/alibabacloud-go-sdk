// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanOnceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVirusScanOnceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVirusScanOnceTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVirusScanOnceTaskResponseBody) *CreateVirusScanOnceTaskResponse
	GetBody() *CreateVirusScanOnceTaskResponseBody
}

type CreateVirusScanOnceTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirusScanOnceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirusScanOnceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanOnceTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVirusScanOnceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVirusScanOnceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVirusScanOnceTaskResponse) GetBody() *CreateVirusScanOnceTaskResponseBody {
	return s.Body
}

func (s *CreateVirusScanOnceTaskResponse) SetHeaders(v map[string]*string) *CreateVirusScanOnceTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVirusScanOnceTaskResponse) SetStatusCode(v int32) *CreateVirusScanOnceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponse) SetBody(v *CreateVirusScanOnceTaskResponseBody) *CreateVirusScanOnceTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVirusScanOnceTaskResponse) Validate() error {
	return dara.Validate(s)
}
