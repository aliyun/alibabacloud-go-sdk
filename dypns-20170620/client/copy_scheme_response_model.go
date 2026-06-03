// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopySchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopySchemeResponse
	GetStatusCode() *int32
	SetBody(v *CopySchemeResponseBody) *CopySchemeResponse
	GetBody() *CopySchemeResponseBody
}

type CopySchemeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopySchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s CopySchemeResponse) GoString() string {
	return s.String()
}

func (s *CopySchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopySchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopySchemeResponse) GetBody() *CopySchemeResponseBody {
	return s.Body
}

func (s *CopySchemeResponse) SetHeaders(v map[string]*string) *CopySchemeResponse {
	s.Headers = v
	return s
}

func (s *CopySchemeResponse) SetStatusCode(v int32) *CopySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CopySchemeResponse) SetBody(v *CopySchemeResponseBody) *CopySchemeResponse {
	s.Body = v
	return s
}

func (s *CopySchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
