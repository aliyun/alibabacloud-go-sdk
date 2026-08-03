// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInspectionScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInspectionScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInspectionScheduleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInspectionScheduleResponseBody) *ModifyInspectionScheduleResponse
	GetBody() *ModifyInspectionScheduleResponseBody
}

type ModifyInspectionScheduleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInspectionScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInspectionScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInspectionScheduleResponse) GoString() string {
	return s.String()
}

func (s *ModifyInspectionScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInspectionScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInspectionScheduleResponse) GetBody() *ModifyInspectionScheduleResponseBody {
	return s.Body
}

func (s *ModifyInspectionScheduleResponse) SetHeaders(v map[string]*string) *ModifyInspectionScheduleResponse {
	s.Headers = v
	return s
}

func (s *ModifyInspectionScheduleResponse) SetStatusCode(v int32) *ModifyInspectionScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInspectionScheduleResponse) SetBody(v *ModifyInspectionScheduleResponseBody) *ModifyInspectionScheduleResponse {
	s.Body = v
	return s
}

func (s *ModifyInspectionScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
