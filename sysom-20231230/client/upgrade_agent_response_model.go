// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeAgentResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeAgentResponseBody) *UpgradeAgentResponse
	GetBody() *UpgradeAgentResponseBody
}

type UpgradeAgentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentResponse) GoString() string {
	return s.String()
}

func (s *UpgradeAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeAgentResponse) GetBody() *UpgradeAgentResponseBody {
	return s.Body
}

func (s *UpgradeAgentResponse) SetHeaders(v map[string]*string) *UpgradeAgentResponse {
	s.Headers = v
	return s
}

func (s *UpgradeAgentResponse) SetStatusCode(v int32) *UpgradeAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeAgentResponse) SetBody(v *UpgradeAgentResponseBody) *UpgradeAgentResponse {
	s.Body = v
	return s
}

func (s *UpgradeAgentResponse) Validate() error {
	return dara.Validate(s)
}
