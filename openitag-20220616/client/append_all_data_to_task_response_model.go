// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendAllDataToTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AppendAllDataToTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AppendAllDataToTaskResponse
	GetStatusCode() *int32
	SetBody(v *AppendAllDataToTaskResponseBody) *AppendAllDataToTaskResponse
	GetBody() *AppendAllDataToTaskResponseBody
}

type AppendAllDataToTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppendAllDataToTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendAllDataToTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AppendAllDataToTaskResponse) GoString() string {
	return s.String()
}

func (s *AppendAllDataToTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AppendAllDataToTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AppendAllDataToTaskResponse) GetBody() *AppendAllDataToTaskResponseBody {
	return s.Body
}

func (s *AppendAllDataToTaskResponse) SetHeaders(v map[string]*string) *AppendAllDataToTaskResponse {
	s.Headers = v
	return s
}

func (s *AppendAllDataToTaskResponse) SetStatusCode(v int32) *AppendAllDataToTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AppendAllDataToTaskResponse) SetBody(v *AppendAllDataToTaskResponseBody) *AppendAllDataToTaskResponse {
	s.Body = v
	return s
}

func (s *AppendAllDataToTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
