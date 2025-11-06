// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushAnalysisTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPushAnalysisTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPushAnalysisTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryPushAnalysisTaskDetailResponseBody) *QueryPushAnalysisTaskDetailResponse
	GetBody() *QueryPushAnalysisTaskDetailResponseBody
}

type QueryPushAnalysisTaskDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPushAnalysisTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPushAnalysisTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPushAnalysisTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPushAnalysisTaskDetailResponse) GetBody() *QueryPushAnalysisTaskDetailResponseBody {
	return s.Body
}

func (s *QueryPushAnalysisTaskDetailResponse) SetHeaders(v map[string]*string) *QueryPushAnalysisTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponse) SetStatusCode(v int32) *QueryPushAnalysisTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponse) SetBody(v *QueryPushAnalysisTaskDetailResponseBody) *QueryPushAnalysisTaskDetailResponse {
	s.Body = v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
