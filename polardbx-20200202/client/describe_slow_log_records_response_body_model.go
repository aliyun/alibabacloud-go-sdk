// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSlowLogRecordsResponseBody
	GetDBInstanceId() *string
	SetItems(v []*DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody
	GetItems() []*DescribeSlowLogRecordsResponseBodyItems
	SetPageNumber(v string) *DescribeSlowLogRecordsResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeSlowLogRecordsResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeSlowLogRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeSlowLogRecordsResponseBody
	GetTotalCount() *string
}

type DescribeSlowLogRecordsResponseBody struct {
	// example:
	//
	// pxc-********
	DBInstanceId *string                                    `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Items        []*DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSlowLogRecordsResponseBody) GetItems() []*DescribeSlowLogRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeSlowLogRecordsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeSlowLogRecordsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeSlowLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogRecordsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeSlowLogRecordsResponseBody) SetDBInstanceId(v string) *DescribeSlowLogRecordsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v []*DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageRecordCount(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalCount(v string) *DescribeSlowLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	// example:
	//
	// pxc-i-xxxx
	CNname *string `json:"CNname,omitempty" xml:"CNname,omitempty"`
	// example:
	//
	// dcdev
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// tddl:5.4.19-20240927
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 0
	Fail *string `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// example:
	//
	// 1
	Frows *string `json:"Frows,omitempty" xml:"Frows,omitempty"`
	// example:
	//
	// ****[****] @ [1XX.XX.XX.XX]
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// example:
	//
	// pxc-xdb-s-xxxx
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// example:
	//
	// 0
	IsBind *string `json:"IsBind,omitempty" xml:"IsBind,omitempty"`
	// example:
	//
	// 1
	LockTimeMS *string `json:"LockTimeMS,omitempty" xml:"LockTimeMS,omitempty"`
	// example:
	//
	// ["1"]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// 10
	ParseRowCounts *string `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	// example:
	//
	// 2024-11-22T02:22:22.444Z
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// example:
	//
	// 3.000
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// example:
	//
	// 3000.000
	QueryTimeMS *string `json:"QueryTimeMS,omitempty" xml:"QueryTimeMS,omitempty"`
	// example:
	//
	// 20
	ReturnRowCounts *string `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// example:
	//
	// 1
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 1
	SCNT *string `json:"SCNT,omitempty" xml:"SCNT,omitempty"`
	// example:
	//
	// c8df07e5d45cd68da8b4771c2016e20b
	SQLHash *string `json:"SQLHash,omitempty" xml:"SQLHash,omitempty"`
	// example:
	//
	// select 	- from test
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// example:
	//
	// select
	SqlType    *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 0
	TooLong *string `json:"TooLong,omitempty" xml:"TooLong,omitempty"`
	// example:
	//
	// 17a5c5c840006000
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// example:
	//
	// XA
	TransactionPolicy *string `json:"TransactionPolicy,omitempty" xml:"TransactionPolicy,omitempty"`
	// example:
	//
	// 17a5c5c840006000
	TrxId *string `json:"TrxId,omitempty" xml:"TrxId,omitempty"`
	// example:
	//
	// TP
	WT *string `json:"WT,omitempty" xml:"WT,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetCNname() *string {
	return s.CNname
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetExtension() *string {
	return s.Extension
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetFail() *string {
	return s.Fail
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetFrows() *string {
	return s.Frows
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetInsName() *string {
	return s.InsName
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetIsBind() *string {
	return s.IsBind
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetLockTimeMS() *string {
	return s.LockTimeMS
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetParams() *string {
	return s.Params
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetParseRowCounts() *string {
	return s.ParseRowCounts
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetQueryTime() *string {
	return s.QueryTime
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetQueryTimeMS() *string {
	return s.QueryTimeMS
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetReturnRowCounts() *string {
	return s.ReturnRowCounts
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetRows() *string {
	return s.Rows
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetSCNT() *string {
	return s.SCNT
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetSQLHash() *string {
	return s.SQLHash
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetTooLong() *string {
	return s.TooLong
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetTransactionPolicy() *string {
	return s.TransactionPolicy
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetTrxId() *string {
	return s.TrxId
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetWT() *string {
	return s.WT
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetCNname(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.CNname = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetExtension(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Extension = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetFail(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Fail = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetFrows(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Frows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetInsName(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.InsName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetIsBind(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.IsBind = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetLockTimeMS(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.LockTimeMS = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetParams(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Params = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetParseRowCounts(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetQueryStartTime(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetQueryTime(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.QueryTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetQueryTimeMS(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.QueryTimeMS = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetReturnRowCounts(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetRows(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Rows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSCNT(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.SCNT = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSQLHash(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.SQLHash = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSqlType(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.SqlType = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetTemplateId(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetTooLong(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.TooLong = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetTraceId(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.TraceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetTransactionPolicy(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.TransactionPolicy = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetTrxId(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.TrxId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetWT(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.WT = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
