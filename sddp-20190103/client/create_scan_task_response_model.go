// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateScanTaskResponseBody) *CreateScanTaskResponse
	GetBody() *CreateScanTaskResponseBody
}

type CreateScanTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScanTaskResponse) GetBody() *CreateScanTaskResponseBody {
	return s.Body
}

func (s *CreateScanTaskResponse) SetHeaders(v map[string]*string) *CreateScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateScanTaskResponse) SetStatusCode(v int32) *CreateScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScanTaskResponse) SetBody(v *CreateScanTaskResponseBody) *CreateScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateScanTaskResponse) Validate() error {
	return dara.Validate(s)
}
