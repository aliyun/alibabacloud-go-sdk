// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPipelineByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPipelineByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetPipelineByIdResponseBody) *GetPipelineByIdResponse
	GetBody() *GetPipelineByIdResponseBody
}

type GetPipelineByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPipelineByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPipelineByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPipelineByIdResponse) GetBody() *GetPipelineByIdResponseBody {
	return s.Body
}

func (s *GetPipelineByIdResponse) SetHeaders(v map[string]*string) *GetPipelineByIdResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineByIdResponse) SetStatusCode(v int32) *GetPipelineByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineByIdResponse) SetBody(v *GetPipelineByIdResponseBody) *GetPipelineByIdResponse {
	s.Body = v
	return s
}

func (s *GetPipelineByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
