// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskConversationRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduledTaskConversationRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduledTaskConversationRecordsResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduledTaskConversationRecordsResponseBody) *GetScheduledTaskConversationRecordsResponse
	GetBody() *GetScheduledTaskConversationRecordsResponseBody
}

type GetScheduledTaskConversationRecordsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledTaskConversationRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledTaskConversationRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskConversationRecordsResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskConversationRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduledTaskConversationRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduledTaskConversationRecordsResponse) GetBody() *GetScheduledTaskConversationRecordsResponseBody {
	return s.Body
}

func (s *GetScheduledTaskConversationRecordsResponse) SetHeaders(v map[string]*string) *GetScheduledTaskConversationRecordsResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponse) SetStatusCode(v int32) *GetScheduledTaskConversationRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponse) SetBody(v *GetScheduledTaskConversationRecordsResponseBody) *GetScheduledTaskConversationRecordsResponse {
	s.Body = v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
