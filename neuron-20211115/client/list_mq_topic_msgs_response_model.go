// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMqTopicMsgsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMqTopicMsgsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMqTopicMsgsResponse
	GetStatusCode() *int32
	SetBody(v *MqMsgPageResult) *ListMqTopicMsgsResponse
	GetBody() *MqMsgPageResult
}

type ListMqTopicMsgsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MqMsgPageResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMqTopicMsgsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMqTopicMsgsResponse) GoString() string {
	return s.String()
}

func (s *ListMqTopicMsgsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMqTopicMsgsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMqTopicMsgsResponse) GetBody() *MqMsgPageResult {
	return s.Body
}

func (s *ListMqTopicMsgsResponse) SetHeaders(v map[string]*string) *ListMqTopicMsgsResponse {
	s.Headers = v
	return s
}

func (s *ListMqTopicMsgsResponse) SetStatusCode(v int32) *ListMqTopicMsgsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMqTopicMsgsResponse) SetBody(v *MqMsgPageResult) *ListMqTopicMsgsResponse {
	s.Body = v
	return s
}

func (s *ListMqTopicMsgsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
