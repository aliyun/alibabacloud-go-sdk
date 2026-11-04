// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordPlanDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRecordPlanDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRecordPlanDeviceResponse
	GetStatusCode() *int32
	SetBody(v *AddRecordPlanDeviceResponseBody) *AddRecordPlanDeviceResponse
	GetBody() *AddRecordPlanDeviceResponseBody
}

type AddRecordPlanDeviceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecordPlanDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPlanDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddRecordPlanDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRecordPlanDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRecordPlanDeviceResponse) GetBody() *AddRecordPlanDeviceResponseBody {
	return s.Body
}

func (s *AddRecordPlanDeviceResponse) SetHeaders(v map[string]*string) *AddRecordPlanDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddRecordPlanDeviceResponse) SetStatusCode(v int32) *AddRecordPlanDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecordPlanDeviceResponse) SetBody(v *AddRecordPlanDeviceResponseBody) *AddRecordPlanDeviceResponse {
	s.Body = v
	return s
}

func (s *AddRecordPlanDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
