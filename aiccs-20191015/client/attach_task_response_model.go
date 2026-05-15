// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachTaskResponse
	GetStatusCode() *int32
	SetBody(v *AttachTaskResponseBody) *AttachTaskResponse
	GetBody() *AttachTaskResponseBody
}

type AttachTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachTaskResponse) GoString() string {
	return s.String()
}

func (s *AttachTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachTaskResponse) GetBody() *AttachTaskResponseBody {
	return s.Body
}

func (s *AttachTaskResponse) SetHeaders(v map[string]*string) *AttachTaskResponse {
	s.Headers = v
	return s
}

func (s *AttachTaskResponse) SetStatusCode(v int32) *AttachTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachTaskResponse) SetBody(v *AttachTaskResponseBody) *AttachTaskResponse {
	s.Body = v
	return s
}

func (s *AttachTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
