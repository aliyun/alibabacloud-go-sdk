// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordPlanDeviceByDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordPlanDeviceByDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordPlanDeviceByDeviceResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordPlanDeviceByDeviceResponseBody) *QueryRecordPlanDeviceByDeviceResponse
	GetBody() *QueryRecordPlanDeviceByDeviceResponseBody
}

type QueryRecordPlanDeviceByDeviceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordPlanDeviceByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordPlanDeviceByDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordPlanDeviceByDeviceResponse) GetBody() *QueryRecordPlanDeviceByDeviceResponseBody {
	return s.Body
}

func (s *QueryRecordPlanDeviceByDeviceResponse) SetHeaders(v map[string]*string) *QueryRecordPlanDeviceByDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponse) SetStatusCode(v int32) *QueryRecordPlanDeviceByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponse) SetBody(v *QueryRecordPlanDeviceByDeviceResponseBody) *QueryRecordPlanDeviceByDeviceResponse {
	s.Body = v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
