// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContactGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContactGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListContactGroupsResponseBody) *ListContactGroupsResponse
	GetBody() *ListContactGroupsResponseBody
}

type ListContactGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContactGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContactGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContactGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListContactGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContactGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContactGroupsResponse) GetBody() *ListContactGroupsResponseBody {
	return s.Body
}

func (s *ListContactGroupsResponse) SetHeaders(v map[string]*string) *ListContactGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListContactGroupsResponse) SetStatusCode(v int32) *ListContactGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContactGroupsResponse) SetBody(v *ListContactGroupsResponseBody) *ListContactGroupsResponse {
	s.Body = v
	return s
}

func (s *ListContactGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
