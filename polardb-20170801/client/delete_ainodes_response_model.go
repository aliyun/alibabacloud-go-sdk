// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAINodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAINodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAINodesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAINodesResponseBody) *DeleteAINodesResponse
	GetBody() *DeleteAINodesResponseBody
}

type DeleteAINodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAINodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAINodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAINodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAINodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAINodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAINodesResponse) GetBody() *DeleteAINodesResponseBody {
	return s.Body
}

func (s *DeleteAINodesResponse) SetHeaders(v map[string]*string) *DeleteAINodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAINodesResponse) SetStatusCode(v int32) *DeleteAINodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAINodesResponse) SetBody(v *DeleteAINodesResponseBody) *DeleteAINodesResponse {
	s.Body = v
	return s
}

func (s *DeleteAINodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
