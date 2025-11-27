// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInstanceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeInstanceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeInstanceVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeInstanceVersionResponseBody) *UpgradeInstanceVersionResponse
	GetBody() *UpgradeInstanceVersionResponseBody
}

type UpgradeInstanceVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeInstanceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeInstanceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeInstanceVersionResponse) GetBody() *UpgradeInstanceVersionResponseBody {
	return s.Body
}

func (s *UpgradeInstanceVersionResponse) SetHeaders(v map[string]*string) *UpgradeInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceVersionResponse) SetStatusCode(v int32) *UpgradeInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeInstanceVersionResponse) SetBody(v *UpgradeInstanceVersionResponseBody) *UpgradeInstanceVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeInstanceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
