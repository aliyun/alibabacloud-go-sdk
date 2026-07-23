// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineBuildConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRoutineBuildConfigurationsResponseBody
	GetRequestId() *string
	SetRoutineBuildConfigurations(v []*ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations) *ListRoutineBuildConfigurationsResponseBody
	GetRoutineBuildConfigurations() []*ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations
}

type ListRoutineBuildConfigurationsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6abd807e-ed2a-44de-ac54-ac38a62472e6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of ER build configurations.
	RoutineBuildConfigurations []*ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations `json:"RoutineBuildConfigurations,omitempty" xml:"RoutineBuildConfigurations,omitempty" type:"Repeated"`
}

func (s ListRoutineBuildConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoutineBuildConfigurationsResponseBody) GetRoutineBuildConfigurations() []*ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations {
	return s.RoutineBuildConfigurations
}

func (s *ListRoutineBuildConfigurationsResponseBody) SetRequestId(v string) *ListRoutineBuildConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBody) SetRoutineBuildConfigurations(v []*ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations) *ListRoutineBuildConfigurationsResponseBody {
	s.RoutineBuildConfigurations = v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBody) Validate() error {
	if s.RoutineBuildConfigurations != nil {
		for _, item := range s.RoutineBuildConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations struct {
	// The latest ER build task information.
	LatestRoutineBuildTask *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask `json:"LatestRoutineBuildTask,omitempty" xml:"LatestRoutineBuildTask,omitempty" type:"Struct"`
	// The ER build configuration information.
	RoutineBuildConfiguration *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration `json:"RoutineBuildConfiguration,omitempty" xml:"RoutineBuildConfiguration,omitempty" type:"Struct"`
}

func (s ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations) GetLatestRoutineBuildTask() *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask {
	return s.LatestRoutineBuildTask
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations) GetRoutineBuildConfiguration() *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration {
	return s.RoutineBuildConfiguration
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations) SetLatestRoutineBuildTask(v *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations {
	s.LatestRoutineBuildTask = v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations) SetRoutineBuildConfiguration(v *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations {
	s.RoutineBuildConfiguration = v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurations) Validate() error {
	if s.LatestRoutineBuildTask != nil {
		if err := s.LatestRoutineBuildTask.Validate(); err != nil {
			return err
		}
	}
	if s.RoutineBuildConfiguration != nil {
		if err := s.RoutineBuildConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask struct {
	// The creation time, in ISO 8601 format (UTC), formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-03-10T02:18:55Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ER routine name.
	//
	// example:
	//
	// rwa-test
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
	// The status of the build task. Valid values:
	//
	// - int: Init.
	//
	// - pending: Pending.
	//
	// - building: Building.
	//
	// - succeed: Succeeded.
	//
	// - failed: Failed.
	//
	// - canceled: Canceled.
	//
	// example:
	//
	// building
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) GetRoutineName() *string {
	return s.RoutineName
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) GetStatus() *string {
	return s.Status
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) SetCreateTime(v string) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask {
	s.CreateTime = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) SetRoutineName(v string) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask {
	s.RoutineName = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) SetStatus(v string) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask {
	s.Status = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsLatestRoutineBuildTask) Validate() error {
	return dara.Validate(s)
}

type ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration struct {
	// The Git account name.
	//
	// example:
	//
	// test
	GitAccountName *string `json:"GitAccountName,omitempty" xml:"GitAccountName,omitempty"`
	// The Git platform. Valid values: github, gitee, and upload.
	//
	// example:
	//
	// github
	GitPlatform *string `json:"GitPlatform,omitempty" xml:"GitPlatform,omitempty"`
	// The production branch name.
	//
	// example:
	//
	// main
	ProductionBranch *string `json:"ProductionBranch,omitempty" xml:"ProductionBranch,omitempty"`
	// The repository name.
	//
	// example:
	//
	// example-test
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The ER routine name.
	//
	// example:
	//
	// rwa-test
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
}

func (s ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) GetGitAccountName() *string {
	return s.GitAccountName
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) GetGitPlatform() *string {
	return s.GitPlatform
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) GetProductionBranch() *string {
	return s.ProductionBranch
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) GetRepository() *string {
	return s.Repository
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) GetRoutineName() *string {
	return s.RoutineName
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) SetGitAccountName(v string) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration {
	s.GitAccountName = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) SetGitPlatform(v string) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration {
	s.GitPlatform = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) SetProductionBranch(v string) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration {
	s.ProductionBranch = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) SetRepository(v string) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration {
	s.Repository = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) SetRoutineName(v string) *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration {
	s.RoutineName = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponseBodyRoutineBuildConfigurationsRoutineBuildConfiguration) Validate() error {
	return dara.Validate(s)
}
