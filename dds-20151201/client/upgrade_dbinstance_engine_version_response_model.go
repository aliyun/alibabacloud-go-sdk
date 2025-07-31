// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceEngineVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBInstanceEngineVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBInstanceEngineVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBInstanceEngineVersionResponseBody) *UpgradeDBInstanceEngineVersionResponse
	GetBody() *UpgradeDBInstanceEngineVersionResponseBody
}

type UpgradeDBInstanceEngineVersionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBInstanceEngineVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBInstanceEngineVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBInstanceEngineVersionResponse) GetBody() *UpgradeDBInstanceEngineVersionResponseBody {
	return s.Body
}

func (s *UpgradeDBInstanceEngineVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceEngineVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponse) SetStatusCode(v int32) *UpgradeDBInstanceEngineVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponse) SetBody(v *UpgradeDBInstanceEngineVersionResponseBody) *UpgradeDBInstanceEngineVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponse) Validate() error {
	return dara.Validate(s)
}
