// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbList(v []*ListDataSourcesResponseBodyDbList) *ListDataSourcesResponseBody
	GetDbList() []*ListDataSourcesResponseBodyDbList
	SetRequestId(v string) *ListDataSourcesResponseBody
	GetRequestId() *string
}

type ListDataSourcesResponseBody struct {
	DbList []*ListDataSourcesResponseBodyDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) GetDbList() []*ListDataSourcesResponseBodyDbList {
	return s.DbList
}

func (s *ListDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourcesResponseBody) SetDbList(v []*ListDataSourcesResponseBodyDbList) *ListDataSourcesResponseBody {
	s.DbList = v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) Validate() error {
	if s.DbList != nil {
		for _, item := range s.DbList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourcesResponseBodyDbList struct {
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// 0
	AuditSwitch *int32 `json:"AuditSwitch,omitempty" xml:"AuditSwitch,omitempty"`
	// example:
	//
	// 2019-06-06 09:00:00
	CreateTime  *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbAddresses []*string `json:"DbAddresses,omitempty" xml:"DbAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// -----BEGIN CERTIFICATE----- ...... -----END CERTIFICATE-----
	DbCertificate *string `json:"DbCertificate,omitempty" xml:"DbCertificate,omitempty"`
	// example:
	//
	// 1
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// rds-a235dsdg2a****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DbName       *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 3
	DbType *int32 `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// Mysql
	DbTypeName *string `json:"DbTypeName,omitempty" xml:"DbTypeName,omitempty"`
	// example:
	//
	// root
	DbUsername *string `json:"DbUsername,omitempty" xml:"DbUsername,omitempty"`
	// example:
	//
	// 5700
	DbVersion *int32 `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// example:
	//
	// 5.7
	DbVersionName *string `json:"DbVersionName,omitempty" xml:"DbVersionName,omitempty"`
}

func (s ListDataSourcesResponseBodyDbList) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyDbList) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyDbList) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListDataSourcesResponseBodyDbList) GetAuditSwitch() *int32 {
	return s.AuditSwitch
}

func (s *ListDataSourcesResponseBodyDbList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDataSourcesResponseBodyDbList) GetDbAddresses() []*string {
	return s.DbAddresses
}

func (s *ListDataSourcesResponseBodyDbList) GetDbCertificate() *string {
	return s.DbCertificate
}

func (s *ListDataSourcesResponseBodyDbList) GetDbId() *int32 {
	return s.DbId
}

func (s *ListDataSourcesResponseBodyDbList) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ListDataSourcesResponseBodyDbList) GetDbName() *string {
	return s.DbName
}

func (s *ListDataSourcesResponseBodyDbList) GetDbType() *int32 {
	return s.DbType
}

func (s *ListDataSourcesResponseBodyDbList) GetDbTypeName() *string {
	return s.DbTypeName
}

func (s *ListDataSourcesResponseBodyDbList) GetDbUsername() *string {
	return s.DbUsername
}

func (s *ListDataSourcesResponseBodyDbList) GetDbVersion() *int32 {
	return s.DbVersion
}

func (s *ListDataSourcesResponseBodyDbList) GetDbVersionName() *string {
	return s.DbVersionName
}

func (s *ListDataSourcesResponseBodyDbList) SetAssetType(v int32) *ListDataSourcesResponseBodyDbList {
	s.AssetType = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetAuditSwitch(v int32) *ListDataSourcesResponseBodyDbList {
	s.AuditSwitch = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetCreateTime(v string) *ListDataSourcesResponseBodyDbList {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbAddresses(v []*string) *ListDataSourcesResponseBodyDbList {
	s.DbAddresses = v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbCertificate(v string) *ListDataSourcesResponseBodyDbList {
	s.DbCertificate = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbId(v int32) *ListDataSourcesResponseBodyDbList {
	s.DbId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbInstanceId(v string) *ListDataSourcesResponseBodyDbList {
	s.DbInstanceId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbName(v string) *ListDataSourcesResponseBodyDbList {
	s.DbName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbType(v int32) *ListDataSourcesResponseBodyDbList {
	s.DbType = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbTypeName(v string) *ListDataSourcesResponseBodyDbList {
	s.DbTypeName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbUsername(v string) *ListDataSourcesResponseBodyDbList {
	s.DbUsername = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbVersion(v int32) *ListDataSourcesResponseBodyDbList {
	s.DbVersion = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbVersionName(v string) *ListDataSourcesResponseBodyDbList {
	s.DbVersionName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) Validate() error {
	return dara.Validate(s)
}
