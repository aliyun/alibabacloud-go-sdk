// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalSourcesByObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPersonalSourcesByObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPersonalSourcesByObjectResponse
	GetStatusCode() *int32
	SetBody(v *ListPersonalSourcesByObjectResponseBody) *ListPersonalSourcesByObjectResponse
	GetBody() *ListPersonalSourcesByObjectResponseBody
}

type ListPersonalSourcesByObjectResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPersonalSourcesByObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPersonalSourcesByObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalSourcesByObjectResponse) GoString() string {
	return s.String()
}

func (s *ListPersonalSourcesByObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPersonalSourcesByObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPersonalSourcesByObjectResponse) GetBody() *ListPersonalSourcesByObjectResponseBody {
	return s.Body
}

func (s *ListPersonalSourcesByObjectResponse) SetHeaders(v map[string]*string) *ListPersonalSourcesByObjectResponse {
	s.Headers = v
	return s
}

func (s *ListPersonalSourcesByObjectResponse) SetStatusCode(v int32) *ListPersonalSourcesByObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponse) SetBody(v *ListPersonalSourcesByObjectResponseBody) *ListPersonalSourcesByObjectResponse {
	s.Body = v
	return s
}

func (s *ListPersonalSourcesByObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
