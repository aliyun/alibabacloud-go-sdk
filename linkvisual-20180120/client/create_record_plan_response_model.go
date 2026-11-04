// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecordPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecordPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecordPlanResponseBody) *CreateRecordPlanResponse
	GetBody() *CreateRecordPlanResponseBody
}

type CreateRecordPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecordPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateRecordPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecordPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecordPlanResponse) GetBody() *CreateRecordPlanResponseBody {
	return s.Body
}

func (s *CreateRecordPlanResponse) SetHeaders(v map[string]*string) *CreateRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateRecordPlanResponse) SetStatusCode(v int32) *CreateRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecordPlanResponse) SetBody(v *CreateRecordPlanResponseBody) *CreateRecordPlanResponse {
	s.Body = v
	return s
}

func (s *CreateRecordPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
