// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowDAGPreviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowDAGPreviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowDAGPreviewResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowDAGPreviewResponseBody) *GetWorkflowDAGPreviewResponse
	GetBody() *GetWorkflowDAGPreviewResponseBody
}

type GetWorkflowDAGPreviewResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowDAGPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowDAGPreviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGPreviewResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGPreviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowDAGPreviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowDAGPreviewResponse) GetBody() *GetWorkflowDAGPreviewResponseBody {
	return s.Body
}

func (s *GetWorkflowDAGPreviewResponse) SetHeaders(v map[string]*string) *GetWorkflowDAGPreviewResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowDAGPreviewResponse) SetStatusCode(v int32) *GetWorkflowDAGPreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponse) SetBody(v *GetWorkflowDAGPreviewResponseBody) *GetWorkflowDAGPreviewResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowDAGPreviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
