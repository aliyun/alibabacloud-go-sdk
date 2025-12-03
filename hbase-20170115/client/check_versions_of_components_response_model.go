// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVersionsOfComponentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckVersionsOfComponentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckVersionsOfComponentsResponse
	GetStatusCode() *int32
	SetBody(v *CheckVersionsOfComponentsResponseBody) *CheckVersionsOfComponentsResponse
	GetBody() *CheckVersionsOfComponentsResponseBody
}

type CheckVersionsOfComponentsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckVersionsOfComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckVersionsOfComponentsResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckVersionsOfComponentsResponse) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckVersionsOfComponentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckVersionsOfComponentsResponse) GetBody() *CheckVersionsOfComponentsResponseBody {
	return s.Body
}

func (s *CheckVersionsOfComponentsResponse) SetHeaders(v map[string]*string) *CheckVersionsOfComponentsResponse {
	s.Headers = v
	return s
}

func (s *CheckVersionsOfComponentsResponse) SetStatusCode(v int32) *CheckVersionsOfComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckVersionsOfComponentsResponse) SetBody(v *CheckVersionsOfComponentsResponseBody) *CheckVersionsOfComponentsResponse {
	s.Body = v
	return s
}

func (s *CheckVersionsOfComponentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
