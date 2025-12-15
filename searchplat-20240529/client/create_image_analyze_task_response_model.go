// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageAnalyzeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageAnalyzeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageAnalyzeTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageAnalyzeTaskResponseBody) *CreateImageAnalyzeTaskResponse
	GetBody() *CreateImageAnalyzeTaskResponseBody
}

type CreateImageAnalyzeTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageAnalyzeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageAnalyzeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageAnalyzeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageAnalyzeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageAnalyzeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageAnalyzeTaskResponse) GetBody() *CreateImageAnalyzeTaskResponseBody {
	return s.Body
}

func (s *CreateImageAnalyzeTaskResponse) SetHeaders(v map[string]*string) *CreateImageAnalyzeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageAnalyzeTaskResponse) SetStatusCode(v int32) *CreateImageAnalyzeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageAnalyzeTaskResponse) SetBody(v *CreateImageAnalyzeTaskResponseBody) *CreateImageAnalyzeTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImageAnalyzeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
