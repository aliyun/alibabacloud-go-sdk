// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateWorkflowInstancesResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCreateWorkflowInstancesResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCreateWorkflowInstancesResultResponse
	GetStatusCode() *int32
	SetBody(v *GetCreateWorkflowInstancesResultResponseBody) *GetCreateWorkflowInstancesResultResponse
	GetBody() *GetCreateWorkflowInstancesResultResponseBody
}

type GetCreateWorkflowInstancesResultResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreateWorkflowInstancesResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreateWorkflowInstancesResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCreateWorkflowInstancesResultResponse) GoString() string {
	return s.String()
}

func (s *GetCreateWorkflowInstancesResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCreateWorkflowInstancesResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCreateWorkflowInstancesResultResponse) GetBody() *GetCreateWorkflowInstancesResultResponseBody {
	return s.Body
}

func (s *GetCreateWorkflowInstancesResultResponse) SetHeaders(v map[string]*string) *GetCreateWorkflowInstancesResultResponse {
	s.Headers = v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponse) SetStatusCode(v int32) *GetCreateWorkflowInstancesResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponse) SetBody(v *GetCreateWorkflowInstancesResultResponseBody) *GetCreateWorkflowInstancesResultResponse {
	s.Body = v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
