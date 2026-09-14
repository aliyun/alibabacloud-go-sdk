// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishCrossProjectPipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbolishCrossProjectPipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbolishCrossProjectPipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *AbolishCrossProjectPipelineRunResponseBody) *AbolishCrossProjectPipelineRunResponse
	GetBody() *AbolishCrossProjectPipelineRunResponseBody
}

type AbolishCrossProjectPipelineRunResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbolishCrossProjectPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbolishCrossProjectPipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s AbolishCrossProjectPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *AbolishCrossProjectPipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbolishCrossProjectPipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbolishCrossProjectPipelineRunResponse) GetBody() *AbolishCrossProjectPipelineRunResponseBody {
	return s.Body
}

func (s *AbolishCrossProjectPipelineRunResponse) SetHeaders(v map[string]*string) *AbolishCrossProjectPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *AbolishCrossProjectPipelineRunResponse) SetStatusCode(v int32) *AbolishCrossProjectPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *AbolishCrossProjectPipelineRunResponse) SetBody(v *AbolishCrossProjectPipelineRunResponseBody) *AbolishCrossProjectPipelineRunResponse {
	s.Body = v
	return s
}

func (s *AbolishCrossProjectPipelineRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
