// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstashPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListLogstashPluginsResponseBody
	GetRequestId() *string
	SetResult(v []*ListLogstashPluginsResponseBodyResult) *ListLogstashPluginsResponseBody
	GetResult() []*ListLogstashPluginsResponseBodyResult
}

type ListLogstashPluginsResponseBody struct {
	// The address of the documentation for the plug-in.
	//
	// example:
	//
	// 99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the plug-in. Valid values:
	//
	// 	- INSTALLED: Installed
	//
	// 	- UNINSTALLED: Not installed
	//
	// 	- INSTALLING: The instance is being installed.
	//
	// 	- UNINSTALLING: The instance is being uninstalled.
	//
	// 	- UPGRADING: The backup gateway is being upgraded.
	//
	// 	- FAILED: Installation failed
	//
	// 	- UNKNOWN: The cluster is lost and cannot be created.
	Result []*ListLogstashPluginsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListLogstashPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogstashPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogstashPluginsResponseBody) GetResult() []*ListLogstashPluginsResponseBodyResult {
	return s.Result
}

func (s *ListLogstashPluginsResponseBody) SetRequestId(v string) *ListLogstashPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogstashPluginsResponseBody) SetResult(v []*ListLogstashPluginsResponseBodyResult) *ListLogstashPluginsResponseBody {
	s.Result = v
	return s
}

func (s *ListLogstashPluginsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLogstashPluginsResponseBodyResult struct {
	// The source of the plug-in.
	//
	// example:
	//
	// The clone filter is for duplicating events.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// logstash-filter-clone
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// SYSTEM
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// https://xxx.html
	SpecificationUrl *string `json:"specificationUrl,omitempty" xml:"specificationUrl,omitempty"`
	// example:
	//
	// INSTALLED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListLogstashPluginsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashPluginsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLogstashPluginsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListLogstashPluginsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListLogstashPluginsResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ListLogstashPluginsResponseBodyResult) GetSpecificationUrl() *string {
	return s.SpecificationUrl
}

func (s *ListLogstashPluginsResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListLogstashPluginsResponseBodyResult) SetDescription(v string) *ListLogstashPluginsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListLogstashPluginsResponseBodyResult) SetName(v string) *ListLogstashPluginsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListLogstashPluginsResponseBodyResult) SetSource(v string) *ListLogstashPluginsResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListLogstashPluginsResponseBodyResult) SetSpecificationUrl(v string) *ListLogstashPluginsResponseBodyResult {
	s.SpecificationUrl = &v
	return s
}

func (s *ListLogstashPluginsResponseBodyResult) SetState(v string) *ListLogstashPluginsResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListLogstashPluginsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
