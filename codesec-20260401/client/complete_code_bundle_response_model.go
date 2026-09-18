// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteCodeBundleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteCodeBundleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteCodeBundleResponse
	GetStatusCode() *int32
	SetBody(v *CompleteCodeBundleResponseBody) *CompleteCodeBundleResponse
	GetBody() *CompleteCodeBundleResponseBody
}

type CompleteCodeBundleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompleteCodeBundleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteCodeBundleResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteCodeBundleResponse) GoString() string {
	return s.String()
}

func (s *CompleteCodeBundleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteCodeBundleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteCodeBundleResponse) GetBody() *CompleteCodeBundleResponseBody {
	return s.Body
}

func (s *CompleteCodeBundleResponse) SetHeaders(v map[string]*string) *CompleteCodeBundleResponse {
	s.Headers = v
	return s
}

func (s *CompleteCodeBundleResponse) SetStatusCode(v int32) *CompleteCodeBundleResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteCodeBundleResponse) SetBody(v *CompleteCodeBundleResponseBody) *CompleteCodeBundleResponse {
	s.Body = v
	return s
}

func (s *CompleteCodeBundleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
