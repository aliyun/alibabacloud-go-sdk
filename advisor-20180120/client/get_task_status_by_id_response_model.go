// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskStatusByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskStatusByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskStatusByIdResponseBody) *GetTaskStatusByIdResponse
	GetBody() *GetTaskStatusByIdResponseBody
}

type GetTaskStatusByIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatusByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatusByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusByIdResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskStatusByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskStatusByIdResponse) GetBody() *GetTaskStatusByIdResponseBody {
	return s.Body
}

func (s *GetTaskStatusByIdResponse) SetHeaders(v map[string]*string) *GetTaskStatusByIdResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusByIdResponse) SetStatusCode(v int32) *GetTaskStatusByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatusByIdResponse) SetBody(v *GetTaskStatusByIdResponseBody) *GetTaskStatusByIdResponse {
	s.Body = v
	return s
}

func (s *GetTaskStatusByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
