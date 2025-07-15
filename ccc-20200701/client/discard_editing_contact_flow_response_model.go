// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscardEditingContactFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiscardEditingContactFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiscardEditingContactFlowResponse
	GetStatusCode() *int32
	SetBody(v *DiscardEditingContactFlowResponseBody) *DiscardEditingContactFlowResponse
	GetBody() *DiscardEditingContactFlowResponseBody
}

type DiscardEditingContactFlowResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiscardEditingContactFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiscardEditingContactFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DiscardEditingContactFlowResponse) GoString() string {
	return s.String()
}

func (s *DiscardEditingContactFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiscardEditingContactFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiscardEditingContactFlowResponse) GetBody() *DiscardEditingContactFlowResponseBody {
	return s.Body
}

func (s *DiscardEditingContactFlowResponse) SetHeaders(v map[string]*string) *DiscardEditingContactFlowResponse {
	s.Headers = v
	return s
}

func (s *DiscardEditingContactFlowResponse) SetStatusCode(v int32) *DiscardEditingContactFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DiscardEditingContactFlowResponse) SetBody(v *DiscardEditingContactFlowResponseBody) *DiscardEditingContactFlowResponse {
	s.Body = v
	return s
}

func (s *DiscardEditingContactFlowResponse) Validate() error {
	return dara.Validate(s)
}
