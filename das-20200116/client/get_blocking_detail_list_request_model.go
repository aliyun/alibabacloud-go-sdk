// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBlockingDetailListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbNameList(v string) *GetBlockingDetailListRequest
	GetDbNameList() *string
	SetEndTime(v string) *GetBlockingDetailListRequest
	GetEndTime() *string
	SetInstanceId(v string) *GetBlockingDetailListRequest
	GetInstanceId() *string
	SetPageNo(v string) *GetBlockingDetailListRequest
	GetPageNo() *string
	SetPageSize(v string) *GetBlockingDetailListRequest
	GetPageSize() *string
	SetQueryHash(v string) *GetBlockingDetailListRequest
	GetQueryHash() *string
	SetStartTime(v string) *GetBlockingDetailListRequest
	GetStartTime() *string
}

type GetBlockingDetailListRequest struct {
	// The name of the database. Separate multiple database names with commas (,).
	//
	// example:
	//
	// school1,school2
	DbNameList *string `json:"DbNameList,omitempty" xml:"DbNameList,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1682490480548
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4nfalp2ap421312z
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be an integer that is greater than 0. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The hash value of the SQL statement. The hash values of SQL statements of the same type are the same.
	//
	// example:
	//
	// DC08B955CAD25E7B
	QueryHash *string `json:"QueryHash,omitempty" xml:"QueryHash,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1679429913757
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetBlockingDetailListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBlockingDetailListRequest) GoString() string {
	return s.String()
}

func (s *GetBlockingDetailListRequest) GetDbNameList() *string {
	return s.DbNameList
}

func (s *GetBlockingDetailListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetBlockingDetailListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBlockingDetailListRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *GetBlockingDetailListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetBlockingDetailListRequest) GetQueryHash() *string {
	return s.QueryHash
}

func (s *GetBlockingDetailListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetBlockingDetailListRequest) SetDbNameList(v string) *GetBlockingDetailListRequest {
	s.DbNameList = &v
	return s
}

func (s *GetBlockingDetailListRequest) SetEndTime(v string) *GetBlockingDetailListRequest {
	s.EndTime = &v
	return s
}

func (s *GetBlockingDetailListRequest) SetInstanceId(v string) *GetBlockingDetailListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetBlockingDetailListRequest) SetPageNo(v string) *GetBlockingDetailListRequest {
	s.PageNo = &v
	return s
}

func (s *GetBlockingDetailListRequest) SetPageSize(v string) *GetBlockingDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *GetBlockingDetailListRequest) SetQueryHash(v string) *GetBlockingDetailListRequest {
	s.QueryHash = &v
	return s
}

func (s *GetBlockingDetailListRequest) SetStartTime(v string) *GetBlockingDetailListRequest {
	s.StartTime = &v
	return s
}

func (s *GetBlockingDetailListRequest) Validate() error {
	return dara.Validate(s)
}
