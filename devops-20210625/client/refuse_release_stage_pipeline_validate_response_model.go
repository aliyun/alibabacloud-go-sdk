// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefuseReleaseStagePipelineValidateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefuseReleaseStagePipelineValidateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefuseReleaseStagePipelineValidateResponse
	GetStatusCode() *int32
	SetBody(v *RefuseReleaseStagePipelineValidateResponseBody) *RefuseReleaseStagePipelineValidateResponse
	GetBody() *RefuseReleaseStagePipelineValidateResponseBody
}

type RefuseReleaseStagePipelineValidateResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefuseReleaseStagePipelineValidateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefuseReleaseStagePipelineValidateResponse) String() string {
	return dara.Prettify(s)
}

func (s RefuseReleaseStagePipelineValidateResponse) GoString() string {
	return s.String()
}

func (s *RefuseReleaseStagePipelineValidateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefuseReleaseStagePipelineValidateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefuseReleaseStagePipelineValidateResponse) GetBody() *RefuseReleaseStagePipelineValidateResponseBody {
	return s.Body
}

func (s *RefuseReleaseStagePipelineValidateResponse) SetHeaders(v map[string]*string) *RefuseReleaseStagePipelineValidateResponse {
	s.Headers = v
	return s
}

func (s *RefuseReleaseStagePipelineValidateResponse) SetStatusCode(v int32) *RefuseReleaseStagePipelineValidateResponse {
	s.StatusCode = &v
	return s
}

func (s *RefuseReleaseStagePipelineValidateResponse) SetBody(v *RefuseReleaseStagePipelineValidateResponseBody) *RefuseReleaseStagePipelineValidateResponse {
	s.Body = v
	return s
}

func (s *RefuseReleaseStagePipelineValidateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
