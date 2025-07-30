// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupsForApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupsForApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupsForApplicationResponseBody) *ListGroupsForApplicationResponse
	GetBody() *ListGroupsForApplicationResponseBody
}

type ListGroupsForApplicationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsForApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsForApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupsForApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupsForApplicationResponse) GetBody() *ListGroupsForApplicationResponseBody {
	return s.Body
}

func (s *ListGroupsForApplicationResponse) SetHeaders(v map[string]*string) *ListGroupsForApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsForApplicationResponse) SetStatusCode(v int32) *ListGroupsForApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsForApplicationResponse) SetBody(v *ListGroupsForApplicationResponseBody) *ListGroupsForApplicationResponse {
	s.Body = v
	return s
}

func (s *ListGroupsForApplicationResponse) Validate() error {
	return dara.Validate(s)
}
