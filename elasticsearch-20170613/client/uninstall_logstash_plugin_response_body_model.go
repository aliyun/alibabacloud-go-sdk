// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallLogstashPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]interface{}) *UninstallLogstashPluginResponseBody
	GetHeaders() map[string]interface{}
	SetRequestId(v string) *UninstallLogstashPluginResponseBody
	GetRequestId() *string
	SetResult(v []*string) *UninstallLogstashPluginResponseBody
	GetResult() []*string
}

type UninstallLogstashPluginResponseBody struct {
	Headers map[string]interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UninstallLogstashPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallLogstashPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallLogstashPluginResponseBody) GetHeaders() map[string]interface{} {
	return s.Headers
}

func (s *UninstallLogstashPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallLogstashPluginResponseBody) GetResult() []*string {
	return s.Result
}

func (s *UninstallLogstashPluginResponseBody) SetHeaders(v map[string]interface{}) *UninstallLogstashPluginResponseBody {
	s.Headers = v
	return s
}

func (s *UninstallLogstashPluginResponseBody) SetRequestId(v string) *UninstallLogstashPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallLogstashPluginResponseBody) SetResult(v []*string) *UninstallLogstashPluginResponseBody {
	s.Result = v
	return s
}

func (s *UninstallLogstashPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
