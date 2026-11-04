// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRecordPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventRecordPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventRecordPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventRecordPlanResponseBody) *DeleteEventRecordPlanResponse
	GetBody() *DeleteEventRecordPlanResponseBody
}

type DeleteEventRecordPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventRecordPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventRecordPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventRecordPlanResponse) GetBody() *DeleteEventRecordPlanResponseBody {
	return s.Body
}

func (s *DeleteEventRecordPlanResponse) SetHeaders(v map[string]*string) *DeleteEventRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRecordPlanResponse) SetStatusCode(v int32) *DeleteEventRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventRecordPlanResponse) SetBody(v *DeleteEventRecordPlanResponseBody) *DeleteEventRecordPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteEventRecordPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
