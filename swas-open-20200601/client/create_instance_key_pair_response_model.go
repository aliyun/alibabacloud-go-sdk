// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceKeyPairResponseBody) *CreateInstanceKeyPairResponse
	GetBody() *CreateInstanceKeyPairResponseBody
}

type CreateInstanceKeyPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceKeyPairResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceKeyPairResponse) GetBody() *CreateInstanceKeyPairResponseBody {
	return s.Body
}

func (s *CreateInstanceKeyPairResponse) SetHeaders(v map[string]*string) *CreateInstanceKeyPairResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceKeyPairResponse) SetStatusCode(v int32) *CreateInstanceKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceKeyPairResponse) SetBody(v *CreateInstanceKeyPairResponseBody) *CreateInstanceKeyPairResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
