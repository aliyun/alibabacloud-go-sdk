// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradePolarClawChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradePolarClawChannelResponse
	GetStatusCode() *int32
	SetBody(v *UpgradePolarClawChannelResponseBody) *UpgradePolarClawChannelResponse
	GetBody() *UpgradePolarClawChannelResponseBody
}

type UpgradePolarClawChannelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradePolarClawChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradePolarClawChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawChannelResponse) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradePolarClawChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradePolarClawChannelResponse) GetBody() *UpgradePolarClawChannelResponseBody {
	return s.Body
}

func (s *UpgradePolarClawChannelResponse) SetHeaders(v map[string]*string) *UpgradePolarClawChannelResponse {
	s.Headers = v
	return s
}

func (s *UpgradePolarClawChannelResponse) SetStatusCode(v int32) *UpgradePolarClawChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradePolarClawChannelResponse) SetBody(v *UpgradePolarClawChannelResponseBody) *UpgradePolarClawChannelResponse {
	s.Body = v
	return s
}

func (s *UpgradePolarClawChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
