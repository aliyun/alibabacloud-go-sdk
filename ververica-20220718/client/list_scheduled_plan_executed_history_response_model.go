// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPlanExecutedHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScheduledPlanExecutedHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScheduledPlanExecutedHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListScheduledPlanExecutedHistoryResponseBody) *ListScheduledPlanExecutedHistoryResponse
	GetBody() *ListScheduledPlanExecutedHistoryResponseBody
}

type ListScheduledPlanExecutedHistoryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledPlanExecutedHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledPlanExecutedHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPlanExecutedHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanExecutedHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScheduledPlanExecutedHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScheduledPlanExecutedHistoryResponse) GetBody() *ListScheduledPlanExecutedHistoryResponseBody {
	return s.Body
}

func (s *ListScheduledPlanExecutedHistoryResponse) SetHeaders(v map[string]*string) *ListScheduledPlanExecutedHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponse) SetStatusCode(v int32) *ListScheduledPlanExecutedHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponse) SetBody(v *ListScheduledPlanExecutedHistoryResponseBody) *ListScheduledPlanExecutedHistoryResponse {
	s.Body = v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
