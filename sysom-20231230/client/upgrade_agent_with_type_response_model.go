// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentWithTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeAgentWithTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeAgentWithTypeResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeAgentWithTypeResponseBody) *UpgradeAgentWithTypeResponse
	GetBody() *UpgradeAgentWithTypeResponseBody
}

type UpgradeAgentWithTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeAgentWithTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeAgentWithTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentWithTypeResponse) GoString() string {
	return s.String()
}

func (s *UpgradeAgentWithTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeAgentWithTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeAgentWithTypeResponse) GetBody() *UpgradeAgentWithTypeResponseBody {
	return s.Body
}

func (s *UpgradeAgentWithTypeResponse) SetHeaders(v map[string]*string) *UpgradeAgentWithTypeResponse {
	s.Headers = v
	return s
}

func (s *UpgradeAgentWithTypeResponse) SetStatusCode(v int32) *UpgradeAgentWithTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeAgentWithTypeResponse) SetBody(v *UpgradeAgentWithTypeResponseBody) *UpgradeAgentWithTypeResponse {
	s.Body = v
	return s
}

func (s *UpgradeAgentWithTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
