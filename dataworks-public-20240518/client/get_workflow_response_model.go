// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowResponseBody) *GetWorkflowResponse
	GetBody() *GetWorkflowResponseBody
}

type GetWorkflowResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowResponse) GetBody() *GetWorkflowResponseBody {
	return s.Body
}

func (s *GetWorkflowResponse) SetHeaders(v map[string]*string) *GetWorkflowResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowResponse) SetStatusCode(v int32) *GetWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowResponse) SetBody(v *GetWorkflowResponseBody) *GetWorkflowResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
