// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyScheduledTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyScheduledTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyScheduledTaskResponseBody) *ModifyScheduledTaskResponse
	GetBody() *ModifyScheduledTaskResponseBody
}

type ModifyScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScheduledTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyScheduledTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyScheduledTaskResponse) GetBody() *ModifyScheduledTaskResponseBody {
	return s.Body
}

func (s *ModifyScheduledTaskResponse) SetHeaders(v map[string]*string) *ModifyScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduledTaskResponse) SetStatusCode(v int32) *ModifyScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScheduledTaskResponse) SetBody(v *ModifyScheduledTaskResponseBody) *ModifyScheduledTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyScheduledTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
