// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeVersionResponseBody) *UpgradeVersionResponse
	GetBody() *UpgradeVersionResponseBody
}

type UpgradeVersionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeVersionResponse) GetBody() *UpgradeVersionResponseBody {
	return s.Body
}

func (s *UpgradeVersionResponse) SetHeaders(v map[string]*string) *UpgradeVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeVersionResponse) SetStatusCode(v int32) *UpgradeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeVersionResponse) SetBody(v *UpgradeVersionResponseBody) *UpgradeVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
