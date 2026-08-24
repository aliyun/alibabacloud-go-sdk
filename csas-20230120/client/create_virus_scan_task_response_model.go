// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVirusScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVirusScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVirusScanTaskResponseBody) *CreateVirusScanTaskResponse
	GetBody() *CreateVirusScanTaskResponseBody
}

type CreateVirusScanTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirusScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirusScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVirusScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVirusScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVirusScanTaskResponse) GetBody() *CreateVirusScanTaskResponseBody {
	return s.Body
}

func (s *CreateVirusScanTaskResponse) SetHeaders(v map[string]*string) *CreateVirusScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVirusScanTaskResponse) SetStatusCode(v int32) *CreateVirusScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirusScanTaskResponse) SetBody(v *CreateVirusScanTaskResponseBody) *CreateVirusScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVirusScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
