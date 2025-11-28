// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgreeMemberAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AgreeMemberAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AgreeMemberAccessResponse
	GetStatusCode() *int32
	SetBody(v *AgreeMemberAccessResponseBody) *AgreeMemberAccessResponse
	GetBody() *AgreeMemberAccessResponseBody
}

type AgreeMemberAccessResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgreeMemberAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgreeMemberAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s AgreeMemberAccessResponse) GoString() string {
	return s.String()
}

func (s *AgreeMemberAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AgreeMemberAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AgreeMemberAccessResponse) GetBody() *AgreeMemberAccessResponseBody {
	return s.Body
}

func (s *AgreeMemberAccessResponse) SetHeaders(v map[string]*string) *AgreeMemberAccessResponse {
	s.Headers = v
	return s
}

func (s *AgreeMemberAccessResponse) SetStatusCode(v int32) *AgreeMemberAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *AgreeMemberAccessResponse) SetBody(v *AgreeMemberAccessResponseBody) *AgreeMemberAccessResponse {
	s.Body = v
	return s
}

func (s *AgreeMemberAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
