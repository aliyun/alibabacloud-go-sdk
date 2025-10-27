// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOnceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateOnceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateOnceTaskResponse
	GetStatusCode() *int32
	SetBody(v *GenerateOnceTaskResponseBody) *GenerateOnceTaskResponse
	GetBody() *GenerateOnceTaskResponseBody
}

type GenerateOnceTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateOnceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateOnceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateOnceTaskResponse) GoString() string {
	return s.String()
}

func (s *GenerateOnceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateOnceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateOnceTaskResponse) GetBody() *GenerateOnceTaskResponseBody {
	return s.Body
}

func (s *GenerateOnceTaskResponse) SetHeaders(v map[string]*string) *GenerateOnceTaskResponse {
	s.Headers = v
	return s
}

func (s *GenerateOnceTaskResponse) SetStatusCode(v int32) *GenerateOnceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateOnceTaskResponse) SetBody(v *GenerateOnceTaskResponseBody) *GenerateOnceTaskResponse {
	s.Body = v
	return s
}

func (s *GenerateOnceTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
