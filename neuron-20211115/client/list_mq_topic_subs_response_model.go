// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMqTopicSubsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMqTopicSubsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMqTopicSubsResponse
	GetStatusCode() *int32
	SetBody(v *MqTopicSubsPageResult) *ListMqTopicSubsResponse
	GetBody() *MqTopicSubsPageResult
}

type ListMqTopicSubsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MqTopicSubsPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMqTopicSubsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMqTopicSubsResponse) GoString() string {
	return s.String()
}

func (s *ListMqTopicSubsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMqTopicSubsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMqTopicSubsResponse) GetBody() *MqTopicSubsPageResult {
	return s.Body
}

func (s *ListMqTopicSubsResponse) SetHeaders(v map[string]*string) *ListMqTopicSubsResponse {
	s.Headers = v
	return s
}

func (s *ListMqTopicSubsResponse) SetStatusCode(v int32) *ListMqTopicSubsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMqTopicSubsResponse) SetBody(v *MqTopicSubsPageResult) *ListMqTopicSubsResponse {
	s.Body = v
	return s
}

func (s *ListMqTopicSubsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
