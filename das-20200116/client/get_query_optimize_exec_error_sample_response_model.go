// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeExecErrorSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryOptimizeExecErrorSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryOptimizeExecErrorSampleResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryOptimizeExecErrorSampleResponseBody) *GetQueryOptimizeExecErrorSampleResponse
	GetBody() *GetQueryOptimizeExecErrorSampleResponseBody
}

type GetQueryOptimizeExecErrorSampleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryOptimizeExecErrorSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryOptimizeExecErrorSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryOptimizeExecErrorSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryOptimizeExecErrorSampleResponse) GetBody() *GetQueryOptimizeExecErrorSampleResponseBody {
	return s.Body
}

func (s *GetQueryOptimizeExecErrorSampleResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeExecErrorSampleResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponse) SetStatusCode(v int32) *GetQueryOptimizeExecErrorSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponse) SetBody(v *GetQueryOptimizeExecErrorSampleResponseBody) *GetQueryOptimizeExecErrorSampleResponse {
	s.Body = v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
