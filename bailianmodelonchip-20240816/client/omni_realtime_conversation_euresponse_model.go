// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOmniRealtimeConversationEUResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OmniRealtimeConversationEUResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OmniRealtimeConversationEUResponse
	GetStatusCode() *int32
	SetBody(v *OmniRealtimeConversationEUResponseBody) *OmniRealtimeConversationEUResponse
	GetBody() *OmniRealtimeConversationEUResponseBody
}

type OmniRealtimeConversationEUResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OmniRealtimeConversationEUResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OmniRealtimeConversationEUResponse) String() string {
	return dara.Prettify(s)
}

func (s OmniRealtimeConversationEUResponse) GoString() string {
	return s.String()
}

func (s *OmniRealtimeConversationEUResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OmniRealtimeConversationEUResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OmniRealtimeConversationEUResponse) GetBody() *OmniRealtimeConversationEUResponseBody {
	return s.Body
}

func (s *OmniRealtimeConversationEUResponse) SetHeaders(v map[string]*string) *OmniRealtimeConversationEUResponse {
	s.Headers = v
	return s
}

func (s *OmniRealtimeConversationEUResponse) SetStatusCode(v int32) *OmniRealtimeConversationEUResponse {
	s.StatusCode = &v
	return s
}

func (s *OmniRealtimeConversationEUResponse) SetBody(v *OmniRealtimeConversationEUResponseBody) *OmniRealtimeConversationEUResponse {
	s.Body = v
	return s
}

func (s *OmniRealtimeConversationEUResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
