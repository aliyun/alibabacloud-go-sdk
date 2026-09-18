// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTargetScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTargetScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTargetScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateTargetScanTaskResponseBody) *CreateTargetScanTaskResponse
	GetBody() *CreateTargetScanTaskResponseBody
}

type CreateTargetScanTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTargetScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTargetScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTargetScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTargetScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTargetScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTargetScanTaskResponse) GetBody() *CreateTargetScanTaskResponseBody {
	return s.Body
}

func (s *CreateTargetScanTaskResponse) SetHeaders(v map[string]*string) *CreateTargetScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTargetScanTaskResponse) SetStatusCode(v int32) *CreateTargetScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTargetScanTaskResponse) SetBody(v *CreateTargetScanTaskResponseBody) *CreateTargetScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTargetScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
