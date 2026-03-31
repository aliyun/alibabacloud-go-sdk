// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListChatConfigurationsResponseBody) *ListChatConfigurationsResponse
	GetBody() *ListChatConfigurationsResponseBody
}

type ListChatConfigurationsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListChatConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatConfigurationsResponse) GetBody() *ListChatConfigurationsResponseBody {
	return s.Body
}

func (s *ListChatConfigurationsResponse) SetHeaders(v map[string]*string) *ListChatConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListChatConfigurationsResponse) SetStatusCode(v int32) *ListChatConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatConfigurationsResponse) SetBody(v *ListChatConfigurationsResponseBody) *ListChatConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListChatConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
