// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeginBillingCycle(v string) *CreateReportDefinitionResponseBody
	GetBeginBillingCycle() *string
	SetMetadata(v interface{}) *CreateReportDefinitionResponseBody
	GetMetadata() interface{}
	SetOssBucketName(v string) *CreateReportDefinitionResponseBody
	GetOssBucketName() *string
	SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionResponseBody
	GetOssBucketOwnerAccountId() *int64
	SetOssBucketPath(v string) *CreateReportDefinitionResponseBody
	GetOssBucketPath() *string
	SetReportSourceName(v string) *CreateReportDefinitionResponseBody
	GetReportSourceName() *string
	SetReportSourceType(v string) *CreateReportDefinitionResponseBody
	GetReportSourceType() *string
	SetReportTaskId(v int64) *CreateReportDefinitionResponseBody
	GetReportTaskId() *int64
	SetReportType(v string) *CreateReportDefinitionResponseBody
	GetReportType() *string
	SetRequestId(v string) *CreateReportDefinitionResponseBody
	GetRequestId() *string
	SetSubscribeCreateTime(v string) *CreateReportDefinitionResponseBody
	GetSubscribeCreateTime() *string
}

type CreateReportDefinitionResponseBody struct {
	// example:
	//
	// 2025-05
	BeginBillingCycle *string     `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	Metadata          interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// sh-bill
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// example:
	//
	// bill/
	OssBucketPath *string `json:"OssBucketPath,omitempty" xml:"OssBucketPath,omitempty"`
	// example:
	//
	// OSS
	ReportSourceName *string `json:"ReportSourceName,omitempty" xml:"ReportSourceName,omitempty"`
	// example:
	//
	// OSS
	ReportSourceType *string `json:"ReportSourceType,omitempty" xml:"ReportSourceType,omitempty"`
	// example:
	//
	// 123123
	ReportTaskId *int64 `json:"ReportTaskId,omitempty" xml:"ReportTaskId,omitempty"`
	// example:
	//
	// BillingItemDetailForBillingPeriod
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// example:
	//
	// 340CAB45-0637-5875-9BE4-EFD5750F6BA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2025-05-21 10:36:31
	SubscribeCreateTime *string `json:"SubscribeCreateTime,omitempty" xml:"SubscribeCreateTime,omitempty"`
}

func (s CreateReportDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReportDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReportDefinitionResponseBody) GetBeginBillingCycle() *string {
	return s.BeginBillingCycle
}

func (s *CreateReportDefinitionResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *CreateReportDefinitionResponseBody) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateReportDefinitionResponseBody) GetOssBucketOwnerAccountId() *int64 {
	return s.OssBucketOwnerAccountId
}

func (s *CreateReportDefinitionResponseBody) GetOssBucketPath() *string {
	return s.OssBucketPath
}

func (s *CreateReportDefinitionResponseBody) GetReportSourceName() *string {
	return s.ReportSourceName
}

func (s *CreateReportDefinitionResponseBody) GetReportSourceType() *string {
	return s.ReportSourceType
}

func (s *CreateReportDefinitionResponseBody) GetReportTaskId() *int64 {
	return s.ReportTaskId
}

func (s *CreateReportDefinitionResponseBody) GetReportType() *string {
	return s.ReportType
}

func (s *CreateReportDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReportDefinitionResponseBody) GetSubscribeCreateTime() *string {
	return s.SubscribeCreateTime
}

func (s *CreateReportDefinitionResponseBody) SetBeginBillingCycle(v string) *CreateReportDefinitionResponseBody {
	s.BeginBillingCycle = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetMetadata(v interface{}) *CreateReportDefinitionResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetOssBucketName(v string) *CreateReportDefinitionResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionResponseBody {
	s.OssBucketOwnerAccountId = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetOssBucketPath(v string) *CreateReportDefinitionResponseBody {
	s.OssBucketPath = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportSourceName(v string) *CreateReportDefinitionResponseBody {
	s.ReportSourceName = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportSourceType(v string) *CreateReportDefinitionResponseBody {
	s.ReportSourceType = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportTaskId(v int64) *CreateReportDefinitionResponseBody {
	s.ReportTaskId = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportType(v string) *CreateReportDefinitionResponseBody {
	s.ReportType = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetRequestId(v string) *CreateReportDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetSubscribeCreateTime(v string) *CreateReportDefinitionResponseBody {
	s.SubscribeCreateTime = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
