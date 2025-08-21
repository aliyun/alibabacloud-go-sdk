// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeEngineVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeEngineVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeEngineVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeEngineVersionResponseBody) *UpgradeEngineVersionResponse
	GetBody() *UpgradeEngineVersionResponseBody
}

type UpgradeEngineVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeEngineVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeEngineVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEngineVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeEngineVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeEngineVersionResponse) GetBody() *UpgradeEngineVersionResponseBody {
	return s.Body
}

func (s *UpgradeEngineVersionResponse) SetHeaders(v map[string]*string) *UpgradeEngineVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeEngineVersionResponse) SetStatusCode(v int32) *UpgradeEngineVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeEngineVersionResponse) SetBody(v *UpgradeEngineVersionResponseBody) *UpgradeEngineVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeEngineVersionResponse) Validate() error {
	return dara.Validate(s)
}
