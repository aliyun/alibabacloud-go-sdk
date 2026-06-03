// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyVerifySchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyVerifySchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyVerifySchemeResponse
	GetStatusCode() *int32
	SetBody(v *CopyVerifySchemeResponseBody) *CopyVerifySchemeResponse
	GetBody() *CopyVerifySchemeResponseBody
}

type CopyVerifySchemeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyVerifySchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *CopyVerifySchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyVerifySchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyVerifySchemeResponse) GetBody() *CopyVerifySchemeResponseBody {
	return s.Body
}

func (s *CopyVerifySchemeResponse) SetHeaders(v map[string]*string) *CopyVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *CopyVerifySchemeResponse) SetStatusCode(v int32) *CopyVerifySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyVerifySchemeResponse) SetBody(v *CopyVerifySchemeResponseBody) *CopyVerifySchemeResponse {
	s.Body = v
	return s
}

func (s *CopyVerifySchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
