// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishPipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbolishPipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbolishPipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *AbolishPipelineRunResponseBody) *AbolishPipelineRunResponse
	GetBody() *AbolishPipelineRunResponseBody
}

type AbolishPipelineRunResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbolishPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbolishPipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s AbolishPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *AbolishPipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbolishPipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbolishPipelineRunResponse) GetBody() *AbolishPipelineRunResponseBody {
	return s.Body
}

func (s *AbolishPipelineRunResponse) SetHeaders(v map[string]*string) *AbolishPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *AbolishPipelineRunResponse) SetStatusCode(v int32) *AbolishPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *AbolishPipelineRunResponse) SetBody(v *AbolishPipelineRunResponseBody) *AbolishPipelineRunResponse {
	s.Body = v
	return s
}

func (s *AbolishPipelineRunResponse) Validate() error {
	return dara.Validate(s)
}
