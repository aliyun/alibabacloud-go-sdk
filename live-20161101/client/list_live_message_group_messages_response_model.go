// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveMessageGroupMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveMessageGroupMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveMessageGroupMessagesResponseBody) *ListLiveMessageGroupMessagesResponse
	GetBody() *ListLiveMessageGroupMessagesResponseBody
}

type ListLiveMessageGroupMessagesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveMessageGroupMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveMessageGroupMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveMessageGroupMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveMessageGroupMessagesResponse) GetBody() *ListLiveMessageGroupMessagesResponseBody {
	return s.Body
}

func (s *ListLiveMessageGroupMessagesResponse) SetHeaders(v map[string]*string) *ListLiveMessageGroupMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListLiveMessageGroupMessagesResponse) SetStatusCode(v int32) *ListLiveMessageGroupMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponse) SetBody(v *ListLiveMessageGroupMessagesResponseBody) *ListLiveMessageGroupMessagesResponse {
	s.Body = v
	return s
}

func (s *ListLiveMessageGroupMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
