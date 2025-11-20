// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *ListCredentialsResult) *ListCredentialsResponse
	GetBody() *ListCredentialsResult
}

type ListCredentialsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCredentialsResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponse) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCredentialsResponse) GetBody() *ListCredentialsResult {
	return s.Body
}

func (s *ListCredentialsResponse) SetHeaders(v map[string]*string) *ListCredentialsResponse {
	s.Headers = v
	return s
}

func (s *ListCredentialsResponse) SetStatusCode(v int32) *ListCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCredentialsResponse) SetBody(v *ListCredentialsResult) *ListCredentialsResponse {
	s.Body = v
	return s
}

func (s *ListCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
