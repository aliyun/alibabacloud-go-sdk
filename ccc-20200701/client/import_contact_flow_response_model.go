// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportContactFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportContactFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportContactFlowResponse
	GetStatusCode() *int32
	SetBody(v *ImportContactFlowResponseBody) *ImportContactFlowResponse
	GetBody() *ImportContactFlowResponseBody
}

type ImportContactFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportContactFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportContactFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportContactFlowResponse) GoString() string {
	return s.String()
}

func (s *ImportContactFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportContactFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportContactFlowResponse) GetBody() *ImportContactFlowResponseBody {
	return s.Body
}

func (s *ImportContactFlowResponse) SetHeaders(v map[string]*string) *ImportContactFlowResponse {
	s.Headers = v
	return s
}

func (s *ImportContactFlowResponse) SetStatusCode(v int32) *ImportContactFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportContactFlowResponse) SetBody(v *ImportContactFlowResponseBody) *ImportContactFlowResponse {
	s.Body = v
	return s
}

func (s *ImportContactFlowResponse) Validate() error {
	return dara.Validate(s)
}
