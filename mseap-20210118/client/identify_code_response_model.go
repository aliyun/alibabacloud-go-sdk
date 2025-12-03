// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdentifyCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IdentifyCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IdentifyCodeResponse
	GetStatusCode() *int32
	SetBody(v *IdentifyCodeResponseBody) *IdentifyCodeResponse
	GetBody() *IdentifyCodeResponseBody
}

type IdentifyCodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IdentifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IdentifyCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s IdentifyCodeResponse) GoString() string {
	return s.String()
}

func (s *IdentifyCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IdentifyCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IdentifyCodeResponse) GetBody() *IdentifyCodeResponseBody {
	return s.Body
}

func (s *IdentifyCodeResponse) SetHeaders(v map[string]*string) *IdentifyCodeResponse {
	s.Headers = v
	return s
}

func (s *IdentifyCodeResponse) SetStatusCode(v int32) *IdentifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *IdentifyCodeResponse) SetBody(v *IdentifyCodeResponseBody) *IdentifyCodeResponse {
	s.Body = v
	return s
}

func (s *IdentifyCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
