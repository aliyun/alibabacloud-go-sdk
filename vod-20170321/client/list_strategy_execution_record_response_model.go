// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStrategyExecutionRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStrategyExecutionRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStrategyExecutionRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListStrategyExecutionRecordResponseBody) *ListStrategyExecutionRecordResponse
	GetBody() *ListStrategyExecutionRecordResponseBody
}

type ListStrategyExecutionRecordResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStrategyExecutionRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStrategyExecutionRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStrategyExecutionRecordResponse) GoString() string {
	return s.String()
}

func (s *ListStrategyExecutionRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStrategyExecutionRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStrategyExecutionRecordResponse) GetBody() *ListStrategyExecutionRecordResponseBody {
	return s.Body
}

func (s *ListStrategyExecutionRecordResponse) SetHeaders(v map[string]*string) *ListStrategyExecutionRecordResponse {
	s.Headers = v
	return s
}

func (s *ListStrategyExecutionRecordResponse) SetStatusCode(v int32) *ListStrategyExecutionRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStrategyExecutionRecordResponse) SetBody(v *ListStrategyExecutionRecordResponseBody) *ListStrategyExecutionRecordResponse {
	s.Body = v
	return s
}

func (s *ListStrategyExecutionRecordResponse) Validate() error {
	return dara.Validate(s)
}
