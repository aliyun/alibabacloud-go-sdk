// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeSQLLogRecordsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeSQLLogRecordsRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeSQLLogRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeSQLLogRecordsRequest
	GetEndTime() *string
	SetForm(v string) *DescribeSQLLogRecordsRequest
	GetForm() *string
	SetOwnerAccount(v string) *DescribeSQLLogRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLLogRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSQLLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSQLLogRecordsRequest
	GetPageSize() *int32
	SetQueryKeywords(v string) *DescribeSQLLogRecordsRequest
	GetQueryKeywords() *string
	SetResourceOwnerAccount(v string) *DescribeSQLLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLLogRecordsRequest
	GetResourceOwnerId() *int64
	SetSQLId(v int64) *DescribeSQLLogRecordsRequest
	GetSQLId() *int64
	SetStartTime(v string) *DescribeSQLLogRecordsRequest
	GetStartTime() *string
	SetUser(v string) *DescribeSQLLogRecordsRequest
	GetUser() *string
}

type DescribeSQLLogRecordsRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database     *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// This parameter is required.
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Form                 *string `json:"Form,omitempty" xml:"Form,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryKeywords        *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SQLId                *int64  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User      *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSQLLogRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLLogRecordsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeSQLLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSQLLogRecordsRequest) GetForm() *string {
	return s.Form
}

func (s *DescribeSQLLogRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLLogRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSQLLogRecordsRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeSQLLogRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLLogRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLLogRecordsRequest) GetSQLId() *int64 {
	return s.SQLId
}

func (s *DescribeSQLLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSQLLogRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DescribeSQLLogRecordsRequest) SetClientToken(v string) *DescribeSQLLogRecordsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetDBInstanceId(v string) *DescribeSQLLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetDatabase(v string) *DescribeSQLLogRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetEndTime(v string) *DescribeSQLLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetForm(v string) *DescribeSQLLogRecordsRequest {
	s.Form = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetOwnerAccount(v string) *DescribeSQLLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetOwnerId(v int64) *DescribeSQLLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageNumber(v int32) *DescribeSQLLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageSize(v int32) *DescribeSQLLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetQueryKeywords(v string) *DescribeSQLLogRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeSQLLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeSQLLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetSQLId(v int64) *DescribeSQLLogRecordsRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetStartTime(v string) *DescribeSQLLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetUser(v string) *DescribeSQLLogRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
