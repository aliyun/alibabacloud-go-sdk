// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateNetwrokResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrivateNetwrokResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrivateNetwrokResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrivateNetwrokResponseBody) *UpdatePrivateNetwrokResponse
	GetBody() *UpdatePrivateNetwrokResponseBody
}

type UpdatePrivateNetwrokResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateNetwrokResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateNetwrokResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetwrokResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetwrokResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrivateNetwrokResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrivateNetwrokResponse) GetBody() *UpdatePrivateNetwrokResponseBody {
	return s.Body
}

func (s *UpdatePrivateNetwrokResponse) SetHeaders(v map[string]*string) *UpdatePrivateNetwrokResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateNetwrokResponse) SetStatusCode(v int32) *UpdatePrivateNetwrokResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateNetwrokResponse) SetBody(v *UpdatePrivateNetwrokResponseBody) *UpdatePrivateNetwrokResponse {
	s.Body = v
	return s
}

func (s *UpdatePrivateNetwrokResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
