// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomSourceTopicAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCustomSourceTopicAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCustomSourceTopicAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCustomSourceTopicAnalysisResponseBody) *SubmitCustomSourceTopicAnalysisResponse
	GetBody() *SubmitCustomSourceTopicAnalysisResponseBody
}

type SubmitCustomSourceTopicAnalysisResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCustomSourceTopicAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCustomSourceTopicAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomSourceTopicAnalysisResponse) GoString() string {
	return s.String()
}

func (s *SubmitCustomSourceTopicAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCustomSourceTopicAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCustomSourceTopicAnalysisResponse) GetBody() *SubmitCustomSourceTopicAnalysisResponseBody {
	return s.Body
}

func (s *SubmitCustomSourceTopicAnalysisResponse) SetHeaders(v map[string]*string) *SubmitCustomSourceTopicAnalysisResponse {
	s.Headers = v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponse) SetStatusCode(v int32) *SubmitCustomSourceTopicAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponse) SetBody(v *SubmitCustomSourceTopicAnalysisResponseBody) *SubmitCustomSourceTopicAnalysisResponse {
	s.Body = v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisResponse) Validate() error {
	return dara.Validate(s)
}
