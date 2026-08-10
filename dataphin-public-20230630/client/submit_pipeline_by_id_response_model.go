// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPipelineByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitPipelineByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitPipelineByIdResponse
	GetStatusCode() *int32
	SetBody(v *SubmitPipelineByIdResponseBody) *SubmitPipelineByIdResponse
	GetBody() *SubmitPipelineByIdResponseBody
}

type SubmitPipelineByIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPipelineByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPipelineByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineByIdResponse) GoString() string {
	return s.String()
}

func (s *SubmitPipelineByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitPipelineByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitPipelineByIdResponse) GetBody() *SubmitPipelineByIdResponseBody {
	return s.Body
}

func (s *SubmitPipelineByIdResponse) SetHeaders(v map[string]*string) *SubmitPipelineByIdResponse {
	s.Headers = v
	return s
}

func (s *SubmitPipelineByIdResponse) SetStatusCode(v int32) *SubmitPipelineByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPipelineByIdResponse) SetBody(v *SubmitPipelineByIdResponseBody) *SubmitPipelineByIdResponse {
	s.Body = v
	return s
}

func (s *SubmitPipelineByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
