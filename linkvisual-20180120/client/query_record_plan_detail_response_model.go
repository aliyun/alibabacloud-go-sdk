// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordPlanDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordPlanDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordPlanDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordPlanDetailResponseBody) *QueryRecordPlanDetailResponse
	GetBody() *QueryRecordPlanDetailResponseBody
}

type QueryRecordPlanDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordPlanDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordPlanDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlanDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordPlanDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordPlanDetailResponse) GetBody() *QueryRecordPlanDetailResponseBody {
	return s.Body
}

func (s *QueryRecordPlanDetailResponse) SetHeaders(v map[string]*string) *QueryRecordPlanDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordPlanDetailResponse) SetStatusCode(v int32) *QueryRecordPlanDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordPlanDetailResponse) SetBody(v *QueryRecordPlanDetailResponseBody) *QueryRecordPlanDetailResponse {
	s.Body = v
	return s
}

func (s *QueryRecordPlanDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
