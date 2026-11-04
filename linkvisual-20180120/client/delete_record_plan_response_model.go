// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecordPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecordPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecordPlanResponseBody) *DeleteRecordPlanResponse
	GetBody() *DeleteRecordPlanResponseBody
}

type DeleteRecordPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecordPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecordPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecordPlanResponse) GetBody() *DeleteRecordPlanResponseBody {
	return s.Body
}

func (s *DeleteRecordPlanResponse) SetHeaders(v map[string]*string) *DeleteRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordPlanResponse) SetStatusCode(v int32) *DeleteRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordPlanResponse) SetBody(v *DeleteRecordPlanResponseBody) *DeleteRecordPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteRecordPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
