// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPipelineAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitPipelineAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitPipelineAsyncResponse
	GetStatusCode() *int32
	SetBody(v *SubmitPipelineAsyncResponseBody) *SubmitPipelineAsyncResponse
	GetBody() *SubmitPipelineAsyncResponseBody
}

type SubmitPipelineAsyncResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPipelineAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPipelineAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineAsyncResponse) GoString() string {
	return s.String()
}

func (s *SubmitPipelineAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitPipelineAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitPipelineAsyncResponse) GetBody() *SubmitPipelineAsyncResponseBody {
	return s.Body
}

func (s *SubmitPipelineAsyncResponse) SetHeaders(v map[string]*string) *SubmitPipelineAsyncResponse {
	s.Headers = v
	return s
}

func (s *SubmitPipelineAsyncResponse) SetStatusCode(v int32) *SubmitPipelineAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPipelineAsyncResponse) SetBody(v *SubmitPipelineAsyncResponseBody) *SubmitPipelineAsyncResponse {
	s.Body = v
	return s
}

func (s *SubmitPipelineAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
