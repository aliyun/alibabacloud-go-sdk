// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportWorkflowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportWorkflowsResponse
	GetStatusCode() *int32
	SetBody(v *ImportWorkflowsResponseBody) *ImportWorkflowsResponse
	GetBody() *ImportWorkflowsResponseBody
}

type ImportWorkflowsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportWorkflowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportWorkflowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *ImportWorkflowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportWorkflowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportWorkflowsResponse) GetBody() *ImportWorkflowsResponseBody {
	return s.Body
}

func (s *ImportWorkflowsResponse) SetHeaders(v map[string]*string) *ImportWorkflowsResponse {
	s.Headers = v
	return s
}

func (s *ImportWorkflowsResponse) SetStatusCode(v int32) *ImportWorkflowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportWorkflowsResponse) SetBody(v *ImportWorkflowsResponseBody) *ImportWorkflowsResponse {
	s.Body = v
	return s
}

func (s *ImportWorkflowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
