// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartVirusScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartVirusScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartVirusScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartVirusScanTaskResponseBody) *StartVirusScanTaskResponse
	GetBody() *StartVirusScanTaskResponseBody
}

type StartVirusScanTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartVirusScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartVirusScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartVirusScanTaskResponse) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartVirusScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartVirusScanTaskResponse) GetBody() *StartVirusScanTaskResponseBody {
	return s.Body
}

func (s *StartVirusScanTaskResponse) SetHeaders(v map[string]*string) *StartVirusScanTaskResponse {
	s.Headers = v
	return s
}

func (s *StartVirusScanTaskResponse) SetStatusCode(v int32) *StartVirusScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartVirusScanTaskResponse) SetBody(v *StartVirusScanTaskResponseBody) *StartVirusScanTaskResponse {
	s.Body = v
	return s
}

func (s *StartVirusScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
