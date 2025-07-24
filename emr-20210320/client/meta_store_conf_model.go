// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaStoreConf interface {
	dara.Model
	String() string
	GoString() string
	SetDbPassword(v string) *MetaStoreConf
	GetDbPassword() *string
	SetDbUrl(v string) *MetaStoreConf
	GetDbUrl() *string
	SetDbUserName(v string) *MetaStoreConf
	GetDbUserName() *string
}

type MetaStoreConf struct {
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	DbUrl      *string `json:"DbUrl,omitempty" xml:"DbUrl,omitempty"`
	DbUserName *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
}

func (s MetaStoreConf) String() string {
	return dara.Prettify(s)
}

func (s MetaStoreConf) GoString() string {
	return s.String()
}

func (s *MetaStoreConf) GetDbPassword() *string {
	return s.DbPassword
}

func (s *MetaStoreConf) GetDbUrl() *string {
	return s.DbUrl
}

func (s *MetaStoreConf) GetDbUserName() *string {
	return s.DbUserName
}

func (s *MetaStoreConf) SetDbPassword(v string) *MetaStoreConf {
	s.DbPassword = &v
	return s
}

func (s *MetaStoreConf) SetDbUrl(v string) *MetaStoreConf {
	s.DbUrl = &v
	return s
}

func (s *MetaStoreConf) SetDbUserName(v string) *MetaStoreConf {
	s.DbUserName = &v
	return s
}

func (s *MetaStoreConf) Validate() error {
	return dara.Validate(s)
}
