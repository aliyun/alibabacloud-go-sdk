// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMediaWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMediaWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *AddMediaWorkflowResponseBody) *AddMediaWorkflowResponse
	GetBody() *AddMediaWorkflowResponseBody
}

type AddMediaWorkflowResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMediaWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMediaWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMediaWorkflowResponse) GoString() string {
	return s.String()
}

func (s *AddMediaWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMediaWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMediaWorkflowResponse) GetBody() *AddMediaWorkflowResponseBody {
	return s.Body
}

func (s *AddMediaWorkflowResponse) SetHeaders(v map[string]*string) *AddMediaWorkflowResponse {
	s.Headers = v
	return s
}

func (s *AddMediaWorkflowResponse) SetStatusCode(v int32) *AddMediaWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMediaWorkflowResponse) SetBody(v *AddMediaWorkflowResponseBody) *AddMediaWorkflowResponse {
	s.Body = v
	return s
}

func (s *AddMediaWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
