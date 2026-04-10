// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSandboxTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSandboxTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSandboxTemplatesResponseBody) *DescribeSandboxTemplatesResponse
	GetBody() *DescribeSandboxTemplatesResponseBody
}

type DescribeSandboxTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSandboxTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSandboxTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSandboxTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSandboxTemplatesResponse) GetBody() *DescribeSandboxTemplatesResponseBody {
	return s.Body
}

func (s *DescribeSandboxTemplatesResponse) SetHeaders(v map[string]*string) *DescribeSandboxTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxTemplatesResponse) SetStatusCode(v int32) *DescribeSandboxTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxTemplatesResponse) SetBody(v *DescribeSandboxTemplatesResponseBody) *DescribeSandboxTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeSandboxTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
