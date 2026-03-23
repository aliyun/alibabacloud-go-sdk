// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogReportListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSQLLogReportListRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeSQLLogReportListRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeSQLLogReportListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLLogReportListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSQLLogReportListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSQLLogReportListRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeSQLLogReportListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLLogReportListRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeSQLLogReportListRequest
	GetStartTime() *string
}

type DescribeSQLLogReportListRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSQLLogReportListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogReportListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLLogReportListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSQLLogReportListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLLogReportListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLLogReportListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogReportListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSQLLogReportListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLLogReportListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLLogReportListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSQLLogReportListRequest) SetDBInstanceId(v string) *DescribeSQLLogReportListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogReportListRequest) SetEndTime(v string) *DescribeSQLLogReportListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogReportListRequest) SetOwnerAccount(v string) *DescribeSQLLogReportListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLLogReportListRequest) SetOwnerId(v int64) *DescribeSQLLogReportListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLLogReportListRequest) SetPageNumber(v int32) *DescribeSQLLogReportListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogReportListRequest) SetPageSize(v int32) *DescribeSQLLogReportListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogReportListRequest) SetResourceOwnerAccount(v string) *DescribeSQLLogReportListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLLogReportListRequest) SetResourceOwnerId(v int64) *DescribeSQLLogReportListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLLogReportListRequest) SetStartTime(v string) *DescribeSQLLogReportListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogReportListRequest) Validate() error {
	return dara.Validate(s)
}
