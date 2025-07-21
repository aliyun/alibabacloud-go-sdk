// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowResponse
	GetStatusCode() *int32
	SetBody(v *ListFlowResponseBody) *ListFlowResponse
	GetBody() *ListFlowResponseBody
}

type ListFlowResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlowResponse) GoString() string {
	return s.String()
}

func (s *ListFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowResponse) GetBody() *ListFlowResponseBody {
	return s.Body
}

func (s *ListFlowResponse) SetHeaders(v map[string]*string) *ListFlowResponse {
	s.Headers = v
	return s
}

func (s *ListFlowResponse) SetStatusCode(v int32) *ListFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowResponse) SetBody(v *ListFlowResponseBody) *ListFlowResponse {
	s.Body = v
	return s
}

func (s *ListFlowResponse) Validate() error {
	return dara.Validate(s)
}
