// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOtaTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApproveOtaTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApproveOtaTaskResponse
	GetStatusCode() *int32
	SetBody(v *ApproveOtaTaskResponseBody) *ApproveOtaTaskResponse
	GetBody() *ApproveOtaTaskResponseBody
}

type ApproveOtaTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveOtaTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ApproveOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApproveOtaTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApproveOtaTaskResponse) GetBody() *ApproveOtaTaskResponseBody {
	return s.Body
}

func (s *ApproveOtaTaskResponse) SetHeaders(v map[string]*string) *ApproveOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *ApproveOtaTaskResponse) SetStatusCode(v int32) *ApproveOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveOtaTaskResponse) SetBody(v *ApproveOtaTaskResponseBody) *ApproveOtaTaskResponse {
	s.Body = v
	return s
}

func (s *ApproveOtaTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
