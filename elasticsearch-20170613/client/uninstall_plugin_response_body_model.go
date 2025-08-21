// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallPluginResponseBody
	GetRequestId() *string
	SetResult(v []*string) *UninstallPluginResponseBody
	GetResult() []*string
}

type UninstallPluginResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of plug-ins to be unloaded. If the unloading fails, an exception is returned.
	Result []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UninstallPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallPluginResponseBody) GetResult() []*string {
	return s.Result
}

func (s *UninstallPluginResponseBody) SetRequestId(v string) *UninstallPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallPluginResponseBody) SetResult(v []*string) *UninstallPluginResponseBody {
	s.Result = v
	return s
}

func (s *UninstallPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
