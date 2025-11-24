// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetaData(v *DescribeMetadataResponseBodyMetaData) *DescribeMetadataResponseBody
	GetMetaData() *DescribeMetadataResponseBodyMetaData
	SetRequestId(v string) *DescribeMetadataResponseBody
	GetRequestId() *string
}

type DescribeMetadataResponseBody struct {
	// The ASM metadata, including the basic information about ASM.
	MetaData *DescribeMetadataResponseBodyMetaData `json:"MetaData,omitempty" xml:"MetaData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F93DDAD7-6E04-5AC3-86F4-852708******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetadataResponseBody) GetMetaData() *DescribeMetadataResponseBodyMetaData {
	return s.MetaData
}

func (s *DescribeMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetadataResponseBody) SetMetaData(v *DescribeMetadataResponseBodyMetaData) *DescribeMetadataResponseBody {
	s.MetaData = v
	return s
}

func (s *DescribeMetadataResponseBody) SetRequestId(v string) *DescribeMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetadataResponseBody) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetadataResponseBodyMetaData struct {
	// The Kubernetes versions compatible with ASM.
	CompatibilityInfoList []map[string]interface{} `json:"CompatibilityInfoList,omitempty" xml:"CompatibilityInfoList,omitempty" type:"Repeated"`
	// The current version.
	//
	// example:
	//
	// v1.14.3.87-g96cf7305-aliyun
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The data of the ASM Playground.
	PlaygroundScene *DescribeMetadataResponseBodyMetaDataPlaygroundScene `json:"PlaygroundScene,omitempty" xml:"PlaygroundScene,omitempty" type:"Struct"`
	// The version information about ASM of a commercial edition.
	ProEdition *DescribeMetadataResponseBodyMetaDataProEdition `json:"ProEdition,omitempty" xml:"ProEdition,omitempty" type:"Struct"`
	// The regions where ASM instances can be created.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The custom resource definitions (CRDs) of all ASM versions.
	VersionCrds []map[string]interface{} `json:"VersionCrds,omitempty" xml:"VersionCrds,omitempty" type:"Repeated"`
	// The Istio versions corresponding to the ASM versions.
	VersionRegistry []map[string]interface{} `json:"VersionRegistry,omitempty" xml:"VersionRegistry,omitempty" type:"Repeated"`
	// The list of ASM versions.
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeMetadataResponseBodyMetaData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetadataResponseBodyMetaData) GoString() string {
	return s.String()
}

func (s *DescribeMetadataResponseBodyMetaData) GetCompatibilityInfoList() []map[string]interface{} {
	return s.CompatibilityInfoList
}

func (s *DescribeMetadataResponseBodyMetaData) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeMetadataResponseBodyMetaData) GetPlaygroundScene() *DescribeMetadataResponseBodyMetaDataPlaygroundScene {
	return s.PlaygroundScene
}

func (s *DescribeMetadataResponseBodyMetaData) GetProEdition() *DescribeMetadataResponseBodyMetaDataProEdition {
	return s.ProEdition
}

func (s *DescribeMetadataResponseBodyMetaData) GetRegions() []*string {
	return s.Regions
}

func (s *DescribeMetadataResponseBodyMetaData) GetVersionCrds() []map[string]interface{} {
	return s.VersionCrds
}

func (s *DescribeMetadataResponseBodyMetaData) GetVersionRegistry() []map[string]interface{} {
	return s.VersionRegistry
}

func (s *DescribeMetadataResponseBodyMetaData) GetVersions() []*string {
	return s.Versions
}

func (s *DescribeMetadataResponseBodyMetaData) SetCompatibilityInfoList(v []map[string]interface{}) *DescribeMetadataResponseBodyMetaData {
	s.CompatibilityInfoList = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetCurrentVersion(v string) *DescribeMetadataResponseBodyMetaData {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetPlaygroundScene(v *DescribeMetadataResponseBodyMetaDataPlaygroundScene) *DescribeMetadataResponseBodyMetaData {
	s.PlaygroundScene = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetProEdition(v *DescribeMetadataResponseBodyMetaDataProEdition) *DescribeMetadataResponseBodyMetaData {
	s.ProEdition = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetRegions(v []*string) *DescribeMetadataResponseBodyMetaData {
	s.Regions = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetVersionCrds(v []map[string]interface{}) *DescribeMetadataResponseBodyMetaData {
	s.VersionCrds = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetVersionRegistry(v []map[string]interface{}) *DescribeMetadataResponseBodyMetaData {
	s.VersionRegistry = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetVersions(v []*string) *DescribeMetadataResponseBodyMetaData {
	s.Versions = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) Validate() error {
	if s.PlaygroundScene != nil {
		if err := s.PlaygroundScene.Validate(); err != nil {
			return err
		}
	}
	if s.ProEdition != nil {
		if err := s.ProEdition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetadataResponseBodyMetaDataPlaygroundScene struct {
	// The observability scenarios.
	Observability []*string `json:"observability,omitempty" xml:"observability,omitempty" type:"Repeated"`
	// Other scenarios.
	Other []*string `json:"other,omitempty" xml:"other,omitempty" type:"Repeated"`
	// The security scenarios.
	Security []*string `json:"security,omitempty" xml:"security,omitempty" type:"Repeated"`
	// The traffic management scenarios.
	TrafficManagement []*string `json:"trafficManagement,omitempty" xml:"trafficManagement,omitempty" type:"Repeated"`
}

func (s DescribeMetadataResponseBodyMetaDataPlaygroundScene) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetadataResponseBodyMetaDataPlaygroundScene) GoString() string {
	return s.String()
}

func (s *DescribeMetadataResponseBodyMetaDataPlaygroundScene) GetObservability() []*string {
	return s.Observability
}

func (s *DescribeMetadataResponseBodyMetaDataPlaygroundScene) GetOther() []*string {
	return s.Other
}

func (s *DescribeMetadataResponseBodyMetaDataPlaygroundScene) GetSecurity() []*string {
	return s.Security
}

func (s *DescribeMetadataResponseBodyMetaDataPlaygroundScene) GetTrafficManagement() []*string {
	return s.TrafficManagement
}

func (s *DescribeMetadataResponseBodyMetaDataPlaygroundScene) SetObservability(v []*string) *DescribeMetadataResponseBodyMetaDataPlaygroundScene {
	s.Observability = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataPlaygroundScene) SetOther(v []*string) *DescribeMetadataResponseBodyMetaDataPlaygroundScene {
	s.Other = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataPlaygroundScene) SetSecurity(v []*string) *DescribeMetadataResponseBodyMetaDataPlaygroundScene {
	s.Security = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataPlaygroundScene) SetTrafficManagement(v []*string) *DescribeMetadataResponseBodyMetaDataPlaygroundScene {
	s.TrafficManagement = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataPlaygroundScene) Validate() error {
	return dara.Validate(s)
}

type DescribeMetadataResponseBodyMetaDataProEdition struct {
	// The current version.
	//
	// example:
	//
	// v1.14.3.87-g96cf7305-aliyun
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The CRDs of all ASM versions.
	VersionCrds []map[string]interface{} `json:"VersionCrds,omitempty" xml:"VersionCrds,omitempty" type:"Repeated"`
	// The Istio versions corresponding to the ASM versions.
	VersionRegistry []map[string]interface{} `json:"VersionRegistry,omitempty" xml:"VersionRegistry,omitempty" type:"Repeated"`
	// The list of ASM versions.
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeMetadataResponseBodyMetaDataProEdition) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetadataResponseBodyMetaDataProEdition) GoString() string {
	return s.String()
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) GetVersionCrds() []map[string]interface{} {
	return s.VersionCrds
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) GetVersionRegistry() []map[string]interface{} {
	return s.VersionRegistry
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) GetVersions() []*string {
	return s.Versions
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) SetCurrentVersion(v string) *DescribeMetadataResponseBodyMetaDataProEdition {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) SetVersionCrds(v []map[string]interface{}) *DescribeMetadataResponseBodyMetaDataProEdition {
	s.VersionCrds = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) SetVersionRegistry(v []map[string]interface{}) *DescribeMetadataResponseBodyMetaDataProEdition {
	s.VersionRegistry = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) SetVersions(v []*string) *DescribeMetadataResponseBodyMetaDataProEdition {
	s.Versions = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) Validate() error {
	return dara.Validate(s)
}
