// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceCredentialsResponseBody) *ListServiceCredentialsResponse
	GetBody() *ListServiceCredentialsResponseBody
}

type ListServiceCredentialsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceCredentialsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceCredentialsResponse) GetBody() *ListServiceCredentialsResponseBody {
	return s.Body
}

func (s *ListServiceCredentialsResponse) SetHeaders(v map[string]*string) *ListServiceCredentialsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceCredentialsResponse) SetStatusCode(v int32) *ListServiceCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceCredentialsResponse) SetBody(v *ListServiceCredentialsResponseBody) *ListServiceCredentialsResponse {
	s.Body = v
	return s
}

func (s *ListServiceCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
