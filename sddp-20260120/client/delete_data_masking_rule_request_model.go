// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataMaskingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngineType(v string) *DeleteDataMaskingRuleRequest
	GetEngineType() *string
	SetInstanceId(v string) *DeleteDataMaskingRuleRequest
	GetInstanceId() *string
	SetLang(v string) *DeleteDataMaskingRuleRequest
	GetLang() *string
	SetProductCode(v string) *DeleteDataMaskingRuleRequest
	GetProductCode() *string
	SetProductId(v int64) *DeleteDataMaskingRuleRequest
	GetProductId() *int64
	SetSubRuleList(v []*DeleteDataMaskingRuleRequestSubRuleList) *DeleteDataMaskingRuleRequest
	GetSubRuleList() []*DeleteDataMaskingRuleRequestSubRuleList
}

type DeleteDataMaskingRuleRequest struct {
	// The database engine type. The value must be the same as the EngineType of the target instance, such as MySQL or PostgreSQL.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The database instance identifier. You can obtain this value from the Items[].InstanceId field in the response of the ListDataMaskingInstances operation.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response message. Valid values: zh: Chinese. en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The cloud service code. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also specified, the two parameters must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud service type identifier. Valid values: 5: RDS. 8: POLARDB. 27: PolarDBX2. If this parameter is specified together with ProductCode, the values must match.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The list of column encryption targets grouped by database and table.
	SubRuleList []*DeleteDataMaskingRuleRequestSubRuleList `json:"SubRuleList,omitempty" xml:"SubRuleList,omitempty" type:"Repeated"`
}

func (s DeleteDataMaskingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataMaskingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataMaskingRuleRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DeleteDataMaskingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDataMaskingRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDataMaskingRuleRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DeleteDataMaskingRuleRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *DeleteDataMaskingRuleRequest) GetSubRuleList() []*DeleteDataMaskingRuleRequestSubRuleList {
	return s.SubRuleList
}

func (s *DeleteDataMaskingRuleRequest) SetEngineType(v string) *DeleteDataMaskingRuleRequest {
	s.EngineType = &v
	return s
}

func (s *DeleteDataMaskingRuleRequest) SetInstanceId(v string) *DeleteDataMaskingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDataMaskingRuleRequest) SetLang(v string) *DeleteDataMaskingRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataMaskingRuleRequest) SetProductCode(v string) *DeleteDataMaskingRuleRequest {
	s.ProductCode = &v
	return s
}

func (s *DeleteDataMaskingRuleRequest) SetProductId(v int64) *DeleteDataMaskingRuleRequest {
	s.ProductId = &v
	return s
}

func (s *DeleteDataMaskingRuleRequest) SetSubRuleList(v []*DeleteDataMaskingRuleRequestSubRuleList) *DeleteDataMaskingRuleRequest {
	s.SubRuleList = v
	return s
}

func (s *DeleteDataMaskingRuleRequest) Validate() error {
	if s.SubRuleList != nil {
		for _, item := range s.SubRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteDataMaskingRuleRequestSubRuleList struct {
	// The list of target column names. Separate multiple column names with commas (,).
	//
	// example:
	//
	// phone,email
	Columns *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// The name of the target database.
	//
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The name of the target table.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DeleteDataMaskingRuleRequestSubRuleList) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataMaskingRuleRequestSubRuleList) GoString() string {
	return s.String()
}

func (s *DeleteDataMaskingRuleRequestSubRuleList) GetColumns() *string {
	return s.Columns
}

func (s *DeleteDataMaskingRuleRequestSubRuleList) GetDbName() *string {
	return s.DbName
}

func (s *DeleteDataMaskingRuleRequestSubRuleList) GetTableName() *string {
	return s.TableName
}

func (s *DeleteDataMaskingRuleRequestSubRuleList) SetColumns(v string) *DeleteDataMaskingRuleRequestSubRuleList {
	s.Columns = &v
	return s
}

func (s *DeleteDataMaskingRuleRequestSubRuleList) SetDbName(v string) *DeleteDataMaskingRuleRequestSubRuleList {
	s.DbName = &v
	return s
}

func (s *DeleteDataMaskingRuleRequestSubRuleList) SetTableName(v string) *DeleteDataMaskingRuleRequestSubRuleList {
	s.TableName = &v
	return s
}

func (s *DeleteDataMaskingRuleRequestSubRuleList) Validate() error {
	return dara.Validate(s)
}
