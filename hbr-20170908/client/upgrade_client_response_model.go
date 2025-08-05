// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeClientResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeClientResponseBody) *UpgradeClientResponse
	GetBody() *UpgradeClientResponseBody
}

type UpgradeClientResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeClientResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClientResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeClientResponse) GetBody() *UpgradeClientResponseBody {
	return s.Body
}

func (s *UpgradeClientResponse) SetHeaders(v map[string]*string) *UpgradeClientResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClientResponse) SetStatusCode(v int32) *UpgradeClientResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeClientResponse) SetBody(v *UpgradeClientResponseBody) *UpgradeClientResponse {
	s.Body = v
	return s
}

func (s *UpgradeClientResponse) Validate() error {
	return dara.Validate(s)
}
