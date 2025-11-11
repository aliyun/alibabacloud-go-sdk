// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeepWriteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeepWriteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeepWriteTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetDeepWriteTaskResponseBody) *GetDeepWriteTaskResponse
	GetBody() *GetDeepWriteTaskResponseBody
}

type GetDeepWriteTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeepWriteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeepWriteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeepWriteTaskResponse) GoString() string {
	return s.String()
}

func (s *GetDeepWriteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeepWriteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeepWriteTaskResponse) GetBody() *GetDeepWriteTaskResponseBody {
	return s.Body
}

func (s *GetDeepWriteTaskResponse) SetHeaders(v map[string]*string) *GetDeepWriteTaskResponse {
	s.Headers = v
	return s
}

func (s *GetDeepWriteTaskResponse) SetStatusCode(v int32) *GetDeepWriteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeepWriteTaskResponse) SetBody(v *GetDeepWriteTaskResponseBody) *GetDeepWriteTaskResponse {
	s.Body = v
	return s
}

func (s *GetDeepWriteTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
