// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateCdnUrlAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateCdnUrlAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateCdnUrlAuthResponse
	GetStatusCode() *int32
	SetBody(v *ValidateCdnUrlAuthResponseBody) *ValidateCdnUrlAuthResponse
	GetBody() *ValidateCdnUrlAuthResponseBody
}

type ValidateCdnUrlAuthResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateCdnUrlAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateCdnUrlAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateCdnUrlAuthResponse) GoString() string {
	return s.String()
}

func (s *ValidateCdnUrlAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateCdnUrlAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateCdnUrlAuthResponse) GetBody() *ValidateCdnUrlAuthResponseBody {
	return s.Body
}

func (s *ValidateCdnUrlAuthResponse) SetHeaders(v map[string]*string) *ValidateCdnUrlAuthResponse {
	s.Headers = v
	return s
}

func (s *ValidateCdnUrlAuthResponse) SetStatusCode(v int32) *ValidateCdnUrlAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateCdnUrlAuthResponse) SetBody(v *ValidateCdnUrlAuthResponseBody) *ValidateCdnUrlAuthResponse {
	s.Body = v
	return s
}

func (s *ValidateCdnUrlAuthResponse) Validate() error {
	return dara.Validate(s)
}
