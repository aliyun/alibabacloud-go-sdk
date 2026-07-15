// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumerGroupsResponseBody) *ListConsumerGroupsResponse
	GetBody() *ListConsumerGroupsResponseBody
}

type ListConsumerGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumerGroupsResponse) GetBody() *ListConsumerGroupsResponseBody {
	return s.Body
}

func (s *ListConsumerGroupsResponse) SetHeaders(v map[string]*string) *ListConsumerGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerGroupsResponse) SetStatusCode(v int32) *ListConsumerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerGroupsResponse) SetBody(v *ListConsumerGroupsResponseBody) *ListConsumerGroupsResponse {
	s.Body = v
	return s
}

func (s *ListConsumerGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
