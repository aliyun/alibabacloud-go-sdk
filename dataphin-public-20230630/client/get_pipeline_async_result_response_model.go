// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineAsyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPipelineAsyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPipelineAsyncResultResponse
	GetStatusCode() *int32
	SetBody(v *GetPipelineAsyncResultResponseBody) *GetPipelineAsyncResultResponse
	GetBody() *GetPipelineAsyncResultResponseBody
}

type GetPipelineAsyncResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPipelineAsyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineAsyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineAsyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineAsyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPipelineAsyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPipelineAsyncResultResponse) GetBody() *GetPipelineAsyncResultResponseBody {
	return s.Body
}

func (s *GetPipelineAsyncResultResponse) SetHeaders(v map[string]*string) *GetPipelineAsyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineAsyncResultResponse) SetStatusCode(v int32) *GetPipelineAsyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineAsyncResultResponse) SetBody(v *GetPipelineAsyncResultResponseBody) *GetPipelineAsyncResultResponse {
	s.Body = v
	return s
}

func (s *GetPipelineAsyncResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
