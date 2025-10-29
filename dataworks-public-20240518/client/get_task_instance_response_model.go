// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskInstanceResponseBody) *GetTaskInstanceResponse
	GetBody() *GetTaskInstanceResponseBody
}

type GetTaskInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskInstanceResponse) GetBody() *GetTaskInstanceResponseBody {
	return s.Body
}

func (s *GetTaskInstanceResponse) SetHeaders(v map[string]*string) *GetTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetTaskInstanceResponse) SetStatusCode(v int32) *GetTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskInstanceResponse) SetBody(v *GetTaskInstanceResponseBody) *GetTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *GetTaskInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
