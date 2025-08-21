// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallUserPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallUserPluginsResponseBody
	GetRequestId() *string
	SetResult(v []*string) *InstallUserPluginsResponseBody
	GetResult() []*string
}

type InstallUserPluginsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6F*****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InstallUserPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallUserPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *InstallUserPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallUserPluginsResponseBody) GetResult() []*string {
	return s.Result
}

func (s *InstallUserPluginsResponseBody) SetRequestId(v string) *InstallUserPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallUserPluginsResponseBody) SetResult(v []*string) *InstallUserPluginsResponseBody {
	s.Result = v
	return s
}

func (s *InstallUserPluginsResponseBody) Validate() error {
	return dara.Validate(s)
}
