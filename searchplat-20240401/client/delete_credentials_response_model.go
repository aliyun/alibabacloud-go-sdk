// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCredentialsResponseBody) *DeleteCredentialsResponse
	GetBody() *DeleteCredentialsResponseBody
}

type DeleteCredentialsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialsResponse) GoString() string {
	return s.String()
}

func (s *DeleteCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCredentialsResponse) GetBody() *DeleteCredentialsResponseBody {
	return s.Body
}

func (s *DeleteCredentialsResponse) SetHeaders(v map[string]*string) *DeleteCredentialsResponse {
	s.Headers = v
	return s
}

func (s *DeleteCredentialsResponse) SetStatusCode(v int32) *DeleteCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCredentialsResponse) SetBody(v *DeleteCredentialsResponseBody) *DeleteCredentialsResponse {
	s.Body = v
	return s
}

func (s *DeleteCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
