// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePxfsVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradePxfsVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradePxfsVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradePxfsVersionResponseBody) *UpgradePxfsVersionResponse
	GetBody() *UpgradePxfsVersionResponseBody
}

type UpgradePxfsVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradePxfsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradePxfsVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradePxfsVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradePxfsVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradePxfsVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradePxfsVersionResponse) GetBody() *UpgradePxfsVersionResponseBody {
	return s.Body
}

func (s *UpgradePxfsVersionResponse) SetHeaders(v map[string]*string) *UpgradePxfsVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradePxfsVersionResponse) SetStatusCode(v int32) *UpgradePxfsVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradePxfsVersionResponse) SetBody(v *UpgradePxfsVersionResponseBody) *UpgradePxfsVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradePxfsVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
