// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContainerScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContainerScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContainerScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateContainerScanTaskResponseBody) *CreateContainerScanTaskResponse
	GetBody() *CreateContainerScanTaskResponseBody
}

type CreateContainerScanTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContainerScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContainerScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContainerScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateContainerScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContainerScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContainerScanTaskResponse) GetBody() *CreateContainerScanTaskResponseBody {
	return s.Body
}

func (s *CreateContainerScanTaskResponse) SetHeaders(v map[string]*string) *CreateContainerScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateContainerScanTaskResponse) SetStatusCode(v int32) *CreateContainerScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContainerScanTaskResponse) SetBody(v *CreateContainerScanTaskResponseBody) *CreateContainerScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateContainerScanTaskResponse) Validate() error {
	return dara.Validate(s)
}
