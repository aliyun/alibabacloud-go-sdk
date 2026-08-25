// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSCIMServerCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSCIMServerCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSCIMServerCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *ListSCIMServerCredentialsResponseBody) *ListSCIMServerCredentialsResponse
	GetBody() *ListSCIMServerCredentialsResponseBody
}

type ListSCIMServerCredentialsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSCIMServerCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSCIMServerCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSCIMServerCredentialsResponse) GoString() string {
	return s.String()
}

func (s *ListSCIMServerCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSCIMServerCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSCIMServerCredentialsResponse) GetBody() *ListSCIMServerCredentialsResponseBody {
	return s.Body
}

func (s *ListSCIMServerCredentialsResponse) SetHeaders(v map[string]*string) *ListSCIMServerCredentialsResponse {
	s.Headers = v
	return s
}

func (s *ListSCIMServerCredentialsResponse) SetStatusCode(v int32) *ListSCIMServerCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSCIMServerCredentialsResponse) SetBody(v *ListSCIMServerCredentialsResponseBody) *ListSCIMServerCredentialsResponse {
	s.Body = v
	return s
}

func (s *ListSCIMServerCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
