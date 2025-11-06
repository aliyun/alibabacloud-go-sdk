// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushAnalysisTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPushAnalysisTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPushAnalysisTaskListResponse
	GetStatusCode() *int32
	SetBody(v *QueryPushAnalysisTaskListResponseBody) *QueryPushAnalysisTaskListResponse
	GetBody() *QueryPushAnalysisTaskListResponseBody
}

type QueryPushAnalysisTaskListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPushAnalysisTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPushAnalysisTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPushAnalysisTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPushAnalysisTaskListResponse) GetBody() *QueryPushAnalysisTaskListResponseBody {
	return s.Body
}

func (s *QueryPushAnalysisTaskListResponse) SetHeaders(v map[string]*string) *QueryPushAnalysisTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryPushAnalysisTaskListResponse) SetStatusCode(v int32) *QueryPushAnalysisTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushAnalysisTaskListResponse) SetBody(v *QueryPushAnalysisTaskListResponseBody) *QueryPushAnalysisTaskListResponse {
	s.Body = v
	return s
}

func (s *QueryPushAnalysisTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
