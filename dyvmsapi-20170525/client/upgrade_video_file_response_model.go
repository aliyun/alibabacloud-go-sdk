// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeVideoFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeVideoFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeVideoFileResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeVideoFileResponseBody) *UpgradeVideoFileResponse
	GetBody() *UpgradeVideoFileResponseBody
}

type UpgradeVideoFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeVideoFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeVideoFileResponse) GoString() string {
	return s.String()
}

func (s *UpgradeVideoFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeVideoFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeVideoFileResponse) GetBody() *UpgradeVideoFileResponseBody {
	return s.Body
}

func (s *UpgradeVideoFileResponse) SetHeaders(v map[string]*string) *UpgradeVideoFileResponse {
	s.Headers = v
	return s
}

func (s *UpgradeVideoFileResponse) SetStatusCode(v int32) *UpgradeVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeVideoFileResponse) SetBody(v *UpgradeVideoFileResponseBody) *UpgradeVideoFileResponse {
	s.Body = v
	return s
}

func (s *UpgradeVideoFileResponse) Validate() error {
	return dara.Validate(s)
}
