// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateSystemRunningPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateSystemRunningPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateSystemRunningPlanResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateSystemRunningPlanResponseBody) *BatchUpdateSystemRunningPlanResponse
	GetBody() *BatchUpdateSystemRunningPlanResponseBody
}

type BatchUpdateSystemRunningPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateSystemRunningPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateSystemRunningPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateSystemRunningPlanResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateSystemRunningPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateSystemRunningPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateSystemRunningPlanResponse) GetBody() *BatchUpdateSystemRunningPlanResponseBody {
	return s.Body
}

func (s *BatchUpdateSystemRunningPlanResponse) SetHeaders(v map[string]*string) *BatchUpdateSystemRunningPlanResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateSystemRunningPlanResponse) SetStatusCode(v int32) *BatchUpdateSystemRunningPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanResponse) SetBody(v *BatchUpdateSystemRunningPlanResponseBody) *BatchUpdateSystemRunningPlanResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateSystemRunningPlanResponse) Validate() error {
	return dara.Validate(s)
}
