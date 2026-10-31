// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPipelineSubmitResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPipelineSubmitResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPipelineSubmitResultResponse
	GetStatusCode() *int32
	SetBody(v *QueryPipelineSubmitResultResponseBody) *QueryPipelineSubmitResultResponse
	GetBody() *QueryPipelineSubmitResultResponseBody
}

type QueryPipelineSubmitResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPipelineSubmitResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPipelineSubmitResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineSubmitResultResponse) GoString() string {
	return s.String()
}

func (s *QueryPipelineSubmitResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPipelineSubmitResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPipelineSubmitResultResponse) GetBody() *QueryPipelineSubmitResultResponseBody {
	return s.Body
}

func (s *QueryPipelineSubmitResultResponse) SetHeaders(v map[string]*string) *QueryPipelineSubmitResultResponse {
	s.Headers = v
	return s
}

func (s *QueryPipelineSubmitResultResponse) SetStatusCode(v int32) *QueryPipelineSubmitResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPipelineSubmitResultResponse) SetBody(v *QueryPipelineSubmitResultResponseBody) *QueryPipelineSubmitResultResponse {
	s.Body = v
	return s
}

func (s *QueryPipelineSubmitResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
