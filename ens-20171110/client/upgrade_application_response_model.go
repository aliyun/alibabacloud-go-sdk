// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeApplicationResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeApplicationResponseBody) *UpgradeApplicationResponse
	GetBody() *UpgradeApplicationResponseBody
}

type UpgradeApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeApplicationResponse) GetBody() *UpgradeApplicationResponseBody {
	return s.Body
}

func (s *UpgradeApplicationResponse) SetHeaders(v map[string]*string) *UpgradeApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpgradeApplicationResponse) SetStatusCode(v int32) *UpgradeApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeApplicationResponse) SetBody(v *UpgradeApplicationResponseBody) *UpgradeApplicationResponse {
	s.Body = v
	return s
}

func (s *UpgradeApplicationResponse) Validate() error {
	return dara.Validate(s)
}
