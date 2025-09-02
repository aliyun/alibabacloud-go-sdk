// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertMessagesResponseBody) *ListAlertMessagesResponse
	GetBody() *ListAlertMessagesResponseBody
}

type ListAlertMessagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListAlertMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertMessagesResponse) GetBody() *ListAlertMessagesResponseBody {
	return s.Body
}

func (s *ListAlertMessagesResponse) SetHeaders(v map[string]*string) *ListAlertMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListAlertMessagesResponse) SetStatusCode(v int32) *ListAlertMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertMessagesResponse) SetBody(v *ListAlertMessagesResponseBody) *ListAlertMessagesResponse {
	s.Body = v
	return s
}

func (s *ListAlertMessagesResponse) Validate() error {
	return dara.Validate(s)
}
