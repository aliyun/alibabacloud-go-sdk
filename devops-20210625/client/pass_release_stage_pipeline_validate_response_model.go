// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPassReleaseStagePipelineValidateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PassReleaseStagePipelineValidateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PassReleaseStagePipelineValidateResponse
	GetStatusCode() *int32
	SetBody(v *PassReleaseStagePipelineValidateResponseBody) *PassReleaseStagePipelineValidateResponse
	GetBody() *PassReleaseStagePipelineValidateResponseBody
}

type PassReleaseStagePipelineValidateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PassReleaseStagePipelineValidateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PassReleaseStagePipelineValidateResponse) String() string {
	return dara.Prettify(s)
}

func (s PassReleaseStagePipelineValidateResponse) GoString() string {
	return s.String()
}

func (s *PassReleaseStagePipelineValidateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PassReleaseStagePipelineValidateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PassReleaseStagePipelineValidateResponse) GetBody() *PassReleaseStagePipelineValidateResponseBody {
	return s.Body
}

func (s *PassReleaseStagePipelineValidateResponse) SetHeaders(v map[string]*string) *PassReleaseStagePipelineValidateResponse {
	s.Headers = v
	return s
}

func (s *PassReleaseStagePipelineValidateResponse) SetStatusCode(v int32) *PassReleaseStagePipelineValidateResponse {
	s.StatusCode = &v
	return s
}

func (s *PassReleaseStagePipelineValidateResponse) SetBody(v *PassReleaseStagePipelineValidateResponseBody) *PassReleaseStagePipelineValidateResponse {
	s.Body = v
	return s
}

func (s *PassReleaseStagePipelineValidateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
