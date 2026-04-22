// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterprisePptTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnterprisePptTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnterprisePptTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListEnterprisePptTemplatesResponseBody) *ListEnterprisePptTemplatesResponse
	GetBody() *ListEnterprisePptTemplatesResponseBody
}

type ListEnterprisePptTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterprisePptTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterprisePptTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnterprisePptTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListEnterprisePptTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnterprisePptTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnterprisePptTemplatesResponse) GetBody() *ListEnterprisePptTemplatesResponseBody {
	return s.Body
}

func (s *ListEnterprisePptTemplatesResponse) SetHeaders(v map[string]*string) *ListEnterprisePptTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListEnterprisePptTemplatesResponse) SetStatusCode(v int32) *ListEnterprisePptTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponse) SetBody(v *ListEnterprisePptTemplatesResponseBody) *ListEnterprisePptTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListEnterprisePptTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
