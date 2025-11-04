// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentDialoguesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIAgentDialoguesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIAgentDialoguesResponse
	GetStatusCode() *int32
	SetBody(v *ListAIAgentDialoguesResponseBody) *ListAIAgentDialoguesResponse
	GetBody() *ListAIAgentDialoguesResponseBody
}

type ListAIAgentDialoguesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIAgentDialoguesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIAgentDialoguesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentDialoguesResponse) GoString() string {
	return s.String()
}

func (s *ListAIAgentDialoguesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIAgentDialoguesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIAgentDialoguesResponse) GetBody() *ListAIAgentDialoguesResponseBody {
	return s.Body
}

func (s *ListAIAgentDialoguesResponse) SetHeaders(v map[string]*string) *ListAIAgentDialoguesResponse {
	s.Headers = v
	return s
}

func (s *ListAIAgentDialoguesResponse) SetStatusCode(v int32) *ListAIAgentDialoguesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIAgentDialoguesResponse) SetBody(v *ListAIAgentDialoguesResponseBody) *ListAIAgentDialoguesResponse {
	s.Body = v
	return s
}

func (s *ListAIAgentDialoguesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
