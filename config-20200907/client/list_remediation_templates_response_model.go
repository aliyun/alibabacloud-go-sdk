// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRemediationTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRemediationTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListRemediationTemplatesResponseBody) *ListRemediationTemplatesResponse
	GetBody() *ListRemediationTemplatesResponseBody
}

type ListRemediationTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRemediationTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRemediationTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRemediationTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRemediationTemplatesResponse) GetBody() *ListRemediationTemplatesResponseBody {
	return s.Body
}

func (s *ListRemediationTemplatesResponse) SetHeaders(v map[string]*string) *ListRemediationTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListRemediationTemplatesResponse) SetStatusCode(v int32) *ListRemediationTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRemediationTemplatesResponse) SetBody(v *ListRemediationTemplatesResponseBody) *ListRemediationTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListRemediationTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
