// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeRenderingServerImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeRenderingServerImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeRenderingServerImageResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeRenderingServerImageResponseBody) *UpgradeRenderingServerImageResponse
	GetBody() *UpgradeRenderingServerImageResponseBody
}

type UpgradeRenderingServerImageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeRenderingServerImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeRenderingServerImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingServerImageResponse) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingServerImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeRenderingServerImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeRenderingServerImageResponse) GetBody() *UpgradeRenderingServerImageResponseBody {
	return s.Body
}

func (s *UpgradeRenderingServerImageResponse) SetHeaders(v map[string]*string) *UpgradeRenderingServerImageResponse {
	s.Headers = v
	return s
}

func (s *UpgradeRenderingServerImageResponse) SetStatusCode(v int32) *UpgradeRenderingServerImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeRenderingServerImageResponse) SetBody(v *UpgradeRenderingServerImageResponseBody) *UpgradeRenderingServerImageResponse {
	s.Body = v
	return s
}

func (s *UpgradeRenderingServerImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
