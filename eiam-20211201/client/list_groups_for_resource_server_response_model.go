// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForResourceServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupsForResourceServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupsForResourceServerResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupsForResourceServerResponseBody) *ListGroupsForResourceServerResponse
	GetBody() *ListGroupsForResourceServerResponseBody
}

type ListGroupsForResourceServerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsForResourceServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsForResourceServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForResourceServerResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForResourceServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupsForResourceServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupsForResourceServerResponse) GetBody() *ListGroupsForResourceServerResponseBody {
	return s.Body
}

func (s *ListGroupsForResourceServerResponse) SetHeaders(v map[string]*string) *ListGroupsForResourceServerResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsForResourceServerResponse) SetStatusCode(v int32) *ListGroupsForResourceServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsForResourceServerResponse) SetBody(v *ListGroupsForResourceServerResponseBody) *ListGroupsForResourceServerResponse {
	s.Body = v
	return s
}

func (s *ListGroupsForResourceServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
