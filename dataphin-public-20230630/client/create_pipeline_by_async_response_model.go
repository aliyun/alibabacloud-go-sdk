// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineByAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePipelineByAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePipelineByAsyncResponse
	GetStatusCode() *int32
	SetBody(v *CreatePipelineByAsyncResponseBody) *CreatePipelineByAsyncResponse
	GetBody() *CreatePipelineByAsyncResponseBody
}

type CreatePipelineByAsyncResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePipelineByAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineByAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePipelineByAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePipelineByAsyncResponse) GetBody() *CreatePipelineByAsyncResponseBody {
	return s.Body
}

func (s *CreatePipelineByAsyncResponse) SetHeaders(v map[string]*string) *CreatePipelineByAsyncResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineByAsyncResponse) SetStatusCode(v int32) *CreatePipelineByAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineByAsyncResponse) SetBody(v *CreatePipelineByAsyncResponseBody) *CreatePipelineByAsyncResponse {
	s.Body = v
	return s
}

func (s *CreatePipelineByAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
