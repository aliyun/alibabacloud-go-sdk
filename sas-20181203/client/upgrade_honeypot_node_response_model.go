// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeHoneypotNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeHoneypotNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeHoneypotNodeResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeHoneypotNodeResponseBody) *UpgradeHoneypotNodeResponse
	GetBody() *UpgradeHoneypotNodeResponseBody
}

type UpgradeHoneypotNodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeHoneypotNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeHoneypotNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeHoneypotNodeResponse) GoString() string {
	return s.String()
}

func (s *UpgradeHoneypotNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeHoneypotNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeHoneypotNodeResponse) GetBody() *UpgradeHoneypotNodeResponseBody {
	return s.Body
}

func (s *UpgradeHoneypotNodeResponse) SetHeaders(v map[string]*string) *UpgradeHoneypotNodeResponse {
	s.Headers = v
	return s
}

func (s *UpgradeHoneypotNodeResponse) SetStatusCode(v int32) *UpgradeHoneypotNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeHoneypotNodeResponse) SetBody(v *UpgradeHoneypotNodeResponseBody) *UpgradeHoneypotNodeResponse {
	s.Body = v
	return s
}

func (s *UpgradeHoneypotNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
