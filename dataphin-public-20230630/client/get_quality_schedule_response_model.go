// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityScheduleResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityScheduleResponseBody) *GetQualityScheduleResponse
	GetBody() *GetQualityScheduleResponseBody
}

type GetQualityScheduleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetQualityScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityScheduleResponse) GetBody() *GetQualityScheduleResponseBody {
	return s.Body
}

func (s *GetQualityScheduleResponse) SetHeaders(v map[string]*string) *GetQualityScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetQualityScheduleResponse) SetStatusCode(v int32) *GetQualityScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityScheduleResponse) SetBody(v *GetQualityScheduleResponseBody) *GetQualityScheduleResponse {
	s.Body = v
	return s
}

func (s *GetQualityScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
