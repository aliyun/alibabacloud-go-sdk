// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDialogAnalysisResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDialogAnalysisResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDialogAnalysisResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDialogAnalysisResultResponseBody) *GetDialogAnalysisResultResponse
	GetBody() *GetDialogAnalysisResultResponseBody
}

type GetDialogAnalysisResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDialogAnalysisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDialogAnalysisResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDialogAnalysisResultResponse) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDialogAnalysisResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDialogAnalysisResultResponse) GetBody() *GetDialogAnalysisResultResponseBody {
	return s.Body
}

func (s *GetDialogAnalysisResultResponse) SetHeaders(v map[string]*string) *GetDialogAnalysisResultResponse {
	s.Headers = v
	return s
}

func (s *GetDialogAnalysisResultResponse) SetStatusCode(v int32) *GetDialogAnalysisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDialogAnalysisResultResponse) SetBody(v *GetDialogAnalysisResultResponseBody) *GetDialogAnalysisResultResponse {
	s.Body = v
	return s
}

func (s *GetDialogAnalysisResultResponse) Validate() error {
	return dara.Validate(s)
}
