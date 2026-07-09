// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *CancelPipelineRunResponseBody) *CancelPipelineRunResponse
	GetBody() *CancelPipelineRunResponseBody
}

type CancelPipelineRunResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *CancelPipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPipelineRunResponse) GetBody() *CancelPipelineRunResponseBody {
	return s.Body
}

func (s *CancelPipelineRunResponse) SetHeaders(v map[string]*string) *CancelPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *CancelPipelineRunResponse) SetStatusCode(v int32) *CancelPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPipelineRunResponse) SetBody(v *CancelPipelineRunResponseBody) *CancelPipelineRunResponse {
	s.Body = v
	return s
}

func (s *CancelPipelineRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
