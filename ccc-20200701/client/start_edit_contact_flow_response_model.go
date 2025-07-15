// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEditContactFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartEditContactFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartEditContactFlowResponse
	GetStatusCode() *int32
	SetBody(v *StartEditContactFlowResponseBody) *StartEditContactFlowResponse
	GetBody() *StartEditContactFlowResponseBody
}

type StartEditContactFlowResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartEditContactFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartEditContactFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s StartEditContactFlowResponse) GoString() string {
	return s.String()
}

func (s *StartEditContactFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartEditContactFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartEditContactFlowResponse) GetBody() *StartEditContactFlowResponseBody {
	return s.Body
}

func (s *StartEditContactFlowResponse) SetHeaders(v map[string]*string) *StartEditContactFlowResponse {
	s.Headers = v
	return s
}

func (s *StartEditContactFlowResponse) SetStatusCode(v int32) *StartEditContactFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *StartEditContactFlowResponse) SetBody(v *StartEditContactFlowResponseBody) *StartEditContactFlowResponse {
	s.Body = v
	return s
}

func (s *StartEditContactFlowResponse) Validate() error {
	return dara.Validate(s)
}
