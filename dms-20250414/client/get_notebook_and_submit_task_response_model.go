// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotebookAndSubmitTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNotebookAndSubmitTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNotebookAndSubmitTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetNotebookAndSubmitTaskResponseBody) *GetNotebookAndSubmitTaskResponse
	GetBody() *GetNotebookAndSubmitTaskResponseBody
}

type GetNotebookAndSubmitTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNotebookAndSubmitTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNotebookAndSubmitTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNotebookAndSubmitTaskResponse) GoString() string {
	return s.String()
}

func (s *GetNotebookAndSubmitTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNotebookAndSubmitTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNotebookAndSubmitTaskResponse) GetBody() *GetNotebookAndSubmitTaskResponseBody {
	return s.Body
}

func (s *GetNotebookAndSubmitTaskResponse) SetHeaders(v map[string]*string) *GetNotebookAndSubmitTaskResponse {
	s.Headers = v
	return s
}

func (s *GetNotebookAndSubmitTaskResponse) SetStatusCode(v int32) *GetNotebookAndSubmitTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNotebookAndSubmitTaskResponse) SetBody(v *GetNotebookAndSubmitTaskResponseBody) *GetNotebookAndSubmitTaskResponse {
	s.Body = v
	return s
}

func (s *GetNotebookAndSubmitTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
