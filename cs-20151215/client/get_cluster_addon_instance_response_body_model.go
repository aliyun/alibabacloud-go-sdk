// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterAddonInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetClusterAddonInstanceResponseBody
	GetConfig() *string
	SetLogging(v *GetClusterAddonInstanceResponseBodyLogging) *GetClusterAddonInstanceResponseBody
	GetLogging() *GetClusterAddonInstanceResponseBodyLogging
	SetName(v string) *GetClusterAddonInstanceResponseBody
	GetName() *string
	SetState(v string) *GetClusterAddonInstanceResponseBody
	GetState() *string
	SetVersion(v string) *GetClusterAddonInstanceResponseBody
	GetVersion() *string
}

type GetClusterAddonInstanceResponseBody struct {
	// The custom configurations of the component.
	//
	// example:
	//
	// {"sls_project_name":""}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The status of Simple Log Service.
	Logging *GetClusterAddonInstanceResponseBodyLogging `json:"logging,omitempty" xml:"logging,omitempty" type:"Struct"`
	// The name of the component instance.
	//
	// example:
	//
	// ack-node-problem-detector
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the component. Valid values:
	//
	// 	- active: The component is installed.
	//
	// 	- updating: The component is being modified.
	//
	// 	- upgrading: The component is being updated.
	//
	// 	- deleting: The component is being uninstalled.
	//
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The version of the component instance.
	//
	// example:
	//
	// 1.2.16
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetClusterAddonInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAddonInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterAddonInstanceResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetClusterAddonInstanceResponseBody) GetLogging() *GetClusterAddonInstanceResponseBodyLogging {
	return s.Logging
}

func (s *GetClusterAddonInstanceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetClusterAddonInstanceResponseBody) GetState() *string {
	return s.State
}

func (s *GetClusterAddonInstanceResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetClusterAddonInstanceResponseBody) SetConfig(v string) *GetClusterAddonInstanceResponseBody {
	s.Config = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) SetLogging(v *GetClusterAddonInstanceResponseBodyLogging) *GetClusterAddonInstanceResponseBody {
	s.Logging = v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) SetName(v string) *GetClusterAddonInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) SetState(v string) *GetClusterAddonInstanceResponseBody {
	s.State = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) SetVersion(v string) *GetClusterAddonInstanceResponseBody {
	s.Version = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClusterAddonInstanceResponseBodyLogging struct {
	// Indicates whether Simple Log Service is supported by the component.
	//
	// example:
	//
	// false
	Capable *bool `json:"capable,omitempty" xml:"capable,omitempty"`
	// Indicates whether Simple Log Service is enabled for the component.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The Simple Log Service project that is used to collect logs for the component.
	//
	// example:
	//
	// my-log-project
	LogProject *string `json:"log_project,omitempty" xml:"log_project,omitempty"`
	// The Simple Log Service Logstore that is used to collect logs for the component.
	//
	// example:
	//
	// my-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
}

func (s GetClusterAddonInstanceResponseBodyLogging) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAddonInstanceResponseBodyLogging) GoString() string {
	return s.String()
}

func (s *GetClusterAddonInstanceResponseBodyLogging) GetCapable() *bool {
	return s.Capable
}

func (s *GetClusterAddonInstanceResponseBodyLogging) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetClusterAddonInstanceResponseBodyLogging) GetLogProject() *string {
	return s.LogProject
}

func (s *GetClusterAddonInstanceResponseBodyLogging) GetLogstore() *string {
	return s.Logstore
}

func (s *GetClusterAddonInstanceResponseBodyLogging) SetCapable(v bool) *GetClusterAddonInstanceResponseBodyLogging {
	s.Capable = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBodyLogging) SetEnabled(v bool) *GetClusterAddonInstanceResponseBodyLogging {
	s.Enabled = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBodyLogging) SetLogProject(v string) *GetClusterAddonInstanceResponseBodyLogging {
	s.LogProject = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBodyLogging) SetLogstore(v string) *GetClusterAddonInstanceResponseBodyLogging {
	s.Logstore = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBodyLogging) Validate() error {
	return dara.Validate(s)
}
