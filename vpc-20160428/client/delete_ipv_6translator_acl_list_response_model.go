// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIPv6TranslatorAclListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIPv6TranslatorAclListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIPv6TranslatorAclListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIPv6TranslatorAclListResponseBody) *DeleteIPv6TranslatorAclListResponse
	GetBody() *DeleteIPv6TranslatorAclListResponseBody
}

type DeleteIPv6TranslatorAclListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIPv6TranslatorAclListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIPv6TranslatorAclListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIPv6TranslatorAclListResponse) GoString() string {
	return s.String()
}

func (s *DeleteIPv6TranslatorAclListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIPv6TranslatorAclListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIPv6TranslatorAclListResponse) GetBody() *DeleteIPv6TranslatorAclListResponseBody {
	return s.Body
}

func (s *DeleteIPv6TranslatorAclListResponse) SetHeaders(v map[string]*string) *DeleteIPv6TranslatorAclListResponse {
	s.Headers = v
	return s
}

func (s *DeleteIPv6TranslatorAclListResponse) SetStatusCode(v int32) *DeleteIPv6TranslatorAclListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIPv6TranslatorAclListResponse) SetBody(v *DeleteIPv6TranslatorAclListResponseBody) *DeleteIPv6TranslatorAclListResponse {
	s.Body = v
	return s
}

func (s *DeleteIPv6TranslatorAclListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
