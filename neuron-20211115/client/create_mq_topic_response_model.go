// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMqTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMqTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMqTopicResponse
	GetStatusCode() *int32
	SetBody(v *MqTopic) *CreateMqTopicResponse
	GetBody() *MqTopic
}

type CreateMqTopicResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MqTopic           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMqTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMqTopicResponse) GoString() string {
	return s.String()
}

func (s *CreateMqTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMqTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMqTopicResponse) GetBody() *MqTopic {
	return s.Body
}

func (s *CreateMqTopicResponse) SetHeaders(v map[string]*string) *CreateMqTopicResponse {
	s.Headers = v
	return s
}

func (s *CreateMqTopicResponse) SetStatusCode(v int32) *CreateMqTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMqTopicResponse) SetBody(v *MqTopic) *CreateMqTopicResponse {
	s.Body = v
	return s
}

func (s *CreateMqTopicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
