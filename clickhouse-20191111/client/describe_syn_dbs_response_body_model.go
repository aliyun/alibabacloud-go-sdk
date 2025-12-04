// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynDbsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSynDbsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSynDbsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSynDbsResponseBody
	GetRequestId() *string
	SetSynDbs(v []*DescribeSynDbsResponseBodySynDbs) *DescribeSynDbsResponseBody
	GetSynDbs() []*DescribeSynDbsResponseBodySynDbs
	SetTotalCount(v int32) *DescribeSynDbsResponseBody
	GetTotalCount() *int32
}

type DescribeSynDbsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7655F5F9-1313-5ABA-8516-F6EB79605A5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about data synchronization between the ApsaraDB for ClickHouse cluster and an ApsaraDB RDS for MySQL instance.
	SynDbs []*DescribeSynDbsResponseBodySynDbs `json:"SynDbs,omitempty" xml:"SynDbs,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSynDbsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynDbsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSynDbsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSynDbsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSynDbsResponseBody) GetSynDbs() []*DescribeSynDbsResponseBodySynDbs {
	return s.SynDbs
}

func (s *DescribeSynDbsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSynDbsResponseBody) SetPageNumber(v int32) *DescribeSynDbsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynDbsResponseBody) SetPageSize(v int32) *DescribeSynDbsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSynDbsResponseBody) SetRequestId(v string) *DescribeSynDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynDbsResponseBody) SetSynDbs(v []*DescribeSynDbsResponseBodySynDbs) *DescribeSynDbsResponseBody {
	s.SynDbs = v
	return s
}

func (s *DescribeSynDbsResponseBody) SetTotalCount(v int32) *DescribeSynDbsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSynDbsResponseBody) Validate() error {
	if s.SynDbs != nil {
		for _, item := range s.SynDbs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSynDbsResponseBodySynDbs struct {
	// 	- When the value **true*	- is returned for the **SynStatus*	- parameter, the system does not return the ErrorMsg parameter.
	//
	// 	- When the value **false*	- is returned for the **SynStatus*	- parameter, the system returns for the ErrorMsg parameter the cause why the data synchronization failed.
	//
	// example:
	//
	// ClickHouse exception, code: 49, host: 100.100.118.132, port: 49670; Code: 49, e.displayText() = DB::Exception: Logical error: there is no global context (version 20.8.17.25)n
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// rm-wz9d11qg1j0h4****
	RdsId *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	// The database account that is used to log on to the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// test
	RdsUserName *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	// The internal endpoint of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// rm-bp16t9h3999xb0a711****.mysql.rds.aliyuncs.com:3306
	RdsVpcUrl *string `json:"RdsVpcUrl,omitempty" xml:"RdsVpcUrl,omitempty"`
	// The name of the database in the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// database
	SynDb *string `json:"SynDb,omitempty" xml:"SynDb,omitempty"`
	// Indicates whether the data synchronization succeeded. Valid values:
	//
	// 	- **true**: The data synchronization succeeded.
	//
	// 	- **false**: The data synchronization failed.
	//
	// example:
	//
	// true
	SynStatus *bool `json:"SynStatus,omitempty" xml:"SynStatus,omitempty"`
}

func (s DescribeSynDbsResponseBodySynDbs) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynDbsResponseBodySynDbs) GoString() string {
	return s.String()
}

func (s *DescribeSynDbsResponseBodySynDbs) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeSynDbsResponseBodySynDbs) GetRdsId() *string {
	return s.RdsId
}

func (s *DescribeSynDbsResponseBodySynDbs) GetRdsUserName() *string {
	return s.RdsUserName
}

func (s *DescribeSynDbsResponseBodySynDbs) GetRdsVpcUrl() *string {
	return s.RdsVpcUrl
}

func (s *DescribeSynDbsResponseBodySynDbs) GetSynDb() *string {
	return s.SynDb
}

func (s *DescribeSynDbsResponseBodySynDbs) GetSynStatus() *bool {
	return s.SynStatus
}

func (s *DescribeSynDbsResponseBodySynDbs) SetErrorMsg(v string) *DescribeSynDbsResponseBodySynDbs {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetRdsId(v string) *DescribeSynDbsResponseBodySynDbs {
	s.RdsId = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetRdsUserName(v string) *DescribeSynDbsResponseBodySynDbs {
	s.RdsUserName = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetRdsVpcUrl(v string) *DescribeSynDbsResponseBodySynDbs {
	s.RdsVpcUrl = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetSynDb(v string) *DescribeSynDbsResponseBodySynDbs {
	s.SynDb = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetSynStatus(v bool) *DescribeSynDbsResponseBodySynDbs {
	s.SynStatus = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) Validate() error {
	return dara.Validate(s)
}
