// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInspectionScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInspectionScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInspectionScheduleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInspectionScheduleResponseBody) *DeleteInspectionScheduleResponse
	GetBody() *DeleteInspectionScheduleResponseBody
}

type DeleteInspectionScheduleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInspectionScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInspectionScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInspectionScheduleResponse) GoString() string {
	return s.String()
}

func (s *DeleteInspectionScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInspectionScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInspectionScheduleResponse) GetBody() *DeleteInspectionScheduleResponseBody {
	return s.Body
}

func (s *DeleteInspectionScheduleResponse) SetHeaders(v map[string]*string) *DeleteInspectionScheduleResponse {
	s.Headers = v
	return s
}

func (s *DeleteInspectionScheduleResponse) SetStatusCode(v int32) *DeleteInspectionScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInspectionScheduleResponse) SetBody(v *DeleteInspectionScheduleResponseBody) *DeleteInspectionScheduleResponse {
	s.Body = v
	return s
}

func (s *DeleteInspectionScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
