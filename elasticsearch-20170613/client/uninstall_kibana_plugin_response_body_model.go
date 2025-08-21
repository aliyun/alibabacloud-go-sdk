// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallKibanaPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallKibanaPluginResponseBody
	GetRequestId() *string
	SetResult(v []*string) *UninstallKibanaPluginResponseBody
	GetResult() []*string
}

type UninstallKibanaPluginResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result shows a list of uninstalled plug-ins.
	Result []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UninstallKibanaPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallKibanaPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallKibanaPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallKibanaPluginResponseBody) GetResult() []*string {
	return s.Result
}

func (s *UninstallKibanaPluginResponseBody) SetRequestId(v string) *UninstallKibanaPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallKibanaPluginResponseBody) SetResult(v []*string) *UninstallKibanaPluginResponseBody {
	s.Result = v
	return s
}

func (s *UninstallKibanaPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
