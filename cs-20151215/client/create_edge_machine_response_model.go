// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEdgeMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEdgeMachineResponse
	GetStatusCode() *int32
	SetBody(v *CreateEdgeMachineResponseBody) *CreateEdgeMachineResponse
	GetBody() *CreateEdgeMachineResponseBody
}

type CreateEdgeMachineResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEdgeMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEdgeMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeMachineResponse) GoString() string {
	return s.String()
}

func (s *CreateEdgeMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEdgeMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEdgeMachineResponse) GetBody() *CreateEdgeMachineResponseBody {
	return s.Body
}

func (s *CreateEdgeMachineResponse) SetHeaders(v map[string]*string) *CreateEdgeMachineResponse {
	s.Headers = v
	return s
}

func (s *CreateEdgeMachineResponse) SetStatusCode(v int32) *CreateEdgeMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEdgeMachineResponse) SetBody(v *CreateEdgeMachineResponseBody) *CreateEdgeMachineResponse {
	s.Body = v
	return s
}

func (s *CreateEdgeMachineResponse) Validate() error {
	return dara.Validate(s)
}
