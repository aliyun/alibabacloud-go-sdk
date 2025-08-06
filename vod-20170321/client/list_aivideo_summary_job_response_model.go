// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoSummaryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIVideoSummaryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIVideoSummaryJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIVideoSummaryJobResponseBody) *ListAIVideoSummaryJobResponse
	GetBody() *ListAIVideoSummaryJobResponseBody
}

type ListAIVideoSummaryJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIVideoSummaryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIVideoSummaryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoSummaryJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIVideoSummaryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIVideoSummaryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIVideoSummaryJobResponse) GetBody() *ListAIVideoSummaryJobResponseBody {
	return s.Body
}

func (s *ListAIVideoSummaryJobResponse) SetHeaders(v map[string]*string) *ListAIVideoSummaryJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIVideoSummaryJobResponse) SetStatusCode(v int32) *ListAIVideoSummaryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIVideoSummaryJobResponse) SetBody(v *ListAIVideoSummaryJobResponseBody) *ListAIVideoSummaryJobResponse {
	s.Body = v
	return s
}

func (s *ListAIVideoSummaryJobResponse) Validate() error {
	return dara.Validate(s)
}
