// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMinorVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeMinorVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeMinorVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeMinorVersionResponseBody) *UpgradeMinorVersionResponse
	GetBody() *UpgradeMinorVersionResponseBody
}

type UpgradeMinorVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeMinorVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeMinorVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeMinorVersionResponse) GetBody() *UpgradeMinorVersionResponseBody {
	return s.Body
}

func (s *UpgradeMinorVersionResponse) SetHeaders(v map[string]*string) *UpgradeMinorVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMinorVersionResponse) SetStatusCode(v int32) *UpgradeMinorVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMinorVersionResponse) SetBody(v *UpgradeMinorVersionResponseBody) *UpgradeMinorVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeMinorVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
