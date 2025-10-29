// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInstanceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskInstanceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskInstanceLogResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskInstanceLogResponseBody) *GetTaskInstanceLogResponse
	GetBody() *GetTaskInstanceLogResponseBody
}

type GetTaskInstanceLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskInstanceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskInstanceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceLogResponse) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskInstanceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskInstanceLogResponse) GetBody() *GetTaskInstanceLogResponseBody {
	return s.Body
}

func (s *GetTaskInstanceLogResponse) SetHeaders(v map[string]*string) *GetTaskInstanceLogResponse {
	s.Headers = v
	return s
}

func (s *GetTaskInstanceLogResponse) SetStatusCode(v int32) *GetTaskInstanceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskInstanceLogResponse) SetBody(v *GetTaskInstanceLogResponseBody) *GetTaskInstanceLogResponse {
	s.Body = v
	return s
}

func (s *GetTaskInstanceLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
