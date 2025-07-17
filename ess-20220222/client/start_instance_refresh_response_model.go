// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRefreshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartInstanceRefreshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartInstanceRefreshResponse
	GetStatusCode() *int32
	SetBody(v *StartInstanceRefreshResponseBody) *StartInstanceRefreshResponse
	GetBody() *StartInstanceRefreshResponseBody
}

type StartInstanceRefreshResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartInstanceRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartInstanceRefreshResponse) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRefreshResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceRefreshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartInstanceRefreshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartInstanceRefreshResponse) GetBody() *StartInstanceRefreshResponseBody {
	return s.Body
}

func (s *StartInstanceRefreshResponse) SetHeaders(v map[string]*string) *StartInstanceRefreshResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceRefreshResponse) SetStatusCode(v int32) *StartInstanceRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceRefreshResponse) SetBody(v *StartInstanceRefreshResponseBody) *StartInstanceRefreshResponse {
	s.Body = v
	return s
}

func (s *StartInstanceRefreshResponse) Validate() error {
	return dara.Validate(s)
}
