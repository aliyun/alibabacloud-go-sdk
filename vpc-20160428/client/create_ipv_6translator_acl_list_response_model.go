// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIPv6TranslatorAclListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIPv6TranslatorAclListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIPv6TranslatorAclListResponse
	GetStatusCode() *int32
	SetBody(v *CreateIPv6TranslatorAclListResponseBody) *CreateIPv6TranslatorAclListResponse
	GetBody() *CreateIPv6TranslatorAclListResponseBody
}

type CreateIPv6TranslatorAclListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIPv6TranslatorAclListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIPv6TranslatorAclListResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIPv6TranslatorAclListResponse) GoString() string {
	return s.String()
}

func (s *CreateIPv6TranslatorAclListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIPv6TranslatorAclListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIPv6TranslatorAclListResponse) GetBody() *CreateIPv6TranslatorAclListResponseBody {
	return s.Body
}

func (s *CreateIPv6TranslatorAclListResponse) SetHeaders(v map[string]*string) *CreateIPv6TranslatorAclListResponse {
	s.Headers = v
	return s
}

func (s *CreateIPv6TranslatorAclListResponse) SetStatusCode(v int32) *CreateIPv6TranslatorAclListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIPv6TranslatorAclListResponse) SetBody(v *CreateIPv6TranslatorAclListResponseBody) *CreateIPv6TranslatorAclListResponse {
	s.Body = v
	return s
}

func (s *CreateIPv6TranslatorAclListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
