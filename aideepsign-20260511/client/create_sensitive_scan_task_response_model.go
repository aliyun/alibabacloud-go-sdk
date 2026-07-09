// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSensitiveScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSensitiveScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSensitiveScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSensitiveScanTaskResponseBody) *CreateSensitiveScanTaskResponse
	GetBody() *CreateSensitiveScanTaskResponseBody
}

type CreateSensitiveScanTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSensitiveScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSensitiveScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSensitiveScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSensitiveScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSensitiveScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSensitiveScanTaskResponse) GetBody() *CreateSensitiveScanTaskResponseBody {
	return s.Body
}

func (s *CreateSensitiveScanTaskResponse) SetHeaders(v map[string]*string) *CreateSensitiveScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSensitiveScanTaskResponse) SetStatusCode(v int32) *CreateSensitiveScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSensitiveScanTaskResponse) SetBody(v *CreateSensitiveScanTaskResponseBody) *CreateSensitiveScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSensitiveScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
