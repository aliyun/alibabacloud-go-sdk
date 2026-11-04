// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecordPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecordPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecordPlanResponseBody) *UpdateRecordPlanResponse
	GetBody() *UpdateRecordPlanResponseBody
}

type UpdateRecordPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecordPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecordPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecordPlanResponse) GetBody() *UpdateRecordPlanResponseBody {
	return s.Body
}

func (s *UpdateRecordPlanResponse) SetHeaders(v map[string]*string) *UpdateRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordPlanResponse) SetStatusCode(v int32) *UpdateRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordPlanResponse) SetBody(v *UpdateRecordPlanResponseBody) *UpdateRecordPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateRecordPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
