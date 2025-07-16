// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetDomainDetailModel(v *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) *DescribeCdnDomainDetailResponseBody
	GetGetDomainDetailModel() *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel
	SetRequestId(v string) *DescribeCdnDomainDetailResponseBody
	GetRequestId() *string
}

type DescribeCdnDomainDetailResponseBody struct {
	// The details about the accelerated domain name.
	GetDomainDetailModel *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel `json:"GetDomainDetailModel,omitempty" xml:"GetDomainDetailModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 18CF38AA-1275-451D-A12B-4EC0BF1C5E30
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBody) GetGetDomainDetailModel() *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	return s.GetDomainDetailModel
}

func (s *DescribeCdnDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDomainDetailResponseBody) SetGetDomainDetailModel(v *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) *DescribeCdnDomainDetailResponseBody {
	s.GetDomainDetailModel = v
	return s
}

func (s *DescribeCdnDomainDetailResponseBody) SetRequestId(v string) *DescribeCdnDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModel struct {
	// The workload type of the accelerated domain name. Valid values:
	//
	// 	- **web**: images and small files
	//
	// 	- **download**: large files
	//
	// 	- **video**: on-demand video and audio streaming
	//
	// example:
	//
	// web
	CdnType *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	// The CNAME that is assigned to the accelerated domain name. You must add the CNAME record in the system of your DNS service provider to map the accelerated domain name to the CNAME.
	//
	// example:
	//
	// example.com.w.kunlunle.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The description of the domain name.
	//
	// example:
	//
	// Streaming domain
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the accelerated domain name. Valid values:
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// 	- **configuring**
	//
	// 	- **configure_failed**
	//
	// 	- **checking**
	//
	// 	- **check_failed**
	//
	// 	- **stopping**
	//
	// 	- **deleting**
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the domain name was created.
	//
	// example:
	//
	// 2015-06-25T03:30:50Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the domain name was last modified.
	//
	// example:
	//
	// 2017-06-25T03:30:50Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The CNAME for which HTTPS is enabled.
	//
	// example:
	//
	// example.com.w.kunlunle.com
	HttpsCname *string `json:"HttpsCname,omitempty" xml:"HttpsCname,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// abcd1234abcd1234
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The acceleration region.
	//
	// example:
	//
	// domestic
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Indicates whether the SSL certificate is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	// The information about the origin server.
	SourceModels *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels `json:"SourceModels,omitempty" xml:"SourceModels,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetCdnType() *string {
	return s.CdnType
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetCname() *string {
	return s.Cname
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetDescription() *string {
	return s.Description
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetHttpsCname() *string {
	return s.HttpsCname
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetScope() *string {
	return s.Scope
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetServerCertificateStatus() *string {
	return s.ServerCertificateStatus
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GetSourceModels() *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels {
	return s.SourceModels
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetCdnType(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.CdnType = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetCname(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Cname = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetDescription(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Description = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetDomainName(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetDomainStatus(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.DomainStatus = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetGmtCreated(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetGmtModified(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.GmtModified = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetHttpsCname(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.HttpsCname = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetResourceGroupId(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetScope(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Scope = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetServerCertificateStatus(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetSourceModels(v *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.SourceModels = v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels struct {
	SourceModel []*DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel `json:"SourceModel,omitempty" xml:"SourceModel,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) GetSourceModel() []*DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	return s.SourceModel
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) SetSourceModel(v []*DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels {
	s.SourceModel = v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel struct {
	// The address of the origin server.
	//
	// example:
	//
	// example.com
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The status.
	//
	// example:
	//
	// online
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The port over which requests are redirected to the origin server. Ports 443 and 80 are supported.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority.
	//
	// example:
	//
	// 20
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the origin server. Valid values:
	//
	// 	- **ipaddr**: an origin IP address
	//
	// 	- **domain**: an origin domain name
	//
	// 	- **oss**: the domain name of an Object Storage Service (OSS) bucket
	//
	// 	- **fc_domain:*	- a Function Compute domain name
	//
	// example:
	//
	// domain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin server if multiple origin servers have been specified.
	//
	// example:
	//
	// 10
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) GetContent() *string {
	return s.Content
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) GetEnabled() *string {
	return s.Enabled
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) GetPort() *int32 {
	return s.Port
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) GetPriority() *string {
	return s.Priority
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) GetType() *string {
	return s.Type
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) GetWeight() *string {
	return s.Weight
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetContent(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Content = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetEnabled(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Enabled = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetPort(v int32) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Port = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetPriority(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Priority = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetType(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Type = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetWeight(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Weight = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) Validate() error {
	return dara.Validate(s)
}
