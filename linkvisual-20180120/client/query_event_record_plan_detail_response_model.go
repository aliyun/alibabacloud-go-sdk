// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventRecordPlanDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEventRecordPlanDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEventRecordPlanDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryEventRecordPlanDetailResponseBody) *QueryEventRecordPlanDetailResponse
	GetBody() *QueryEventRecordPlanDetailResponseBody
}

type QueryEventRecordPlanDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventRecordPlanDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventRecordPlanDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEventRecordPlanDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEventRecordPlanDetailResponse) GetBody() *QueryEventRecordPlanDetailResponseBody {
	return s.Body
}

func (s *QueryEventRecordPlanDetailResponse) SetHeaders(v map[string]*string) *QueryEventRecordPlanDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryEventRecordPlanDetailResponse) SetStatusCode(v int32) *QueryEventRecordPlanDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponse) SetBody(v *QueryEventRecordPlanDetailResponseBody) *QueryEventRecordPlanDetailResponse {
	s.Body = v
	return s
}

func (s *QueryEventRecordPlanDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
