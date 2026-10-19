// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowInstanceResponseBody) *GetWorkflowInstanceResponse
	GetBody() *GetWorkflowInstanceResponseBody
}

type GetWorkflowInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowInstanceResponse) GetBody() *GetWorkflowInstanceResponseBody {
	return s.Body
}

func (s *GetWorkflowInstanceResponse) SetHeaders(v map[string]*string) *GetWorkflowInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowInstanceResponse) SetStatusCode(v int32) *GetWorkflowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowInstanceResponse) SetBody(v *GetWorkflowInstanceResponseBody) *GetWorkflowInstanceResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
