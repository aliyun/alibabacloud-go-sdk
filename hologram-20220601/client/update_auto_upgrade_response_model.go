// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutoUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutoUpgradeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutoUpgradeResponseBody) *UpdateAutoUpgradeResponse
	GetBody() *UpdateAutoUpgradeResponseBody
}

type UpdateAutoUpgradeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoUpgradeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutoUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutoUpgradeResponse) GetBody() *UpdateAutoUpgradeResponseBody {
	return s.Body
}

func (s *UpdateAutoUpgradeResponse) SetHeaders(v map[string]*string) *UpdateAutoUpgradeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoUpgradeResponse) SetStatusCode(v int32) *UpdateAutoUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoUpgradeResponse) SetBody(v *UpdateAutoUpgradeResponseBody) *UpdateAutoUpgradeResponse {
	s.Body = v
	return s
}

func (s *UpdateAutoUpgradeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
