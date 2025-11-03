// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIPv6TranslatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIPv6TranslatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIPv6TranslatorResponse
	GetStatusCode() *int32
	SetBody(v *CreateIPv6TranslatorResponseBody) *CreateIPv6TranslatorResponse
	GetBody() *CreateIPv6TranslatorResponseBody
}

type CreateIPv6TranslatorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIPv6TranslatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIPv6TranslatorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIPv6TranslatorResponse) GoString() string {
	return s.String()
}

func (s *CreateIPv6TranslatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIPv6TranslatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIPv6TranslatorResponse) GetBody() *CreateIPv6TranslatorResponseBody {
	return s.Body
}

func (s *CreateIPv6TranslatorResponse) SetHeaders(v map[string]*string) *CreateIPv6TranslatorResponse {
	s.Headers = v
	return s
}

func (s *CreateIPv6TranslatorResponse) SetStatusCode(v int32) *CreateIPv6TranslatorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIPv6TranslatorResponse) SetBody(v *CreateIPv6TranslatorResponseBody) *CreateIPv6TranslatorResponse {
	s.Body = v
	return s
}

func (s *CreateIPv6TranslatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
