// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeRenderingInstanceImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeRenderingInstanceImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeRenderingInstanceImageResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeRenderingInstanceImageResponseBody) *UpgradeRenderingInstanceImageResponse
	GetBody() *UpgradeRenderingInstanceImageResponseBody
}

type UpgradeRenderingInstanceImageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeRenderingInstanceImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeRenderingInstanceImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingInstanceImageResponse) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingInstanceImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeRenderingInstanceImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeRenderingInstanceImageResponse) GetBody() *UpgradeRenderingInstanceImageResponseBody {
	return s.Body
}

func (s *UpgradeRenderingInstanceImageResponse) SetHeaders(v map[string]*string) *UpgradeRenderingInstanceImageResponse {
	s.Headers = v
	return s
}

func (s *UpgradeRenderingInstanceImageResponse) SetStatusCode(v int32) *UpgradeRenderingInstanceImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeRenderingInstanceImageResponse) SetBody(v *UpgradeRenderingInstanceImageResponseBody) *UpgradeRenderingInstanceImageResponse {
	s.Body = v
	return s
}

func (s *UpgradeRenderingInstanceImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
