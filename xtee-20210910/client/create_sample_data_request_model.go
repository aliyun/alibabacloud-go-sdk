// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateSampleDataRequest
	GetLang() *string
	SetDescription(v string) *CreateSampleDataRequest
	GetDescription() *string
	SetEncryptType(v string) *CreateSampleDataRequest
	GetEncryptType() *string
	SetName(v string) *CreateSampleDataRequest
	GetName() *string
	SetRegId(v string) *CreateSampleDataRequest
	GetRegId() *string
	SetRiskValue(v string) *CreateSampleDataRequest
	GetRiskValue() *string
	SetScene(v string) *CreateSampleDataRequest
	GetScene() *string
	SetStorePath(v string) *CreateSampleDataRequest
	GetStorePath() *string
	SetStoreType(v string) *CreateSampleDataRequest
	GetStoreType() *string
}

type CreateSampleDataRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Encryption type
	//
	// example:
	//
	// LABEL
	EncryptType *string `json:"encryptType,omitempty" xml:"encryptType,omitempty"`
	// Name
	//
	// example:
	//
	// 注册样本
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Specified risk value
	//
	// example:
	//
	// black
	RiskValue *string `json:"riskValue,omitempty" xml:"riskValue,omitempty"`
	// Scene
	//
	// example:
	//
	// PHONE
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// Storage path
	//
	// example:
	//
	// saf/de/sample/3dc2spspHKq4G3YI9d08
	StorePath *string `json:"storePath,omitempty" xml:"storePath,omitempty"`
	// Storage type
	//
	// example:
	//
	// OSS
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s CreateSampleDataRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleDataRequest) GoString() string {
	return s.String()
}

func (s *CreateSampleDataRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSampleDataRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSampleDataRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *CreateSampleDataRequest) GetName() *string {
	return s.Name
}

func (s *CreateSampleDataRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateSampleDataRequest) GetRiskValue() *string {
	return s.RiskValue
}

func (s *CreateSampleDataRequest) GetScene() *string {
	return s.Scene
}

func (s *CreateSampleDataRequest) GetStorePath() *string {
	return s.StorePath
}

func (s *CreateSampleDataRequest) GetStoreType() *string {
	return s.StoreType
}

func (s *CreateSampleDataRequest) SetLang(v string) *CreateSampleDataRequest {
	s.Lang = &v
	return s
}

func (s *CreateSampleDataRequest) SetDescription(v string) *CreateSampleDataRequest {
	s.Description = &v
	return s
}

func (s *CreateSampleDataRequest) SetEncryptType(v string) *CreateSampleDataRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateSampleDataRequest) SetName(v string) *CreateSampleDataRequest {
	s.Name = &v
	return s
}

func (s *CreateSampleDataRequest) SetRegId(v string) *CreateSampleDataRequest {
	s.RegId = &v
	return s
}

func (s *CreateSampleDataRequest) SetRiskValue(v string) *CreateSampleDataRequest {
	s.RiskValue = &v
	return s
}

func (s *CreateSampleDataRequest) SetScene(v string) *CreateSampleDataRequest {
	s.Scene = &v
	return s
}

func (s *CreateSampleDataRequest) SetStorePath(v string) *CreateSampleDataRequest {
	s.StorePath = &v
	return s
}

func (s *CreateSampleDataRequest) SetStoreType(v string) *CreateSampleDataRequest {
	s.StoreType = &v
	return s
}

func (s *CreateSampleDataRequest) Validate() error {
	return dara.Validate(s)
}
