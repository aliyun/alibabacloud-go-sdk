// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliDingGroupMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAliDingGroupMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAliDingGroupMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListAliDingGroupMessagesResponseBody) *ListAliDingGroupMessagesResponse
	GetBody() *ListAliDingGroupMessagesResponseBody
}

type ListAliDingGroupMessagesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliDingGroupMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliDingGroupMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingGroupMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListAliDingGroupMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAliDingGroupMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAliDingGroupMessagesResponse) GetBody() *ListAliDingGroupMessagesResponseBody {
	return s.Body
}

func (s *ListAliDingGroupMessagesResponse) SetHeaders(v map[string]*string) *ListAliDingGroupMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListAliDingGroupMessagesResponse) SetStatusCode(v int32) *ListAliDingGroupMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliDingGroupMessagesResponse) SetBody(v *ListAliDingGroupMessagesResponseBody) *ListAliDingGroupMessagesResponse {
	s.Body = v
	return s
}

func (s *ListAliDingGroupMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
