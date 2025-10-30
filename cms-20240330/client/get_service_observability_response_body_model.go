// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceObservabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntryPointInfo(v *GetServiceObservabilityResponseBodyEntryPointInfo) *GetServiceObservabilityResponseBody
	GetEntryPointInfo() *GetServiceObservabilityResponseBodyEntryPointInfo
	SetFeeType(v string) *GetServiceObservabilityResponseBody
	GetFeeType() *string
	SetQuotas(v map[string]*string) *GetServiceObservabilityResponseBody
	GetQuotas() map[string]*string
	SetRegionId(v string) *GetServiceObservabilityResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetServiceObservabilityResponseBody
	GetRequestId() *string
	SetSettings(v map[string]*string) *GetServiceObservabilityResponseBody
	GetSettings() map[string]*string
	SetStatus(v string) *GetServiceObservabilityResponseBody
	GetStatus() *string
	SetType(v string) *GetServiceObservabilityResponseBody
	GetType() *string
	SetWorkspace(v string) *GetServiceObservabilityResponseBody
	GetWorkspace() *string
}

type GetServiceObservabilityResponseBody struct {
	// Endpoint and Authentication Information
	EntryPointInfo *GetServiceObservabilityResponseBodyEntryPointInfo `json:"entryPointInfo,omitempty" xml:"entryPointInfo,omitempty" type:"Struct"`
	// Billing Type
	//
	// example:
	//
	// arms=serverless;xtrace=serverless
	FeeType *string `json:"feeType,omitempty" xml:"feeType,omitempty"`
	// Quota Configuration
	Quotas map[string]*string `json:"quotas,omitempty" xml:"quotas,omitempty"`
	// Region
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Request ID
	//
	// example:
	//
	// 4852B9B5-345C-5CBC-A15F-786D83ECCBBA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// System Configuration
	Settings map[string]*string `json:"settings,omitempty" xml:"settings,omitempty"`
	// Resource Initialization Status
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Application Observability Type
	//
	// example:
	//
	// apm
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Workspace Name
	//
	// example:
	//
	// default-cms-1654218***343050-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetServiceObservabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceObservabilityResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceObservabilityResponseBody) GetEntryPointInfo() *GetServiceObservabilityResponseBodyEntryPointInfo {
	return s.EntryPointInfo
}

func (s *GetServiceObservabilityResponseBody) GetFeeType() *string {
	return s.FeeType
}

func (s *GetServiceObservabilityResponseBody) GetQuotas() map[string]*string {
	return s.Quotas
}

func (s *GetServiceObservabilityResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceObservabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceObservabilityResponseBody) GetSettings() map[string]*string {
	return s.Settings
}

func (s *GetServiceObservabilityResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetServiceObservabilityResponseBody) GetType() *string {
	return s.Type
}

func (s *GetServiceObservabilityResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetServiceObservabilityResponseBody) SetEntryPointInfo(v *GetServiceObservabilityResponseBodyEntryPointInfo) *GetServiceObservabilityResponseBody {
	s.EntryPointInfo = v
	return s
}

func (s *GetServiceObservabilityResponseBody) SetFeeType(v string) *GetServiceObservabilityResponseBody {
	s.FeeType = &v
	return s
}

func (s *GetServiceObservabilityResponseBody) SetQuotas(v map[string]*string) *GetServiceObservabilityResponseBody {
	s.Quotas = v
	return s
}

func (s *GetServiceObservabilityResponseBody) SetRegionId(v string) *GetServiceObservabilityResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetServiceObservabilityResponseBody) SetRequestId(v string) *GetServiceObservabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceObservabilityResponseBody) SetSettings(v map[string]*string) *GetServiceObservabilityResponseBody {
	s.Settings = v
	return s
}

func (s *GetServiceObservabilityResponseBody) SetStatus(v string) *GetServiceObservabilityResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceObservabilityResponseBody) SetType(v string) *GetServiceObservabilityResponseBody {
	s.Type = &v
	return s
}

func (s *GetServiceObservabilityResponseBody) SetWorkspace(v string) *GetServiceObservabilityResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetServiceObservabilityResponseBody) Validate() error {
	if s.EntryPointInfo != nil {
		if err := s.EntryPointInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceObservabilityResponseBodyEntryPointInfo struct {
	// Authentication Token for Data Reporting
	//
	// example:
	//
	// gaddp****@de20f2***1ce***
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// Private Network Access Address
	//
	// example:
	//
	// project-xtrace-xxxx-cn-hangzhou.cn-hangzhou-intranet.log.aliyuncs.com
	PrivateDomain *string `json:"privateDomain,omitempty" xml:"privateDomain,omitempty"`
	// SLS Project
	//
	// example:
	//
	// proj-xtrace-xxxxx
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// Public Network Access Address
	//
	// example:
	//
	// project-xtrace-xxxx-cn-hangzhou.cn-hangzhou.log.aliyuncs.com
	PublicDomain *string `json:"publicDomain,omitempty" xml:"publicDomain,omitempty"`
}

func (s GetServiceObservabilityResponseBodyEntryPointInfo) String() string {
	return dara.Prettify(s)
}

func (s GetServiceObservabilityResponseBodyEntryPointInfo) GoString() string {
	return s.String()
}

func (s *GetServiceObservabilityResponseBodyEntryPointInfo) GetAuthToken() *string {
	return s.AuthToken
}

func (s *GetServiceObservabilityResponseBodyEntryPointInfo) GetPrivateDomain() *string {
	return s.PrivateDomain
}

func (s *GetServiceObservabilityResponseBodyEntryPointInfo) GetProject() *string {
	return s.Project
}

func (s *GetServiceObservabilityResponseBodyEntryPointInfo) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *GetServiceObservabilityResponseBodyEntryPointInfo) SetAuthToken(v string) *GetServiceObservabilityResponseBodyEntryPointInfo {
	s.AuthToken = &v
	return s
}

func (s *GetServiceObservabilityResponseBodyEntryPointInfo) SetPrivateDomain(v string) *GetServiceObservabilityResponseBodyEntryPointInfo {
	s.PrivateDomain = &v
	return s
}

func (s *GetServiceObservabilityResponseBodyEntryPointInfo) SetProject(v string) *GetServiceObservabilityResponseBodyEntryPointInfo {
	s.Project = &v
	return s
}

func (s *GetServiceObservabilityResponseBodyEntryPointInfo) SetPublicDomain(v string) *GetServiceObservabilityResponseBodyEntryPointInfo {
	s.PublicDomain = &v
	return s
}

func (s *GetServiceObservabilityResponseBodyEntryPointInfo) Validate() error {
	return dara.Validate(s)
}
