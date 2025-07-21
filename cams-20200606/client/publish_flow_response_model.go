// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishFlowResponse
	GetStatusCode() *int32
	SetBody(v *PublishFlowResponseBody) *PublishFlowResponse
	GetBody() *PublishFlowResponseBody
}

type PublishFlowResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishFlowResponse) GoString() string {
	return s.String()
}

func (s *PublishFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishFlowResponse) GetBody() *PublishFlowResponseBody {
	return s.Body
}

func (s *PublishFlowResponse) SetHeaders(v map[string]*string) *PublishFlowResponse {
	s.Headers = v
	return s
}

func (s *PublishFlowResponse) SetStatusCode(v int32) *PublishFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishFlowResponse) SetBody(v *PublishFlowResponseBody) *PublishFlowResponse {
	s.Body = v
	return s
}

func (s *PublishFlowResponse) Validate() error {
	return dara.Validate(s)
}
