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
	// The request ID.
	//
	// example:
	//
	// 99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
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
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLogstashPluginsResponseBodyResult struct {
	// The plugin description.
	//
	// example:
	//
	// The clone filter is for duplicating events.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The plugin name.
	//
	// example:
	//
	// logstash-filter-clone
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The plugin source.
	//
	// example:
	//
	// SYSTEM
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The URL of the plugin documentation.
	//
	// example:
	//
	// https://xxx.html
	SpecificationUrl *string `json:"specificationUrl,omitempty" xml:"specificationUrl,omitempty"`
	// The plugin status. Valid values:
	//
	// - INSTALLED: installed
	//
	// - UNINSTALLED: not installed
	//
	// - INSTALLING: being installed
	//
	// - UNINSTALLING: being uninstalled
	//
	// - UPGRADING: being upgraded
	//
	// - FAILED: installation failed
	//
	// - UNKNOWN: the cluster is disconnected and the creation status cannot be retrieved.
	//
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
