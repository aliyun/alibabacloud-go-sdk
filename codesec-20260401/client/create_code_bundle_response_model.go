// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeBundleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCodeBundleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCodeBundleResponse
	GetStatusCode() *int32
	SetBody(v *CreateCodeBundleResponseBody) *CreateCodeBundleResponse
	GetBody() *CreateCodeBundleResponseBody
}

type CreateCodeBundleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCodeBundleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCodeBundleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeBundleResponse) GoString() string {
	return s.String()
}

func (s *CreateCodeBundleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCodeBundleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCodeBundleResponse) GetBody() *CreateCodeBundleResponseBody {
	return s.Body
}

func (s *CreateCodeBundleResponse) SetHeaders(v map[string]*string) *CreateCodeBundleResponse {
	s.Headers = v
	return s
}

func (s *CreateCodeBundleResponse) SetStatusCode(v int32) *CreateCodeBundleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCodeBundleResponse) SetBody(v *CreateCodeBundleResponseBody) *CreateCodeBundleResponse {
	s.Body = v
	return s
}

func (s *CreateCodeBundleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
