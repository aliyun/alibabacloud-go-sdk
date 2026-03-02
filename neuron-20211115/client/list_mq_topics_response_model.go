// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMqTopicsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMqTopicsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMqTopicsResponse
	GetStatusCode() *int32
	SetBody(v *MqTopicPageResult) *ListMqTopicsResponse
	GetBody() *MqTopicPageResult
}

type ListMqTopicsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MqTopicPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMqTopicsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMqTopicsResponse) GoString() string {
	return s.String()
}

func (s *ListMqTopicsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMqTopicsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMqTopicsResponse) GetBody() *MqTopicPageResult {
	return s.Body
}

func (s *ListMqTopicsResponse) SetHeaders(v map[string]*string) *ListMqTopicsResponse {
	s.Headers = v
	return s
}

func (s *ListMqTopicsResponse) SetStatusCode(v int32) *ListMqTopicsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMqTopicsResponse) SetBody(v *MqTopicPageResult) *ListMqTopicsResponse {
	s.Body = v
	return s
}

func (s *ListMqTopicsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
