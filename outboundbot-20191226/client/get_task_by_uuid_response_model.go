// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskByUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskByUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskByUuidResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskByUuidResponseBody) *GetTaskByUuidResponse
	GetBody() *GetTaskByUuidResponseBody
}

type GetTaskByUuidResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskByUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskByUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUuidResponse) GoString() string {
	return s.String()
}

func (s *GetTaskByUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskByUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskByUuidResponse) GetBody() *GetTaskByUuidResponseBody {
	return s.Body
}

func (s *GetTaskByUuidResponse) SetHeaders(v map[string]*string) *GetTaskByUuidResponse {
	s.Headers = v
	return s
}

func (s *GetTaskByUuidResponse) SetStatusCode(v int32) *GetTaskByUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskByUuidResponse) SetBody(v *GetTaskByUuidResponseBody) *GetTaskByUuidResponse {
	s.Body = v
	return s
}

func (s *GetTaskByUuidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
