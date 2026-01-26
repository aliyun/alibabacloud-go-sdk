// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvCustomJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvCustomJobsResponseBody
	GetCode() *int32
	SetData(v []*ListEnvCustomJobsResponseBodyData) *ListEnvCustomJobsResponseBody
	GetData() []*ListEnvCustomJobsResponseBodyData
	SetMessage(v string) *ListEnvCustomJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvCustomJobsResponseBody
	GetRequestId() *string
}

type ListEnvCustomJobsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data []*ListEnvCustomJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A0CEDF1-06FE-44AC-8E21-21A5BE65****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnvCustomJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvCustomJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvCustomJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvCustomJobsResponseBody) GetData() []*ListEnvCustomJobsResponseBodyData {
	return s.Data
}

func (s *ListEnvCustomJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvCustomJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvCustomJobsResponseBody) SetCode(v int32) *ListEnvCustomJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvCustomJobsResponseBody) SetData(v []*ListEnvCustomJobsResponseBodyData) *ListEnvCustomJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvCustomJobsResponseBody) SetMessage(v string) *ListEnvCustomJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvCustomJobsResponseBody) SetRequestId(v string) *ListEnvCustomJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvCustomJobsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnvCustomJobsResponseBodyData struct {
	// The name of the add-on to which the custom job belongs.
	//
	// example:
	//
	// mysql
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The instance name of the add-on.
	//
	// example:
	//
	// mysql1
	AddonReleaseName *string `json:"AddonReleaseName,omitempty" xml:"AddonReleaseName,omitempty"`
	// The version of the add-on.
	//
	// example:
	//
	// 1.1.0
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// If the request parameter EncryptYaml is set to true, a Base64-encoded YAML string is returned. Otherwise, a plaintext YAML string is returned.
	//
	// example:
	//
	// Refer to supplementary instructions.
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// The time when the custom job was created. The value of this parameter is a timestamp.
	//
	// example:
	//
	// 2022-01-01T10:11:34Z
	CreationTimestamp *string `json:"CreationTimestamp,omitempty" xml:"CreationTimestamp,omitempty"`
	// The name of the custom job.
	//
	// example:
	//
	// job1
	CustomJobName *string `json:"CustomJobName,omitempty" xml:"CustomJobName,omitempty"`
	// The ID of the environment instance.
	//
	// example:
	//
	// env-xxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The capture configurations.
	ScrapeConfigs []*ListEnvCustomJobsResponseBodyDataScrapeConfigs `json:"ScrapeConfigs,omitempty" xml:"ScrapeConfigs,omitempty" type:"Repeated"`
	// The status of the custom job.
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEnvCustomJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvCustomJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvCustomJobsResponseBodyData) GetAddonName() *string {
	return s.AddonName
}

func (s *ListEnvCustomJobsResponseBodyData) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListEnvCustomJobsResponseBodyData) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *ListEnvCustomJobsResponseBodyData) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *ListEnvCustomJobsResponseBodyData) GetCreationTimestamp() *string {
	return s.CreationTimestamp
}

func (s *ListEnvCustomJobsResponseBodyData) GetCustomJobName() *string {
	return s.CustomJobName
}

func (s *ListEnvCustomJobsResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvCustomJobsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvCustomJobsResponseBodyData) GetScrapeConfigs() []*ListEnvCustomJobsResponseBodyDataScrapeConfigs {
	return s.ScrapeConfigs
}

func (s *ListEnvCustomJobsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListEnvCustomJobsResponseBodyData) SetAddonName(v string) *ListEnvCustomJobsResponseBodyData {
	s.AddonName = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) SetAddonReleaseName(v string) *ListEnvCustomJobsResponseBodyData {
	s.AddonReleaseName = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) SetAddonVersion(v string) *ListEnvCustomJobsResponseBodyData {
	s.AddonVersion = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) SetConfigYaml(v string) *ListEnvCustomJobsResponseBodyData {
	s.ConfigYaml = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) SetCreationTimestamp(v string) *ListEnvCustomJobsResponseBodyData {
	s.CreationTimestamp = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) SetCustomJobName(v string) *ListEnvCustomJobsResponseBodyData {
	s.CustomJobName = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) SetEnvironmentId(v string) *ListEnvCustomJobsResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) SetRegionId(v string) *ListEnvCustomJobsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) SetScrapeConfigs(v []*ListEnvCustomJobsResponseBodyDataScrapeConfigs) *ListEnvCustomJobsResponseBodyData {
	s.ScrapeConfigs = v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) SetStatus(v string) *ListEnvCustomJobsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyData) Validate() error {
	if s.ScrapeConfigs != nil {
		for _, item := range s.ScrapeConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnvCustomJobsResponseBodyDataScrapeConfigs struct {
	// The name of the job.
	//
	// example:
	//
	// custom-sd-demo
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The path of the metric.
	//
	// example:
	//
	// /metrics
	MetricsPath *string `json:"MetricsPath,omitempty" xml:"MetricsPath,omitempty"`
	// The service discovery methods.
	ScrapeDiscoverys []*string `json:"ScrapeDiscoverys,omitempty" xml:"ScrapeDiscoverys,omitempty" type:"Repeated"`
	// The capture interval.
	//
	// example:
	//
	// 30s
	ScrapeInterval *string `json:"ScrapeInterval,omitempty" xml:"ScrapeInterval,omitempty"`
}

func (s ListEnvCustomJobsResponseBodyDataScrapeConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListEnvCustomJobsResponseBodyDataScrapeConfigs) GoString() string {
	return s.String()
}

func (s *ListEnvCustomJobsResponseBodyDataScrapeConfigs) GetJobName() *string {
	return s.JobName
}

func (s *ListEnvCustomJobsResponseBodyDataScrapeConfigs) GetMetricsPath() *string {
	return s.MetricsPath
}

func (s *ListEnvCustomJobsResponseBodyDataScrapeConfigs) GetScrapeDiscoverys() []*string {
	return s.ScrapeDiscoverys
}

func (s *ListEnvCustomJobsResponseBodyDataScrapeConfigs) GetScrapeInterval() *string {
	return s.ScrapeInterval
}

func (s *ListEnvCustomJobsResponseBodyDataScrapeConfigs) SetJobName(v string) *ListEnvCustomJobsResponseBodyDataScrapeConfigs {
	s.JobName = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyDataScrapeConfigs) SetMetricsPath(v string) *ListEnvCustomJobsResponseBodyDataScrapeConfigs {
	s.MetricsPath = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyDataScrapeConfigs) SetScrapeDiscoverys(v []*string) *ListEnvCustomJobsResponseBodyDataScrapeConfigs {
	s.ScrapeDiscoverys = v
	return s
}

func (s *ListEnvCustomJobsResponseBodyDataScrapeConfigs) SetScrapeInterval(v string) *ListEnvCustomJobsResponseBodyDataScrapeConfigs {
	s.ScrapeInterval = &v
	return s
}

func (s *ListEnvCustomJobsResponseBodyDataScrapeConfigs) Validate() error {
	return dara.Validate(s)
}
