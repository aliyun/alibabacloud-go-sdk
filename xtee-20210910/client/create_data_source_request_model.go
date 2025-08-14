// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateDataSourceRequest
	GetLang() *string
	SetDescription(v string) *CreateDataSourceRequest
	GetDescription() *string
	SetName(v string) *CreateDataSourceRequest
	GetName() *string
	SetOssKey(v string) *CreateDataSourceRequest
	GetOssKey() *string
	SetRegId(v string) *CreateDataSourceRequest
	GetRegId() *string
	SetType(v string) *CreateDataSourceRequest
	GetType() *string
}

type CreateDataSourceRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Data source description.
	//
	// example:
	//
	// 数据源描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Data source name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDispatch
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// OSS file key.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss上传key
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Data source type
	//
	// This parameter is required.
	//
	// example:
	//
	// FILE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDataSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataSourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataSourceRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateDataSourceRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateDataSourceRequest) GetType() *string {
	return s.Type
}

func (s *CreateDataSourceRequest) SetLang(v string) *CreateDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataSourceRequest) SetDescription(v string) *CreateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequest) SetName(v string) *CreateDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequest) SetOssKey(v string) *CreateDataSourceRequest {
	s.OssKey = &v
	return s
}

func (s *CreateDataSourceRequest) SetRegId(v string) *CreateDataSourceRequest {
	s.RegId = &v
	return s
}

func (s *CreateDataSourceRequest) SetType(v string) *CreateDataSourceRequest {
	s.Type = &v
	return s
}

func (s *CreateDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
