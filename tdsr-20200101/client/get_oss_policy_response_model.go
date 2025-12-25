// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetOssPolicyResponseBody) *GetOssPolicyResponse
	GetBody() *GetOssPolicyResponseBody
}

type GetOssPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetOssPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssPolicyResponse) GetBody() *GetOssPolicyResponseBody {
	return s.Body
}

func (s *GetOssPolicyResponse) SetHeaders(v map[string]*string) *GetOssPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetOssPolicyResponse) SetStatusCode(v int32) *GetOssPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssPolicyResponse) SetBody(v *GetOssPolicyResponseBody) *GetOssPolicyResponse {
	s.Body = v
	return s
}

func (s *GetOssPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
