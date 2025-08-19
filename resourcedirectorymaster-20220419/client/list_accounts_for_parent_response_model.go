// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsForParentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccountsForParentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccountsForParentResponse
	GetStatusCode() *int32
	SetBody(v *ListAccountsForParentResponseBody) *ListAccountsForParentResponse
	GetBody() *ListAccountsForParentResponseBody
}

type ListAccountsForParentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsForParentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsForParentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsForParentResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccountsForParentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccountsForParentResponse) GetBody() *ListAccountsForParentResponseBody {
	return s.Body
}

func (s *ListAccountsForParentResponse) SetHeaders(v map[string]*string) *ListAccountsForParentResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsForParentResponse) SetStatusCode(v int32) *ListAccountsForParentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsForParentResponse) SetBody(v *ListAccountsForParentResponseBody) *ListAccountsForParentResponse {
	s.Body = v
	return s
}

func (s *ListAccountsForParentResponse) Validate() error {
	return dara.Validate(s)
}
