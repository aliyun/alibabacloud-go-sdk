// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPublishHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppPublishHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppPublishHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListAppPublishHistoryResponseBody) *ListAppPublishHistoryResponse
	GetBody() *ListAppPublishHistoryResponseBody
}

type ListAppPublishHistoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppPublishHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppPublishHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppPublishHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListAppPublishHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppPublishHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppPublishHistoryResponse) GetBody() *ListAppPublishHistoryResponseBody {
	return s.Body
}

func (s *ListAppPublishHistoryResponse) SetHeaders(v map[string]*string) *ListAppPublishHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListAppPublishHistoryResponse) SetStatusCode(v int32) *ListAppPublishHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppPublishHistoryResponse) SetBody(v *ListAppPublishHistoryResponseBody) *ListAppPublishHistoryResponse {
	s.Body = v
	return s
}

func (s *ListAppPublishHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
