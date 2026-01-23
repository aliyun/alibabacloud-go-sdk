// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualitySchedulesByWatchIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualitySchedulesByWatchIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualitySchedulesByWatchIdResponse
	GetStatusCode() *int32
	SetBody(v *GetQualitySchedulesByWatchIdResponseBody) *GetQualitySchedulesByWatchIdResponse
	GetBody() *GetQualitySchedulesByWatchIdResponseBody
}

type GetQualitySchedulesByWatchIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualitySchedulesByWatchIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualitySchedulesByWatchIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualitySchedulesByWatchIdResponse) GoString() string {
	return s.String()
}

func (s *GetQualitySchedulesByWatchIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualitySchedulesByWatchIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualitySchedulesByWatchIdResponse) GetBody() *GetQualitySchedulesByWatchIdResponseBody {
	return s.Body
}

func (s *GetQualitySchedulesByWatchIdResponse) SetHeaders(v map[string]*string) *GetQualitySchedulesByWatchIdResponse {
	s.Headers = v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponse) SetStatusCode(v int32) *GetQualitySchedulesByWatchIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponse) SetBody(v *GetQualitySchedulesByWatchIdResponseBody) *GetQualitySchedulesByWatchIdResponse {
	s.Body = v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
