// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVulScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVulScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVulScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVulScanTaskResponseBody) *CreateVulScanTaskResponse
	GetBody() *CreateVulScanTaskResponseBody
}

type CreateVulScanTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVulScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVulScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVulScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVulScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVulScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVulScanTaskResponse) GetBody() *CreateVulScanTaskResponseBody {
	return s.Body
}

func (s *CreateVulScanTaskResponse) SetHeaders(v map[string]*string) *CreateVulScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVulScanTaskResponse) SetStatusCode(v int32) *CreateVulScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVulScanTaskResponse) SetBody(v *CreateVulScanTaskResponseBody) *CreateVulScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVulScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
