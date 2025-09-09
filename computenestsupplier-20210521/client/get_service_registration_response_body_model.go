// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRegistrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *GetServiceRegistrationResponseBody
	GetComment() *string
	SetDetail(v *GetServiceRegistrationResponseBodyDetail) *GetServiceRegistrationResponseBody
	GetDetail() *GetServiceRegistrationResponseBodyDetail
	SetFinishTime(v string) *GetServiceRegistrationResponseBody
	GetFinishTime() *string
	SetRegistrationId(v string) *GetServiceRegistrationResponseBody
	GetRegistrationId() *string
	SetRequestId(v string) *GetServiceRegistrationResponseBody
	GetRequestId() *string
	SetServiceId(v string) *GetServiceRegistrationResponseBody
	GetServiceId() *string
	SetServiceInfo(v *GetServiceRegistrationResponseBodyServiceInfo) *GetServiceRegistrationResponseBody
	GetServiceInfo() *GetServiceRegistrationResponseBodyServiceInfo
	SetServiceVersion(v string) *GetServiceRegistrationResponseBody
	GetServiceVersion() *string
	SetStatus(v string) *GetServiceRegistrationResponseBody
	GetStatus() *string
	SetSubmitTime(v string) *GetServiceRegistrationResponseBody
	GetSubmitTime() *string
}

type GetServiceRegistrationResponseBody struct {
	// Comment from reviewer.
	//
	// example:
	//
	// comment message
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The details of service audit.
	Detail *GetServiceRegistrationResponseBodyDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// Finish time.
	//
	// example:
	//
	// 2024-12-07T11:05:50Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// Service registration ID.
	//
	// example:
	//
	// sr-1b4aabc1c9eb4109****
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A361BA9E-xxxx-xxxx-xxxx-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-c2d118c9193e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service details.
	ServiceInfo *GetServiceRegistrationResponseBodyServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Struct"`
	// The service version.
	//
	// example:
	//
	// beta
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The status of service registration. Valid values:
	//
	// 	- Submitted
	//
	// 	- Approved
	//
	// 	- Rejected
	//
	// 	- Canceled
	//
	// 	- Executed
	//
	// example:
	//
	// Submitted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Submit time.
	//
	// example:
	//
	// 2024-12-07T11:05:50Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s GetServiceRegistrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationResponseBody) GetComment() *string {
	return s.Comment
}

func (s *GetServiceRegistrationResponseBody) GetDetail() *GetServiceRegistrationResponseBodyDetail {
	return s.Detail
}

func (s *GetServiceRegistrationResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetServiceRegistrationResponseBody) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *GetServiceRegistrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceRegistrationResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceRegistrationResponseBody) GetServiceInfo() *GetServiceRegistrationResponseBodyServiceInfo {
	return s.ServiceInfo
}

func (s *GetServiceRegistrationResponseBody) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceRegistrationResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetServiceRegistrationResponseBody) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *GetServiceRegistrationResponseBody) SetComment(v string) *GetServiceRegistrationResponseBody {
	s.Comment = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetDetail(v *GetServiceRegistrationResponseBodyDetail) *GetServiceRegistrationResponseBody {
	s.Detail = v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetFinishTime(v string) *GetServiceRegistrationResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetRegistrationId(v string) *GetServiceRegistrationResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetRequestId(v string) *GetServiceRegistrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetServiceId(v string) *GetServiceRegistrationResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetServiceInfo(v *GetServiceRegistrationResponseBodyServiceInfo) *GetServiceRegistrationResponseBody {
	s.ServiceInfo = v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetServiceVersion(v string) *GetServiceRegistrationResponseBody {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetStatus(v string) *GetServiceRegistrationResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetSubmitTime(v string) *GetServiceRegistrationResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceRegistrationResponseBodyDetail struct {
	// Whether risk exists.
	//
	// example:
	//
	// true
	AtRisk *bool `json:"AtRisk,omitempty" xml:"AtRisk,omitempty"`
	// Whether service is associated with artifact.
	//
	// example:
	//
	// true
	HasRelatedArtifact *bool `json:"HasRelatedArtifact,omitempty" xml:"HasRelatedArtifact,omitempty"`
	// The reports.
	//
	// example:
	//
	// { "template1":"https://compute-nest-security-audit-bucket.oss-cn-hangzhou.aliyuncs.com/report" }
	Reports *string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	// The url of template diff file.
	//
	// example:
	//
	// https://compute-nest-template-diff-bucket.oss-cn-hangzhou.aliyuncs.com/service-abc/diff
	TemplateDiffUrl *string `json:"TemplateDiffUrl,omitempty" xml:"TemplateDiffUrl,omitempty"`
}

func (s GetServiceRegistrationResponseBodyDetail) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRegistrationResponseBodyDetail) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationResponseBodyDetail) GetAtRisk() *bool {
	return s.AtRisk
}

func (s *GetServiceRegistrationResponseBodyDetail) GetHasRelatedArtifact() *bool {
	return s.HasRelatedArtifact
}

func (s *GetServiceRegistrationResponseBodyDetail) GetReports() *string {
	return s.Reports
}

func (s *GetServiceRegistrationResponseBodyDetail) GetTemplateDiffUrl() *string {
	return s.TemplateDiffUrl
}

func (s *GetServiceRegistrationResponseBodyDetail) SetAtRisk(v bool) *GetServiceRegistrationResponseBodyDetail {
	s.AtRisk = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyDetail) SetHasRelatedArtifact(v bool) *GetServiceRegistrationResponseBodyDetail {
	s.HasRelatedArtifact = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyDetail) SetReports(v string) *GetServiceRegistrationResponseBodyDetail {
	s.Reports = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyDetail) SetTemplateDiffUrl(v string) *GetServiceRegistrationResponseBodyDetail {
	s.TemplateDiffUrl = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyDetail) Validate() error {
	return dara.Validate(s)
}

type GetServiceRegistrationResponseBodyServiceInfo struct {
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// The version name.
	//
	// example:
	//
	// v1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceRegistrationResponseBodyServiceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRegistrationResponseBodyServiceInfo) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) GetTrialType() *string {
	return s.TrialType
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) GetVersionName() *string {
	return s.VersionName
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) SetServiceType(v string) *GetServiceRegistrationResponseBodyServiceInfo {
	s.ServiceType = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) SetTrialType(v string) *GetServiceRegistrationResponseBodyServiceInfo {
	s.TrialType = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) SetVersionName(v string) *GetServiceRegistrationResponseBodyServiceInfo {
	s.VersionName = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) Validate() error {
	return dara.Validate(s)
}
