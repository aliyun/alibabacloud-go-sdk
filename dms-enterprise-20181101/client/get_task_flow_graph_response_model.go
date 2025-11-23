// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskFlowGraphResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskFlowGraphResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskFlowGraphResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskFlowGraphResponseBody) *GetTaskFlowGraphResponse
	GetBody() *GetTaskFlowGraphResponseBody
}

type GetTaskFlowGraphResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskFlowGraphResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskFlowGraphResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowGraphResponse) GoString() string {
	return s.String()
}

func (s *GetTaskFlowGraphResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskFlowGraphResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskFlowGraphResponse) GetBody() *GetTaskFlowGraphResponseBody {
	return s.Body
}

func (s *GetTaskFlowGraphResponse) SetHeaders(v map[string]*string) *GetTaskFlowGraphResponse {
	s.Headers = v
	return s
}

func (s *GetTaskFlowGraphResponse) SetStatusCode(v int32) *GetTaskFlowGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskFlowGraphResponse) SetBody(v *GetTaskFlowGraphResponseBody) *GetTaskFlowGraphResponse {
	s.Body = v
	return s
}

func (s *GetTaskFlowGraphResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
