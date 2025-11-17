// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListByUserGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListByUserGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListByUserGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *ListByUserGroupIdResponseBody) *ListByUserGroupIdResponse
	GetBody() *ListByUserGroupIdResponseBody
}

type ListByUserGroupIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListByUserGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListByUserGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListByUserGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListByUserGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListByUserGroupIdResponse) GetBody() *ListByUserGroupIdResponseBody {
	return s.Body
}

func (s *ListByUserGroupIdResponse) SetHeaders(v map[string]*string) *ListByUserGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListByUserGroupIdResponse) SetStatusCode(v int32) *ListByUserGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListByUserGroupIdResponse) SetBody(v *ListByUserGroupIdResponseBody) *ListByUserGroupIdResponse {
	s.Body = v
	return s
}

func (s *ListByUserGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
