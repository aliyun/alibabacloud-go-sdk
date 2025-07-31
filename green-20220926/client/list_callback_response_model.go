// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCallbackResponse
	GetStatusCode() *int32
	SetBody(v *ListCallbackResponseBody) *ListCallbackResponse
	GetBody() *ListCallbackResponseBody
}

type ListCallbackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCallbackResponse) GoString() string {
	return s.String()
}

func (s *ListCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCallbackResponse) GetBody() *ListCallbackResponseBody {
	return s.Body
}

func (s *ListCallbackResponse) SetHeaders(v map[string]*string) *ListCallbackResponse {
	s.Headers = v
	return s
}

func (s *ListCallbackResponse) SetStatusCode(v int32) *ListCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallbackResponse) SetBody(v *ListCallbackResponseBody) *ListCallbackResponse {
	s.Body = v
	return s
}

func (s *ListCallbackResponse) Validate() error {
	return dara.Validate(s)
}
