// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradePolarClawPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradePolarClawPluginResponse
	GetStatusCode() *int32
	SetBody(v *UpgradePolarClawPluginResponseBody) *UpgradePolarClawPluginResponse
	GetBody() *UpgradePolarClawPluginResponseBody
}

type UpgradePolarClawPluginResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradePolarClawPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradePolarClawPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawPluginResponse) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradePolarClawPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradePolarClawPluginResponse) GetBody() *UpgradePolarClawPluginResponseBody {
	return s.Body
}

func (s *UpgradePolarClawPluginResponse) SetHeaders(v map[string]*string) *UpgradePolarClawPluginResponse {
	s.Headers = v
	return s
}

func (s *UpgradePolarClawPluginResponse) SetStatusCode(v int32) *UpgradePolarClawPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradePolarClawPluginResponse) SetBody(v *UpgradePolarClawPluginResponseBody) *UpgradePolarClawPluginResponse {
	s.Body = v
	return s
}

func (s *UpgradePolarClawPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
