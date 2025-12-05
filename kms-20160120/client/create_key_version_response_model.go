// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKeyVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKeyVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateKeyVersionResponseBody) *CreateKeyVersionResponse
	GetBody() *CreateKeyVersionResponseBody
}

type CreateKeyVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeyVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKeyVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKeyVersionResponse) GetBody() *CreateKeyVersionResponseBody {
	return s.Body
}

func (s *CreateKeyVersionResponse) SetHeaders(v map[string]*string) *CreateKeyVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyVersionResponse) SetStatusCode(v int32) *CreateKeyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeyVersionResponse) SetBody(v *CreateKeyVersionResponseBody) *CreateKeyVersionResponse {
	s.Body = v
	return s
}

func (s *CreateKeyVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
