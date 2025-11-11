// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCustomHotTopicAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCustomHotTopicAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCustomHotTopicAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *RunCustomHotTopicAnalysisResponseBody) *RunCustomHotTopicAnalysisResponse
	GetBody() *RunCustomHotTopicAnalysisResponseBody
}

type RunCustomHotTopicAnalysisResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCustomHotTopicAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCustomHotTopicAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCustomHotTopicAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCustomHotTopicAnalysisResponse) GetBody() *RunCustomHotTopicAnalysisResponseBody {
	return s.Body
}

func (s *RunCustomHotTopicAnalysisResponse) SetHeaders(v map[string]*string) *RunCustomHotTopicAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponse) SetStatusCode(v int32) *RunCustomHotTopicAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponse) SetBody(v *RunCustomHotTopicAnalysisResponseBody) *RunCustomHotTopicAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
