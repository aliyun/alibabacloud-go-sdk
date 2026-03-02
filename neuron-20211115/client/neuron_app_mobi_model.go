// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppMobi interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *NeuronAppMobi
	GetAppId() *int64
	SetArtifactUrl(v string) *NeuronAppMobi
	GetArtifactUrl() *string
	SetCommitId(v string) *NeuronAppMobi
	GetCommitId() *string
	SetEnterpriseId(v int64) *NeuronAppMobi
	GetEnterpriseId() *int64
	SetGmtCreate(v string) *NeuronAppMobi
	GetGmtCreate() *string
	SetGmtModified(v string) *NeuronAppMobi
	GetGmtModified() *string
	SetId(v int64) *NeuronAppMobi
	GetId() *int64
	SetMobiUserId(v string) *NeuronAppMobi
	GetMobiUserId() *string
	SetModuleId(v string) *NeuronAppMobi
	GetModuleId() *string
	SetModuleName(v string) *NeuronAppMobi
	GetModuleName() *string
	SetSchemaVersion(v string) *NeuronAppMobi
	GetSchemaVersion() *string
	SetStatus(v string) *NeuronAppMobi
	GetStatus() *string
	SetStoreUrl(v string) *NeuronAppMobi
	GetStoreUrl() *string
	SetToken(v string) *NeuronAppMobi
	GetToken() *string
	SetVersion(v string) *NeuronAppMobi
	GetVersion() *string
}

type NeuronAppMobi struct {
	// example:
	//
	// 1
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sda
	ArtifactUrl *string `json:"artifactUrl,omitempty" xml:"artifactUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sda
	CommitId *string `json:"commitId,omitempty" xml:"commitId,omitempty"`
	// example:
	//
	// 1
	EnterpriseId *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	GmtCreate    *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified  *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1343424
	MobiUserId *string `json:"mobiUserId,omitempty" xml:"mobiUserId,omitempty"`
	// example:
	//
	// sda
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// sda
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// sda
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// example:
	//
	// NEW
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// sda
	StoreUrl *string `json:"storeUrl,omitempty" xml:"storeUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sdasdaddsad
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// 1.0.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s NeuronAppMobi) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppMobi) GoString() string {
	return s.String()
}

func (s *NeuronAppMobi) GetAppId() *int64 {
	return s.AppId
}

func (s *NeuronAppMobi) GetArtifactUrl() *string {
	return s.ArtifactUrl
}

func (s *NeuronAppMobi) GetCommitId() *string {
	return s.CommitId
}

func (s *NeuronAppMobi) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *NeuronAppMobi) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *NeuronAppMobi) GetGmtModified() *string {
	return s.GmtModified
}

func (s *NeuronAppMobi) GetId() *int64 {
	return s.Id
}

func (s *NeuronAppMobi) GetMobiUserId() *string {
	return s.MobiUserId
}

func (s *NeuronAppMobi) GetModuleId() *string {
	return s.ModuleId
}

func (s *NeuronAppMobi) GetModuleName() *string {
	return s.ModuleName
}

func (s *NeuronAppMobi) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *NeuronAppMobi) GetStatus() *string {
	return s.Status
}

func (s *NeuronAppMobi) GetStoreUrl() *string {
	return s.StoreUrl
}

func (s *NeuronAppMobi) GetToken() *string {
	return s.Token
}

func (s *NeuronAppMobi) GetVersion() *string {
	return s.Version
}

func (s *NeuronAppMobi) SetAppId(v int64) *NeuronAppMobi {
	s.AppId = &v
	return s
}

func (s *NeuronAppMobi) SetArtifactUrl(v string) *NeuronAppMobi {
	s.ArtifactUrl = &v
	return s
}

func (s *NeuronAppMobi) SetCommitId(v string) *NeuronAppMobi {
	s.CommitId = &v
	return s
}

func (s *NeuronAppMobi) SetEnterpriseId(v int64) *NeuronAppMobi {
	s.EnterpriseId = &v
	return s
}

func (s *NeuronAppMobi) SetGmtCreate(v string) *NeuronAppMobi {
	s.GmtCreate = &v
	return s
}

func (s *NeuronAppMobi) SetGmtModified(v string) *NeuronAppMobi {
	s.GmtModified = &v
	return s
}

func (s *NeuronAppMobi) SetId(v int64) *NeuronAppMobi {
	s.Id = &v
	return s
}

func (s *NeuronAppMobi) SetMobiUserId(v string) *NeuronAppMobi {
	s.MobiUserId = &v
	return s
}

func (s *NeuronAppMobi) SetModuleId(v string) *NeuronAppMobi {
	s.ModuleId = &v
	return s
}

func (s *NeuronAppMobi) SetModuleName(v string) *NeuronAppMobi {
	s.ModuleName = &v
	return s
}

func (s *NeuronAppMobi) SetSchemaVersion(v string) *NeuronAppMobi {
	s.SchemaVersion = &v
	return s
}

func (s *NeuronAppMobi) SetStatus(v string) *NeuronAppMobi {
	s.Status = &v
	return s
}

func (s *NeuronAppMobi) SetStoreUrl(v string) *NeuronAppMobi {
	s.StoreUrl = &v
	return s
}

func (s *NeuronAppMobi) SetToken(v string) *NeuronAppMobi {
	s.Token = &v
	return s
}

func (s *NeuronAppMobi) SetVersion(v string) *NeuronAppMobi {
	s.Version = &v
	return s
}

func (s *NeuronAppMobi) Validate() error {
	return dara.Validate(s)
}
