// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsVerifyStatisticRecordsExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromDate(v int64) *CreateSmsVerifyStatisticRecordsExportTaskRequest
	GetFromDate() *int64
	SetOwnerId(v int64) *CreateSmsVerifyStatisticRecordsExportTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateSmsVerifyStatisticRecordsExportTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsVerifyStatisticRecordsExportTaskRequest
	GetResourceOwnerId() *int64
	SetSchemeName(v string) *CreateSmsVerifyStatisticRecordsExportTaskRequest
	GetSchemeName() *string
	SetToDate(v int64) *CreateSmsVerifyStatisticRecordsExportTaskRequest
	GetToDate() *int64
}

type CreateSmsVerifyStatisticRecordsExportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1772294400000
	FromDate             *int64  `json:"FromDate,omitempty" xml:"FromDate,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1772899200000
	ToDate *int64 `json:"ToDate,omitempty" xml:"ToDate,omitempty"`
}

func (s CreateSmsVerifyStatisticRecordsExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsVerifyStatisticRecordsExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) GetFromDate() *int64 {
	return s.FromDate
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) GetSchemeName() *string {
	return s.SchemeName
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) GetToDate() *int64 {
	return s.ToDate
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) SetFromDate(v int64) *CreateSmsVerifyStatisticRecordsExportTaskRequest {
	s.FromDate = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) SetOwnerId(v int64) *CreateSmsVerifyStatisticRecordsExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) SetResourceOwnerAccount(v string) *CreateSmsVerifyStatisticRecordsExportTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) SetResourceOwnerId(v int64) *CreateSmsVerifyStatisticRecordsExportTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) SetSchemeName(v string) *CreateSmsVerifyStatisticRecordsExportTaskRequest {
	s.SchemeName = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) SetToDate(v int64) *CreateSmsVerifyStatisticRecordsExportTaskRequest {
	s.ToDate = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
