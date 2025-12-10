// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAINodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAINodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAINodesResponse
	GetStatusCode() *int32
	SetBody(v *CreateAINodesResponseBody) *CreateAINodesResponse
	GetBody() *CreateAINodesResponseBody
}

type CreateAINodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAINodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAINodesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAINodesResponse) GoString() string {
	return s.String()
}

func (s *CreateAINodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAINodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAINodesResponse) GetBody() *CreateAINodesResponseBody {
	return s.Body
}

func (s *CreateAINodesResponse) SetHeaders(v map[string]*string) *CreateAINodesResponse {
	s.Headers = v
	return s
}

func (s *CreateAINodesResponse) SetStatusCode(v int32) *CreateAINodesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAINodesResponse) SetBody(v *CreateAINodesResponseBody) *CreateAINodesResponse {
	s.Body = v
	return s
}

func (s *CreateAINodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
