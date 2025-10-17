// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagMiningAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTagMiningAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTagMiningAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetTagMiningAnalysisTaskResponseBody) *GetTagMiningAnalysisTaskResponse
	GetBody() *GetTagMiningAnalysisTaskResponseBody
}

type GetTagMiningAnalysisTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTagMiningAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTagMiningAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTagMiningAnalysisTaskResponse) GetBody() *GetTagMiningAnalysisTaskResponseBody {
	return s.Body
}

func (s *GetTagMiningAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetTagMiningAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponse) SetStatusCode(v int32) *GetTagMiningAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponse) SetBody(v *GetTagMiningAnalysisTaskResponseBody) *GetTagMiningAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
