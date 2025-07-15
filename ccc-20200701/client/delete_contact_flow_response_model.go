// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactFlowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactFlowResponseBody) *DeleteContactFlowResponse
	GetBody() *DeleteContactFlowResponseBody
}

type DeleteContactFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactFlowResponse) GetBody() *DeleteContactFlowResponseBody {
	return s.Body
}

func (s *DeleteContactFlowResponse) SetHeaders(v map[string]*string) *DeleteContactFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactFlowResponse) SetStatusCode(v int32) *DeleteContactFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactFlowResponse) SetBody(v *DeleteContactFlowResponseBody) *DeleteContactFlowResponse {
	s.Body = v
	return s
}

func (s *DeleteContactFlowResponse) Validate() error {
	return dara.Validate(s)
}
