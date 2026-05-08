// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectTaskResponseBody) *GetProjectTaskResponse
	GetBody() *GetProjectTaskResponseBody
}

type GetProjectTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *GetProjectTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectTaskResponse) GetBody() *GetProjectTaskResponseBody {
	return s.Body
}

func (s *GetProjectTaskResponse) SetHeaders(v map[string]*string) *GetProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *GetProjectTaskResponse) SetStatusCode(v int32) *GetProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectTaskResponse) SetBody(v *GetProjectTaskResponseBody) *GetProjectTaskResponse {
	s.Body = v
	return s
}

func (s *GetProjectTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
