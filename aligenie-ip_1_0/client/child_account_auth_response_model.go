// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChildAccountAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChildAccountAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChildAccountAuthResponse
	GetStatusCode() *int32
	SetBody(v *ChildAccountAuthResponseBody) *ChildAccountAuthResponse
	GetBody() *ChildAccountAuthResponseBody
}

type ChildAccountAuthResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChildAccountAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChildAccountAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s ChildAccountAuthResponse) GoString() string {
	return s.String()
}

func (s *ChildAccountAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChildAccountAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChildAccountAuthResponse) GetBody() *ChildAccountAuthResponseBody {
	return s.Body
}

func (s *ChildAccountAuthResponse) SetHeaders(v map[string]*string) *ChildAccountAuthResponse {
	s.Headers = v
	return s
}

func (s *ChildAccountAuthResponse) SetStatusCode(v int32) *ChildAccountAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *ChildAccountAuthResponse) SetBody(v *ChildAccountAuthResponseBody) *ChildAccountAuthResponse {
	s.Body = v
	return s
}

func (s *ChildAccountAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
