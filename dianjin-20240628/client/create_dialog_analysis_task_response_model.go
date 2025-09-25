// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDialogAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDialogAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDialogAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDialogAnalysisTaskResponseBody) *CreateDialogAnalysisTaskResponse
	GetBody() *CreateDialogAnalysisTaskResponseBody
}

type CreateDialogAnalysisTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDialogAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDialogAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDialogAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDialogAnalysisTaskResponse) GetBody() *CreateDialogAnalysisTaskResponseBody {
	return s.Body
}

func (s *CreateDialogAnalysisTaskResponse) SetHeaders(v map[string]*string) *CreateDialogAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDialogAnalysisTaskResponse) SetStatusCode(v int32) *CreateDialogAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponse) SetBody(v *CreateDialogAnalysisTaskResponseBody) *CreateDialogAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDialogAnalysisTaskResponse) Validate() error {
	return dara.Validate(s)
}
