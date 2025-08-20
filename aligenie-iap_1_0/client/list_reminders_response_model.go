// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemindersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRemindersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRemindersResponse
	GetStatusCode() *int32
	SetBody(v *ListRemindersResponseBody) *ListRemindersResponse
	GetBody() *ListRemindersResponseBody
}

type ListRemindersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRemindersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRemindersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersResponse) GoString() string {
	return s.String()
}

func (s *ListRemindersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRemindersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRemindersResponse) GetBody() *ListRemindersResponseBody {
	return s.Body
}

func (s *ListRemindersResponse) SetHeaders(v map[string]*string) *ListRemindersResponse {
	s.Headers = v
	return s
}

func (s *ListRemindersResponse) SetStatusCode(v int32) *ListRemindersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRemindersResponse) SetBody(v *ListRemindersResponseBody) *ListRemindersResponse {
	s.Body = v
	return s
}

func (s *ListRemindersResponse) Validate() error {
	return dara.Validate(s)
}
