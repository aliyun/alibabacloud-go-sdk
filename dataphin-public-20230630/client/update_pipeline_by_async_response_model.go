// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineByAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePipelineByAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePipelineByAsyncResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePipelineByAsyncResponseBody) *UpdatePipelineByAsyncResponse
	GetBody() *UpdatePipelineByAsyncResponseBody
}

type UpdatePipelineByAsyncResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePipelineByAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePipelineByAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePipelineByAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePipelineByAsyncResponse) GetBody() *UpdatePipelineByAsyncResponseBody {
	return s.Body
}

func (s *UpdatePipelineByAsyncResponse) SetHeaders(v map[string]*string) *UpdatePipelineByAsyncResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineByAsyncResponse) SetStatusCode(v int32) *UpdatePipelineByAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineByAsyncResponse) SetBody(v *UpdatePipelineByAsyncResponseBody) *UpdatePipelineByAsyncResponse {
	s.Body = v
	return s
}

func (s *UpdatePipelineByAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
