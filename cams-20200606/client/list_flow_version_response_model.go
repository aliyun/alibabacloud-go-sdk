// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *ListFlowVersionResponseBody) *ListFlowVersionResponse
	GetBody() *ListFlowVersionResponseBody
}

type ListFlowVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *ListFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowVersionResponse) GetBody() *ListFlowVersionResponseBody {
	return s.Body
}

func (s *ListFlowVersionResponse) SetHeaders(v map[string]*string) *ListFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *ListFlowVersionResponse) SetStatusCode(v int32) *ListFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowVersionResponse) SetBody(v *ListFlowVersionResponseBody) *ListFlowVersionResponse {
	s.Body = v
	return s
}

func (s *ListFlowVersionResponse) Validate() error {
	return dara.Validate(s)
}
