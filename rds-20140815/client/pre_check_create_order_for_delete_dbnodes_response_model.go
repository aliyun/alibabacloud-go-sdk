// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckCreateOrderForDeleteDBNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreCheckCreateOrderForDeleteDBNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreCheckCreateOrderForDeleteDBNodesResponse
	GetStatusCode() *int32
	SetBody(v *PreCheckCreateOrderForDeleteDBNodesResponseBody) *PreCheckCreateOrderForDeleteDBNodesResponse
	GetBody() *PreCheckCreateOrderForDeleteDBNodesResponseBody
}

type PreCheckCreateOrderForDeleteDBNodesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreCheckCreateOrderForDeleteDBNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreCheckCreateOrderForDeleteDBNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s PreCheckCreateOrderForDeleteDBNodesResponse) GoString() string {
	return s.String()
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponse) GetBody() *PreCheckCreateOrderForDeleteDBNodesResponseBody {
	return s.Body
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponse) SetHeaders(v map[string]*string) *PreCheckCreateOrderForDeleteDBNodesResponse {
	s.Headers = v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponse) SetStatusCode(v int32) *PreCheckCreateOrderForDeleteDBNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponse) SetBody(v *PreCheckCreateOrderForDeleteDBNodesResponseBody) *PreCheckCreateOrderForDeleteDBNodesResponse {
	s.Body = v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
