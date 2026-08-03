// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspectionScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInspectionScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInspectionScheduleResponse
	GetStatusCode() *int32
	SetBody(v *CreateInspectionScheduleResponseBody) *CreateInspectionScheduleResponse
	GetBody() *CreateInspectionScheduleResponseBody
}

type CreateInspectionScheduleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInspectionScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInspectionScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionScheduleResponse) GoString() string {
	return s.String()
}

func (s *CreateInspectionScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInspectionScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInspectionScheduleResponse) GetBody() *CreateInspectionScheduleResponseBody {
	return s.Body
}

func (s *CreateInspectionScheduleResponse) SetHeaders(v map[string]*string) *CreateInspectionScheduleResponse {
	s.Headers = v
	return s
}

func (s *CreateInspectionScheduleResponse) SetStatusCode(v int32) *CreateInspectionScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInspectionScheduleResponse) SetBody(v *CreateInspectionScheduleResponseBody) *CreateInspectionScheduleResponse {
	s.Body = v
	return s
}

func (s *CreateInspectionScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
