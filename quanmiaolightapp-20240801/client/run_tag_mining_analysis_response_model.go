// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTagMiningAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunTagMiningAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunTagMiningAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *RunTagMiningAnalysisResponseBody) *RunTagMiningAnalysisResponse
	GetBody() *RunTagMiningAnalysisResponseBody
}

type RunTagMiningAnalysisResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunTagMiningAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTagMiningAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunTagMiningAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTagMiningAnalysisResponse) GetBody() *RunTagMiningAnalysisResponseBody {
	return s.Body
}

func (s *RunTagMiningAnalysisResponse) SetHeaders(v map[string]*string) *RunTagMiningAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunTagMiningAnalysisResponse) SetStatusCode(v int32) *RunTagMiningAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTagMiningAnalysisResponse) SetBody(v *RunTagMiningAnalysisResponseBody) *RunTagMiningAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunTagMiningAnalysisResponse) Validate() error {
	return dara.Validate(s)
}
