// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindIdpListByLoginIdentifierResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindIdpListByLoginIdentifierResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindIdpListByLoginIdentifierResponse
	GetStatusCode() *int32
	SetBody(v *FindIdpListByLoginIdentifierResponseBody) *FindIdpListByLoginIdentifierResponse
	GetBody() *FindIdpListByLoginIdentifierResponseBody
}

type FindIdpListByLoginIdentifierResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindIdpListByLoginIdentifierResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindIdpListByLoginIdentifierResponse) String() string {
	return dara.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponse) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindIdpListByLoginIdentifierResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindIdpListByLoginIdentifierResponse) GetBody() *FindIdpListByLoginIdentifierResponseBody {
	return s.Body
}

func (s *FindIdpListByLoginIdentifierResponse) SetHeaders(v map[string]*string) *FindIdpListByLoginIdentifierResponse {
	s.Headers = v
	return s
}

func (s *FindIdpListByLoginIdentifierResponse) SetStatusCode(v int32) *FindIdpListByLoginIdentifierResponse {
	s.StatusCode = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponse) SetBody(v *FindIdpListByLoginIdentifierResponseBody) *FindIdpListByLoginIdentifierResponse {
	s.Body = v
	return s
}

func (s *FindIdpListByLoginIdentifierResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
