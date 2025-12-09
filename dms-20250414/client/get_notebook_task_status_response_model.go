// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotebookTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNotebookTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNotebookTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetNotebookTaskStatusResponseBody) *GetNotebookTaskStatusResponse
	GetBody() *GetNotebookTaskStatusResponseBody
}

type GetNotebookTaskStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNotebookTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNotebookTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNotebookTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetNotebookTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNotebookTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNotebookTaskStatusResponse) GetBody() *GetNotebookTaskStatusResponseBody {
	return s.Body
}

func (s *GetNotebookTaskStatusResponse) SetHeaders(v map[string]*string) *GetNotebookTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetNotebookTaskStatusResponse) SetStatusCode(v int32) *GetNotebookTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNotebookTaskStatusResponse) SetBody(v *GetNotebookTaskStatusResponseBody) *GetNotebookTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetNotebookTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
