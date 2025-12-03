// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPassPipelineValidateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PassPipelineValidateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PassPipelineValidateResponse
	GetStatusCode() *int32
	SetBody(v *PassPipelineValidateResponseBody) *PassPipelineValidateResponse
	GetBody() *PassPipelineValidateResponseBody
}

type PassPipelineValidateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PassPipelineValidateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PassPipelineValidateResponse) String() string {
	return dara.Prettify(s)
}

func (s PassPipelineValidateResponse) GoString() string {
	return s.String()
}

func (s *PassPipelineValidateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PassPipelineValidateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PassPipelineValidateResponse) GetBody() *PassPipelineValidateResponseBody {
	return s.Body
}

func (s *PassPipelineValidateResponse) SetHeaders(v map[string]*string) *PassPipelineValidateResponse {
	s.Headers = v
	return s
}

func (s *PassPipelineValidateResponse) SetStatusCode(v int32) *PassPipelineValidateResponse {
	s.StatusCode = &v
	return s
}

func (s *PassPipelineValidateResponse) SetBody(v *PassPipelineValidateResponseBody) *PassPipelineValidateResponse {
	s.Body = v
	return s
}

func (s *PassPipelineValidateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
