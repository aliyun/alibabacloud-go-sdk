// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduleExecuteTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyScheduleExecuteTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyScheduleExecuteTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyScheduleExecuteTimeResponseBody) *ModifyScheduleExecuteTimeResponse
	GetBody() *ModifyScheduleExecuteTimeResponseBody
}

type ModifyScheduleExecuteTimeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScheduleExecuteTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScheduleExecuteTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduleExecuteTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduleExecuteTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyScheduleExecuteTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyScheduleExecuteTimeResponse) GetBody() *ModifyScheduleExecuteTimeResponseBody {
	return s.Body
}

func (s *ModifyScheduleExecuteTimeResponse) SetHeaders(v map[string]*string) *ModifyScheduleExecuteTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduleExecuteTimeResponse) SetStatusCode(v int32) *ModifyScheduleExecuteTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScheduleExecuteTimeResponse) SetBody(v *ModifyScheduleExecuteTimeResponseBody) *ModifyScheduleExecuteTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyScheduleExecuteTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
