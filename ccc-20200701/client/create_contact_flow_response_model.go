// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContactFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContactFlowResponse
	GetStatusCode() *int32
	SetBody(v *CreateContactFlowResponseBody) *CreateContactFlowResponse
	GetBody() *CreateContactFlowResponseBody
}

type CreateContactFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContactFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContactFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContactFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateContactFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContactFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContactFlowResponse) GetBody() *CreateContactFlowResponseBody {
	return s.Body
}

func (s *CreateContactFlowResponse) SetHeaders(v map[string]*string) *CreateContactFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateContactFlowResponse) SetStatusCode(v int32) *CreateContactFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContactFlowResponse) SetBody(v *CreateContactFlowResponseBody) *CreateContactFlowResponse {
	s.Body = v
	return s
}

func (s *CreateContactFlowResponse) Validate() error {
	return dara.Validate(s)
}
