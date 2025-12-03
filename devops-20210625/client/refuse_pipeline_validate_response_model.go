// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefusePipelineValidateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefusePipelineValidateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefusePipelineValidateResponse
	GetStatusCode() *int32
	SetBody(v *RefusePipelineValidateResponseBody) *RefusePipelineValidateResponse
	GetBody() *RefusePipelineValidateResponseBody
}

type RefusePipelineValidateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefusePipelineValidateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefusePipelineValidateResponse) String() string {
	return dara.Prettify(s)
}

func (s RefusePipelineValidateResponse) GoString() string {
	return s.String()
}

func (s *RefusePipelineValidateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefusePipelineValidateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefusePipelineValidateResponse) GetBody() *RefusePipelineValidateResponseBody {
	return s.Body
}

func (s *RefusePipelineValidateResponse) SetHeaders(v map[string]*string) *RefusePipelineValidateResponse {
	s.Headers = v
	return s
}

func (s *RefusePipelineValidateResponse) SetStatusCode(v int32) *RefusePipelineValidateResponse {
	s.StatusCode = &v
	return s
}

func (s *RefusePipelineValidateResponse) SetBody(v *RefusePipelineValidateResponseBody) *RefusePipelineValidateResponse {
	s.Body = v
	return s
}

func (s *RefusePipelineValidateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
