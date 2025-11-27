// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForDeleteDBNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrderForDeleteDBNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrderForDeleteDBNodesResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrderForDeleteDBNodesResponseBody) *CreateOrderForDeleteDBNodesResponse
	GetBody() *CreateOrderForDeleteDBNodesResponseBody
}

type CreateOrderForDeleteDBNodesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrderForDeleteDBNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrderForDeleteDBNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForDeleteDBNodesResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderForDeleteDBNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrderForDeleteDBNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrderForDeleteDBNodesResponse) GetBody() *CreateOrderForDeleteDBNodesResponseBody {
	return s.Body
}

func (s *CreateOrderForDeleteDBNodesResponse) SetHeaders(v map[string]*string) *CreateOrderForDeleteDBNodesResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderForDeleteDBNodesResponse) SetStatusCode(v int32) *CreateOrderForDeleteDBNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesResponse) SetBody(v *CreateOrderForDeleteDBNodesResponseBody) *CreateOrderForDeleteDBNodesResponse {
	s.Body = v
	return s
}

func (s *CreateOrderForDeleteDBNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
