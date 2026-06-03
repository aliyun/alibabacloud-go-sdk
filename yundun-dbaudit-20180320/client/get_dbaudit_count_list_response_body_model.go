// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBAuditCountListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbList(v []*GetDBAuditCountListResponseBodyDbList) *GetDBAuditCountListResponseBody
	GetDbList() []*GetDBAuditCountListResponseBodyDbList
	SetRequestId(v string) *GetDBAuditCountListResponseBody
	GetRequestId() *string
}

type GetDBAuditCountListResponseBody struct {
	DbList []*GetDBAuditCountListResponseBodyDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDBAuditCountListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDBAuditCountListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBAuditCountListResponseBody) GetDbList() []*GetDBAuditCountListResponseBodyDbList {
	return s.DbList
}

func (s *GetDBAuditCountListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDBAuditCountListResponseBody) SetDbList(v []*GetDBAuditCountListResponseBodyDbList) *GetDBAuditCountListResponseBody {
	s.DbList = v
	return s
}

func (s *GetDBAuditCountListResponseBody) SetRequestId(v string) *GetDBAuditCountListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBAuditCountListResponseBody) Validate() error {
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

type GetDBAuditCountListResponseBodyDbList struct {
	// example:
	//
	// 0
	AssetType   *int32    `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	DbAddresses []*string `json:"DbAddresses,omitempty" xml:"DbAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DbId   *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
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
	// 5700
	DbVersion *int32 `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// example:
	//
	// 5.7
	DbVersionName *string `json:"DbVersionName,omitempty" xml:"DbVersionName,omitempty"`
	// example:
	//
	// 1000
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// example:
	//
	// 2000
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// 100000
	SqlCount *int64 `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
}

func (s GetDBAuditCountListResponseBodyDbList) String() string {
	return dara.Prettify(s)
}

func (s GetDBAuditCountListResponseBodyDbList) GoString() string {
	return s.String()
}

func (s *GetDBAuditCountListResponseBodyDbList) GetAssetType() *int32 {
	return s.AssetType
}

func (s *GetDBAuditCountListResponseBodyDbList) GetDbAddresses() []*string {
	return s.DbAddresses
}

func (s *GetDBAuditCountListResponseBodyDbList) GetDbId() *int32 {
	return s.DbId
}

func (s *GetDBAuditCountListResponseBodyDbList) GetDbName() *string {
	return s.DbName
}

func (s *GetDBAuditCountListResponseBodyDbList) GetDbType() *int32 {
	return s.DbType
}

func (s *GetDBAuditCountListResponseBodyDbList) GetDbTypeName() *string {
	return s.DbTypeName
}

func (s *GetDBAuditCountListResponseBodyDbList) GetDbVersion() *int32 {
	return s.DbVersion
}

func (s *GetDBAuditCountListResponseBodyDbList) GetDbVersionName() *string {
	return s.DbVersionName
}

func (s *GetDBAuditCountListResponseBodyDbList) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *GetDBAuditCountListResponseBodyDbList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *GetDBAuditCountListResponseBodyDbList) GetSqlCount() *int64 {
	return s.SqlCount
}

func (s *GetDBAuditCountListResponseBodyDbList) SetAssetType(v int32) *GetDBAuditCountListResponseBodyDbList {
	s.AssetType = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbAddresses(v []*string) *GetDBAuditCountListResponseBodyDbList {
	s.DbAddresses = v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbId(v int32) *GetDBAuditCountListResponseBodyDbList {
	s.DbId = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbName(v string) *GetDBAuditCountListResponseBodyDbList {
	s.DbName = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbType(v int32) *GetDBAuditCountListResponseBodyDbList {
	s.DbType = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbTypeName(v string) *GetDBAuditCountListResponseBodyDbList {
	s.DbTypeName = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbVersion(v int32) *GetDBAuditCountListResponseBodyDbList {
	s.DbVersion = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbVersionName(v string) *GetDBAuditCountListResponseBodyDbList {
	s.DbVersionName = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetRiskCount(v int64) *GetDBAuditCountListResponseBodyDbList {
	s.RiskCount = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetSessionCount(v int64) *GetDBAuditCountListResponseBodyDbList {
	s.SessionCount = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetSqlCount(v int64) *GetDBAuditCountListResponseBodyDbList {
	s.SqlCount = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) Validate() error {
	return dara.Validate(s)
}
