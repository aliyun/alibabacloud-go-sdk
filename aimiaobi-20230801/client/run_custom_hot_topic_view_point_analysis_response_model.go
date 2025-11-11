// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCustomHotTopicViewPointAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCustomHotTopicViewPointAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCustomHotTopicViewPointAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *RunCustomHotTopicViewPointAnalysisResponseBody) *RunCustomHotTopicViewPointAnalysisResponse
	GetBody() *RunCustomHotTopicViewPointAnalysisResponseBody
}

type RunCustomHotTopicViewPointAnalysisResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCustomHotTopicViewPointAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) GetBody() *RunCustomHotTopicViewPointAnalysisResponseBody {
	return s.Body
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) SetHeaders(v map[string]*string) *RunCustomHotTopicViewPointAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) SetStatusCode(v int32) *RunCustomHotTopicViewPointAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) SetBody(v *RunCustomHotTopicViewPointAnalysisResponseBody) *RunCustomHotTopicViewPointAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
