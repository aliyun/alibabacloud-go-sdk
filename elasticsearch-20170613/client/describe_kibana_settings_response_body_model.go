// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKibanaSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeKibanaSettingsResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DescribeKibanaSettingsResponseBody
	GetResult() map[string]interface{}
}

type DescribeKibanaSettingsResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6D*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {"map.includeElasticMapsService": "false", "server.ssl.cert": "/home/admin/packages/kibana/config/cert/client.crt", "server.ssl.enabled": "true", "server.ssl.key": "/home/admin/packages/kibana/config/cert/client.key", "xpack.reporting.capture.browser.chromium.disableSandbox": "true"}
	Result map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeKibanaSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKibanaSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKibanaSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKibanaSettingsResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DescribeKibanaSettingsResponseBody) SetRequestId(v string) *DescribeKibanaSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKibanaSettingsResponseBody) SetResult(v map[string]interface{}) *DescribeKibanaSettingsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeKibanaSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
