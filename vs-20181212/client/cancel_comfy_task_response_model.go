// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelComfyTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelComfyTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelComfyTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelComfyTaskResponseBody) *CancelComfyTaskResponse
	GetBody() *CancelComfyTaskResponseBody
}

type CancelComfyTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelComfyTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelComfyTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelComfyTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelComfyTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelComfyTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelComfyTaskResponse) GetBody() *CancelComfyTaskResponseBody {
	return s.Body
}

func (s *CancelComfyTaskResponse) SetHeaders(v map[string]*string) *CancelComfyTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelComfyTaskResponse) SetStatusCode(v int32) *CancelComfyTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelComfyTaskResponse) SetBody(v *CancelComfyTaskResponseBody) *CancelComfyTaskResponse {
	s.Body = v
	return s
}

func (s *CancelComfyTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
