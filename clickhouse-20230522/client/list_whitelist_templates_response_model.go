// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWhitelistTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWhitelistTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWhitelistTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListWhitelistTemplatesResponseBody) *ListWhitelistTemplatesResponse
	GetBody() *ListWhitelistTemplatesResponseBody
}

type ListWhitelistTemplatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWhitelistTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWhitelistTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWhitelistTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListWhitelistTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWhitelistTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWhitelistTemplatesResponse) GetBody() *ListWhitelistTemplatesResponseBody {
	return s.Body
}

func (s *ListWhitelistTemplatesResponse) SetHeaders(v map[string]*string) *ListWhitelistTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListWhitelistTemplatesResponse) SetStatusCode(v int32) *ListWhitelistTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWhitelistTemplatesResponse) SetBody(v *ListWhitelistTemplatesResponseBody) *ListWhitelistTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListWhitelistTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
