// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeInstanceResponseBody) *UpgradeInstanceResponse
	GetBody() *UpgradeInstanceResponseBody
}

type UpgradeInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeInstanceResponse) GetBody() *UpgradeInstanceResponseBody {
	return s.Body
}

func (s *UpgradeInstanceResponse) SetHeaders(v map[string]*string) *UpgradeInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceResponse) SetStatusCode(v int32) *UpgradeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeInstanceResponse) SetBody(v *UpgradeInstanceResponseBody) *UpgradeInstanceResponse {
	s.Body = v
	return s
}

func (s *UpgradeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
