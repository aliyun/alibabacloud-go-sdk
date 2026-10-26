// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConversationTimelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentConversationTimelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentConversationTimelineResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentConversationTimelineResponseBody) *GetAgentConversationTimelineResponse
	GetBody() *GetAgentConversationTimelineResponseBody
}

type GetAgentConversationTimelineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentConversationTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentConversationTimelineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConversationTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetAgentConversationTimelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentConversationTimelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentConversationTimelineResponse) GetBody() *GetAgentConversationTimelineResponseBody {
	return s.Body
}

func (s *GetAgentConversationTimelineResponse) SetHeaders(v map[string]*string) *GetAgentConversationTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetAgentConversationTimelineResponse) SetStatusCode(v int32) *GetAgentConversationTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentConversationTimelineResponse) SetBody(v *GetAgentConversationTimelineResponseBody) *GetAgentConversationTimelineResponse {
	s.Body = v
	return s
}

func (s *GetAgentConversationTimelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
