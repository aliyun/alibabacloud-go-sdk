// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAICInstanceImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeAICInstanceImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeAICInstanceImageResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeAICInstanceImageResponseBody) *UpgradeAICInstanceImageResponse
	GetBody() *UpgradeAICInstanceImageResponseBody
}

type UpgradeAICInstanceImageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeAICInstanceImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeAICInstanceImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAICInstanceImageResponse) GoString() string {
	return s.String()
}

func (s *UpgradeAICInstanceImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeAICInstanceImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeAICInstanceImageResponse) GetBody() *UpgradeAICInstanceImageResponseBody {
	return s.Body
}

func (s *UpgradeAICInstanceImageResponse) SetHeaders(v map[string]*string) *UpgradeAICInstanceImageResponse {
	s.Headers = v
	return s
}

func (s *UpgradeAICInstanceImageResponse) SetStatusCode(v int32) *UpgradeAICInstanceImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeAICInstanceImageResponse) SetBody(v *UpgradeAICInstanceImageResponseBody) *UpgradeAICInstanceImageResponse {
	s.Body = v
	return s
}

func (s *UpgradeAICInstanceImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
