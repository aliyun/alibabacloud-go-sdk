// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStorageAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStorageAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateStorageAnalysisTaskResponseBody) *CreateStorageAnalysisTaskResponse
	GetBody() *CreateStorageAnalysisTaskResponseBody
}

type CreateStorageAnalysisTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStorageAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStorageAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateStorageAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStorageAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStorageAnalysisTaskResponse) GetBody() *CreateStorageAnalysisTaskResponseBody {
	return s.Body
}

func (s *CreateStorageAnalysisTaskResponse) SetHeaders(v map[string]*string) *CreateStorageAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateStorageAnalysisTaskResponse) SetStatusCode(v int32) *CreateStorageAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStorageAnalysisTaskResponse) SetBody(v *CreateStorageAnalysisTaskResponseBody) *CreateStorageAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *CreateStorageAnalysisTaskResponse) Validate() error {
	return dara.Validate(s)
}
