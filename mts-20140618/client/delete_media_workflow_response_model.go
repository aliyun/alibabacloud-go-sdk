// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaWorkflowResponseBody) *DeleteMediaWorkflowResponse
	GetBody() *DeleteMediaWorkflowResponseBody
}

type DeleteMediaWorkflowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaWorkflowResponse) GetBody() *DeleteMediaWorkflowResponseBody {
	return s.Body
}

func (s *DeleteMediaWorkflowResponse) SetHeaders(v map[string]*string) *DeleteMediaWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaWorkflowResponse) SetStatusCode(v int32) *DeleteMediaWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaWorkflowResponse) SetBody(v *DeleteMediaWorkflowResponseBody) *DeleteMediaWorkflowResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
