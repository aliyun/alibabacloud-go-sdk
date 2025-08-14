// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateDataSourceRequest
	GetLang() *string
	SetDescription(v string) *UpdateDataSourceRequest
	GetDescription() *string
	SetId(v int64) *UpdateDataSourceRequest
	GetId() *int64
	SetName(v string) *UpdateDataSourceRequest
	GetName() *string
	SetOssKey(v string) *UpdateDataSourceRequest
	GetOssKey() *string
	SetRegId(v string) *UpdateDataSourceRequest
	GetRegId() *string
	SetType(v string) *UpdateDataSourceRequest
	GetType() *string
}

type UpdateDataSourceRequest struct {
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
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Data source name.
	//
	// example:
	//
	// 年龄数据源
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// OSS file key.
	//
	// example:
	//
	// saf/path/xxx
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
	// example:
	//
	// FILE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataSourceRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataSourceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataSourceRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *UpdateDataSourceRequest) GetRegId() *string {
	return s.RegId
}

func (s *UpdateDataSourceRequest) GetType() *string {
	return s.Type
}

func (s *UpdateDataSourceRequest) SetLang(v string) *UpdateDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDescription(v string) *UpdateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataSourceRequest) SetId(v int64) *UpdateDataSourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataSourceRequest) SetName(v string) *UpdateDataSourceRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataSourceRequest) SetOssKey(v string) *UpdateDataSourceRequest {
	s.OssKey = &v
	return s
}

func (s *UpdateDataSourceRequest) SetRegId(v string) *UpdateDataSourceRequest {
	s.RegId = &v
	return s
}

func (s *UpdateDataSourceRequest) SetType(v string) *UpdateDataSourceRequest {
	s.Type = &v
	return s
}

func (s *UpdateDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
