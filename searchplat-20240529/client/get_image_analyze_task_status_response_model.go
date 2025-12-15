// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageAnalyzeTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageAnalyzeTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageAnalyzeTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetImageAnalyzeTaskStatusResponseBody) *GetImageAnalyzeTaskStatusResponse
	GetBody() *GetImageAnalyzeTaskStatusResponseBody
}

type GetImageAnalyzeTaskStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageAnalyzeTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageAnalyzeTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageAnalyzeTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetImageAnalyzeTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageAnalyzeTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageAnalyzeTaskStatusResponse) GetBody() *GetImageAnalyzeTaskStatusResponseBody {
	return s.Body
}

func (s *GetImageAnalyzeTaskStatusResponse) SetHeaders(v map[string]*string) *GetImageAnalyzeTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponse) SetStatusCode(v int32) *GetImageAnalyzeTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponse) SetBody(v *GetImageAnalyzeTaskStatusResponseBody) *GetImageAnalyzeTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
