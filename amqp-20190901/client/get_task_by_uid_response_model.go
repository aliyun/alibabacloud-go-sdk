// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskByUidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskByUidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskByUidResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskByUidResponseBody) *GetTaskByUidResponse
	GetBody() *GetTaskByUidResponseBody
}

type GetTaskByUidResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskByUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskByUidResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUidResponse) GoString() string {
	return s.String()
}

func (s *GetTaskByUidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskByUidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskByUidResponse) GetBody() *GetTaskByUidResponseBody {
	return s.Body
}

func (s *GetTaskByUidResponse) SetHeaders(v map[string]*string) *GetTaskByUidResponse {
	s.Headers = v
	return s
}

func (s *GetTaskByUidResponse) SetStatusCode(v int32) *GetTaskByUidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskByUidResponse) SetBody(v *GetTaskByUidResponseBody) *GetTaskByUidResponse {
	s.Body = v
	return s
}

func (s *GetTaskByUidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
