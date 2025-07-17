// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTaskFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTaskFlowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTaskFlowResponseBody) *DeleteTaskFlowResponse
	GetBody() *DeleteTaskFlowResponseBody
}

type DeleteTaskFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTaskFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTaskFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTaskFlowResponse) GetBody() *DeleteTaskFlowResponseBody {
	return s.Body
}

func (s *DeleteTaskFlowResponse) SetHeaders(v map[string]*string) *DeleteTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskFlowResponse) SetStatusCode(v int32) *DeleteTaskFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskFlowResponse) SetBody(v *DeleteTaskFlowResponseBody) *DeleteTaskFlowResponse {
	s.Body = v
	return s
}

func (s *DeleteTaskFlowResponse) Validate() error {
	return dara.Validate(s)
}
