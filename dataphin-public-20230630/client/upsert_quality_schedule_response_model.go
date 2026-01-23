// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertQualityScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertQualityScheduleResponse
	GetStatusCode() *int32
	SetBody(v *UpsertQualityScheduleResponseBody) *UpsertQualityScheduleResponse
	GetBody() *UpsertQualityScheduleResponseBody
}

type UpsertQualityScheduleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertQualityScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertQualityScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpsertQualityScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertQualityScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertQualityScheduleResponse) GetBody() *UpsertQualityScheduleResponseBody {
	return s.Body
}

func (s *UpsertQualityScheduleResponse) SetHeaders(v map[string]*string) *UpsertQualityScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpsertQualityScheduleResponse) SetStatusCode(v int32) *UpsertQualityScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertQualityScheduleResponse) SetBody(v *UpsertQualityScheduleResponseBody) *UpsertQualityScheduleResponse {
	s.Body = v
	return s
}

func (s *UpsertQualityScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
