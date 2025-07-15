// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishContactFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishContactFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishContactFlowResponse
	GetStatusCode() *int32
	SetBody(v *PublishContactFlowResponseBody) *PublishContactFlowResponse
	GetBody() *PublishContactFlowResponseBody
}

type PublishContactFlowResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishContactFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishContactFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishContactFlowResponse) GoString() string {
	return s.String()
}

func (s *PublishContactFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishContactFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishContactFlowResponse) GetBody() *PublishContactFlowResponseBody {
	return s.Body
}

func (s *PublishContactFlowResponse) SetHeaders(v map[string]*string) *PublishContactFlowResponse {
	s.Headers = v
	return s
}

func (s *PublishContactFlowResponse) SetStatusCode(v int32) *PublishContactFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishContactFlowResponse) SetBody(v *PublishContactFlowResponseBody) *PublishContactFlowResponse {
	s.Body = v
	return s
}

func (s *PublishContactFlowResponse) Validate() error {
	return dara.Validate(s)
}
