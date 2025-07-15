// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListActionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListActionsResponse
	GetStatusCode() *int32
	SetBody(v *ListActionsResponseBody) *ListActionsResponse
	GetBody() *ListActionsResponseBody
}

type ListActionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListActionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListActionsResponse) GoString() string {
	return s.String()
}

func (s *ListActionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListActionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListActionsResponse) GetBody() *ListActionsResponseBody {
	return s.Body
}

func (s *ListActionsResponse) SetHeaders(v map[string]*string) *ListActionsResponse {
	s.Headers = v
	return s
}

func (s *ListActionsResponse) SetStatusCode(v int32) *ListActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActionsResponse) SetBody(v *ListActionsResponseBody) *ListActionsResponse {
	s.Body = v
	return s
}

func (s *ListActionsResponse) Validate() error {
	return dara.Validate(s)
}
