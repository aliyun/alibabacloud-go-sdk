// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRectifyFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRectifyFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRectifyFlowResponse
	GetStatusCode() *int32
	SetBody(v *CreateRectifyFlowResponseBody) *CreateRectifyFlowResponse
	GetBody() *CreateRectifyFlowResponseBody
}

type CreateRectifyFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRectifyFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRectifyFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRectifyFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateRectifyFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRectifyFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRectifyFlowResponse) GetBody() *CreateRectifyFlowResponseBody {
	return s.Body
}

func (s *CreateRectifyFlowResponse) SetHeaders(v map[string]*string) *CreateRectifyFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateRectifyFlowResponse) SetStatusCode(v int32) *CreateRectifyFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRectifyFlowResponse) SetBody(v *CreateRectifyFlowResponseBody) *CreateRectifyFlowResponse {
	s.Body = v
	return s
}

func (s *CreateRectifyFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
