// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAppInstanceGroupByUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedAppInstanceGroupByUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedAppInstanceGroupByUserResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedAppInstanceGroupByUserResponseBody) *ListAuthorizedAppInstanceGroupByUserResponse
	GetBody() *ListAuthorizedAppInstanceGroupByUserResponseBody
}

type ListAuthorizedAppInstanceGroupByUserResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedAppInstanceGroupByUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedAppInstanceGroupByUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAppInstanceGroupByUserResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAppInstanceGroupByUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedAppInstanceGroupByUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedAppInstanceGroupByUserResponse) GetBody() *ListAuthorizedAppInstanceGroupByUserResponseBody {
	return s.Body
}

func (s *ListAuthorizedAppInstanceGroupByUserResponse) SetHeaders(v map[string]*string) *ListAuthorizedAppInstanceGroupByUserResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponse) SetStatusCode(v int32) *ListAuthorizedAppInstanceGroupByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponse) SetBody(v *ListAuthorizedAppInstanceGroupByUserResponseBody) *ListAuthorizedAppInstanceGroupByUserResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
