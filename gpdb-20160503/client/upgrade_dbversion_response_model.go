// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBVersionResponseBody) *UpgradeDBVersionResponse
	GetBody() *UpgradeDBVersionResponseBody
}

type UpgradeDBVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBVersionResponse) GetBody() *UpgradeDBVersionResponseBody {
	return s.Body
}

func (s *UpgradeDBVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBVersionResponse) SetStatusCode(v int32) *UpgradeDBVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBVersionResponse) SetBody(v *UpgradeDBVersionResponseBody) *UpgradeDBVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBVersionResponse) Validate() error {
	return dara.Validate(s)
}
