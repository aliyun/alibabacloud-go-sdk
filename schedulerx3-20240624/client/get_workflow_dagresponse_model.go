// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowDAGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowDAGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowDAGResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowDAGResponseBody) *GetWorkflowDAGResponse
	GetBody() *GetWorkflowDAGResponseBody
}

type GetWorkflowDAGResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowDAGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowDAGResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowDAGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowDAGResponse) GetBody() *GetWorkflowDAGResponseBody {
	return s.Body
}

func (s *GetWorkflowDAGResponse) SetHeaders(v map[string]*string) *GetWorkflowDAGResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowDAGResponse) SetStatusCode(v int32) *GetWorkflowDAGResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowDAGResponse) SetBody(v *GetWorkflowDAGResponseBody) *GetWorkflowDAGResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowDAGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
